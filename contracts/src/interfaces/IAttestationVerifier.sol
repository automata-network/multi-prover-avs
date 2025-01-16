//SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

interface IAttestationVerifier {

    function verifyAttestation(bytes calldata _report) external returns (bytes memory);
    function verifyMrEnclave(bytes32 _mrEnclave) external view returns (bool);
    function verifyMrSigner(bytes32 _mrSigner) external view returns (bool);

}
