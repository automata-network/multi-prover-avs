// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "../src/core/AttestationVerifier.sol";
import "../src/interfaces/IAttestation.sol";

contract MockDcapAttestation is IAttestation {

    bytes public output;

    function setOutput(bytes memory _output) external {
        output = _output;
    }

    function getBp() external pure returns (uint16) {
        return 0;
    }

    function verifyAndAttestOnChain(bytes calldata) external payable returns (bool success, bytes memory) {
        return (true, output);
    }

}

contract AttestationVerifierTest is Test {

    MockDcapAttestation internal dcap;
    AttestationVerifier internal verifier;

    function setUp() public {
        dcap = new MockDcapAttestation();
        verifier = new AttestationVerifier(address(dcap));
        dcap.setOutput(validSgxAttestationOutput());
    }

    function testRejectsDebugModeSgxQuote() public {
        bytes memory quote = sgxQuoteWithAttributes(hex"07000000000000000700000000000000");

        vm.expectRevert(AttestationVerifier.DEBUG_ENCLAVE_NOT_ALLOWED.selector);
        verifier.verifyAttestation(quote);
    }

    function testAcceptsProductionModeSgxQuote() public {
        bytes memory quote = sgxQuoteWithAttributes(hex"05000000000000000500000000000000");

        bytes memory reportData = verifier.verifyAttestation(quote);

        assertEq(reportData.length, 64);
    }

    function sgxQuoteWithAttributes(bytes16 attributes) internal pure returns (bytes memory quote) {
        quote = new bytes(432);
        quote[4] = 0x00;
        quote[5] = 0x00;
        quote[6] = 0x00;
        quote[7] = 0x00;

        for (uint256 i = 0; i < 16; i++) {
            quote[96 + i] = attributes[i];
        }
    }

    function validSgxAttestationOutput() internal pure returns (bytes memory output) {
        output = new bytes(397);
        for (uint256 i = 333; i < 397; i++) {
            output[i] = bytes1(uint8(i - 333 + 1));
        }
    }

}
