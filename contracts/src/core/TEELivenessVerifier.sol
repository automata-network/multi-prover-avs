// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {IAttestationVerifier} from "../interfaces/IAttestationVerifier.sol";
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

    error INVALID_REPORT();
    error INVALID_REPORT_DATA();

    mapping(bytes32 => bool) public attestedReports;
    mapping(bytes32 => Prover) public attestedProvers; // prover's pubkey => attestedTime

    uint256 public attestValiditySeconds = 3600;

    IAttestationVerifier public attestationVerifier;

    // added at v2
    uint256 public maxBlockNumberDiff;
    mapping(bytes32 proverKey => address proverAddr) public attestedProverAddr;

    constructor() {
        _disableInitializers();
    }

    function initialize(
        address _initialOwner,
        address _attestationAddr,
        uint256 _maxBlockNumberDiff,
        uint256 _attestValiditySeconds
    ) public initializer {
        attestationVerifier = IAttestationVerifier(_attestationAddr);
        maxBlockNumberDiff = _maxBlockNumberDiff;
        attestValiditySeconds = _attestValiditySeconds;
        _transferOwnership(_initialOwner);
    }

    function reinitialize(
        uint8 i,
        address _initialOwner,
        address _attestationAddr,
        uint256 _maxBlockNumberDiff,
        uint256 _attestValiditySeconds
    ) public reinitializer(i) {
        attestationVerifier = IAttestationVerifier(_attestationAddr);
        maxBlockNumberDiff = _maxBlockNumberDiff;
        attestValiditySeconds = _attestValiditySeconds;
        _transferOwnership(_initialOwner);
    }

    function changeMaxBlockNumberDiff(uint256 _maxBlockNumberDiff) public onlyOwner {
        maxBlockNumberDiff = _maxBlockNumberDiff;
    }

    function changeAttestationImpl(address _attestationAddr) public onlyOwner {
        attestationVerifier = IAttestationVerifier(_attestationAddr);
    }

    function changeAttestValiditySeconds(uint256 val) public onlyOwner {
        attestValiditySeconds = val;
    }

    function verifyMrEnclave(bytes32 _mrenclave) public view returns (bool) {
        return attestationVerifier.verifyMrEnclave(_mrenclave);
    }

    function verifyMrSigner(bytes32 _mrsigner) public view returns (bool) {
        return attestationVerifier.verifyMrSigner(_mrsigner);
    }

    function submitLivenessProofV2(ReportDataV2 calldata _data, bytes calldata _report) public payable {
        checkBlockNumber(_data.referenceBlockNumber, _data.referenceBlockHash);
        bytes32 dataHash = keccak256(abi.encode(_data));

        (bytes memory reportData) = attestationVerifier.verifyAttestation{value: msg.value}(_report);

        bytes32 reportHash = keccak256(_report);
        require(!attestedReports[reportHash], "report is already used");

        (bytes32 proverBytes, bytes32 reportDataHash) = splitBytes64(reportData);
        require(dataHash == reportDataHash, "report data hash mismatch");

        Prover memory prover = Prover(_data.pubkey, block.timestamp);
        bytes32 proverKey = keccak256(abi.encode(_data.pubkey.x, _data.pubkey.y));
        attestedProvers[proverKey] = prover;
        attestedProverAddr[proverKey] = address(uint160(uint256(proverBytes)));
        attestedReports[reportHash] = true;
    }

    function verifyLivenessProof(bytes32 pubkeyX, bytes32 pubkeyY) public view returns (bool) {
        bytes32 signer = keccak256(abi.encode(pubkeyX, pubkeyY));
        return attestedProvers[signer].time + attestValiditySeconds > block.timestamp;
    }

    function verifyLivenessProofV2(bytes32 pubkeyX, bytes32 pubkeyY, address proverKey) public view returns (bool) {
        bytes32 signer = keccak256(abi.encode(pubkeyX, pubkeyY));
        bool succ = attestedProvers[signer].time + attestValiditySeconds > block.timestamp;
        if (!succ) {
            return false;
        }
        return attestedProverAddr[signer] == proverKey;
    }

    function verifyAttestationV2(bytes32 pubkeyX, bytes32 pubkeyY, bytes calldata data) public returns (bool) {
        bytes memory reportData = attestationVerifier.verifyAttestation(data);

        (bytes32 x, bytes32 y) = splitBytes64(reportData);
        if (x != pubkeyX || y != pubkeyY) {
            return false;
        }

        return true;
    }

    function splitBytes64(bytes memory b) private pure returns (bytes32, bytes32) {
        require(b.length >= 64, "Bytes array too short");

        bytes32 x;
        bytes32 y;
        assembly {
            x := mload(add(b, 32))
            y := mload(add(b, 64))
        }
        return (x, y);
    }

    // this function will make sure the attestation report generated in recent ${maxBlockNumberDiff} blocks
    function checkBlockNumber(uint256 blockNumber, bytes32 blockHash) private view {
        require(blockNumber < block.number, "invalid block number");
        require(block.number - blockNumber < maxBlockNumberDiff, "block number out-of-date");

        require(blockhash(blockNumber) == blockHash, "block number mismatch");
    }

    function estimateBaseFeeVerifyOnChain(bytes calldata rawQuote) external payable returns (uint256) {
        return attestationVerifier.estimateBaseFeeVerifyOnChain{value: msg.value}(rawQuote);
    }

}
