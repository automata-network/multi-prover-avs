// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {IAttestation} from "../interfaces/IAttestation.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract TEELivenessVerifier is OwnableUpgradeable {
    struct Pubkey {
        bytes32 x;
        bytes32 y;
    }

    struct Prover {
        Pubkey pubkey;
        uint256 time;
    }

    struct ReportDataV2 {
        Pubkey pubkey;
        uint256 referenceBlockNumber;
        bytes32 referenceBlockHash;
        bytes32 proverAddressHash;
    }

    mapping(bytes32 => bool) public attestedReports;
    mapping(bytes32 => Prover) public attestedProvers; // prover's pubkey => attestedTime

    uint256 public attestValiditySeconds = 3600;

    IAttestation public dcapAttestation;

    // added at v2
    uint256 public maxBlockNumberDiff;


    constructor() {
        _disableInitializers();
    }

    function initialize(address _initialOwner, address _attestationAddr, uint256 _maxBlockNumberDiff, uint256 _attestValiditySeconds) public initializer {
        dcapAttestation = IAttestation(_attestationAddr);
        maxBlockNumberDiff = _maxBlockNumberDiff;
        attestValiditySeconds = _attestValiditySeconds;
        _transferOwnership(_initialOwner);
    }

    function reinitialize(uint8 i, address _initialOwner, address _attestationAddr, uint256 _maxBlockNumberDiff, uint256 _attestValiditySeconds) public reinitializer(i) {
        dcapAttestation = IAttestation(_attestationAddr);
        maxBlockNumberDiff = _maxBlockNumberDiff;
        attestValiditySeconds = _attestValiditySeconds;
        _transferOwnership(_initialOwner);
    }

    function changeMaxBlockNumberDiff(uint256 _maxBlockNumberDiff) public onlyOwner {
        maxBlockNumberDiff = _maxBlockNumberDiff;
    }

    function changeAttestationImpl(address _attestationAddr) public onlyOwner {
        dcapAttestation = IAttestation(_attestationAddr);
    }

    function changeAttestValiditySeconds(uint256 val) public onlyOwner {
        attestValiditySeconds = val;
    }

    function verifyMrEnclave(bytes32 _mrenclave) public view returns (bool) {
        return dcapAttestation.verifyMrEnclave(_mrenclave);
    }

    function verifyMrSigner(bytes32 _mrsigner) public view returns (bool) {
        return dcapAttestation.verifyMrSigner(_mrsigner);
    }

    function submitLivenessProofV2(
        ReportDataV2 calldata _data,
        bytes calldata _report
    ) public {
        checkBlockNumber(_data.referenceBlockNumber, _data.referenceBlockHash);
        bytes32 dataHash = keccak256(abi.encode(_data));

        (bool succ, bytes memory reportData) = dcapAttestation
            .verifyAttestation(_report);
        require(succ, "attestation report validation fail");

        bytes32 reportHash = keccak256(_report);
        require(!attestedReports[reportHash], "report is already used");

        (, bytes32 reportDataHash) = splitBytes64(reportData);
        require(dataHash == reportDataHash, "report data hash mismatch");

        Prover memory prover = Prover(_data.pubkey, block.timestamp);
        attestedProvers[keccak256(abi.encode(_data.pubkey.x, _data.pubkey.y))] = prover;
        attestedReports[reportHash] = true;
    }

    function verifyLivenessProof(
        bytes32 pubkeyX,
        bytes32 pubkeyY
    ) public view returns (bool) {
        bytes32 signer = keccak256(abi.encode(pubkeyX, pubkeyY));
        return
            attestedProvers[signer].time + attestValiditySeconds >
            block.timestamp;
    }

    function verifyAttestationV2(
        bytes32 pubkeyX,
        bytes32 pubkeyY,
        bytes calldata data
    ) public view returns (bool) {
        (bool succ, bytes memory reportData) = dcapAttestation
            .verifyAttestation(data);
        if (!succ) {
            return false;
        }

        (bytes32 x, bytes32 y) = splitBytes64(reportData);
        if (x != pubkeyX || y != pubkeyY) {
            return false;
        }

        return true;
    }

    function splitBytes64(
        bytes memory b
    ) private pure returns (bytes32, bytes32) {
        require(b.length >= 64, "Bytes array too short");

        bytes32 x;
        bytes32 y;
        assembly {
            x := mload(add(b, 32))
            y := mload(add(b, 64))
        }
        return (x, y);
    }

    // this function will make sure the attestation report
    function checkBlockNumber(uint256 blockNumber, bytes32 blockHash) private view {
        require(
            blockNumber < block.number && block.number - blockNumber < maxBlockNumberDiff,
            "invalid block number"
        );

        require(blockhash(blockNumber) == blockHash, "block number mismatch");
    }
}
