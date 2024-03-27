## Table of Contents
- [About Multi-Prover AVS](#about-multi-prover-avs)
- [Directory Structure](#directory-structure)
- [AVS Task Description](#avs-task-description)
- [AVS Architecture](#avs-architecture)
- [AVS Workflow](#avs-workflow)
- [TEE Committee and Quorum](#tee-committee-and-quorum)
## About Multi-Prover AVS
The Automata Multi-Prover AVS target to build a robust, fortified prover system through the use of diverse, decentralized TEE committees.
![Automata Multi-Prover AVS Design](/assets/multiprover-design.png)

Read this [blog](https://www.notion.so/atanetwork/Elevating-ZK-Security-with-Multi-Prover-AVS-cc1f4d1fc0b341d4a4b90a16f7b8bbb3) to understand more about the Multi-Prover AVS.

## Directory Structure
<pre>
├── <a href="./contracts/">contracts</a>: Solidity contracts, including the AVS contracts and the attestation layer contracts.
│ ├── <a href="./contracts/dcap-v3-attestation/">dcap-v3-attestation</a>: On-chain verification library for Dcap attestation of Intel SGX.
│ ├── <a href="./contracts/src/">src</a>: Source files for AVS contracts.
│ └── <a href="./contracts/test/">test</a>: Tests for smart contracts.
├── <a href="./operator/">operator</a>: The operator implementation.
├── <a href="./aggregator/">aggregator</a>: The aggregator implementation.
├── <a href="https://github.com/automata-network/sgx-prover/tree/avs">sgx-prover</a>: the sgx version of TEE prover.
</pre>

## AVS Task Description
Task definition: A state transition or computational process seeking to leverage the independent execution within a Trusted Execution Environment (TEE) to ascertain its correctness.

```solidity
struct StateHeader {
    uint256 identifier;
    bytes metadata;
    bytes state;
}
```
This is the structure of the state header submitted ty provers, below is the detailed explaination:
- **identifier**: identifier of the handled task, it can be used to distinguish different kinds of tasks and used to calculator the contribution of each operators
- **metadata**: metadata that describe the specific task, for example `keccak256(abi.encodePacked(chainID, blockNumber))` is the metadata for the task to prove blockchain state at specific block height
- **state**: the final state produced by the TEE prover, it can be either a root state of blockchain, or statement proved by a zk circuit

## AVS Architecture
The architecture of the AVS contains:
- [Eigenlayer core contracts](https://github.com/Layr-Labs/eigenlayer-contracts)
- AVS contracts
    - ServiceManager which allow operators to submit tasks, reward and slash logic will be added in the future
- Attestation contracts
    - Manage the register/deregister and livenness of various TEE provers, it will verify attestation of different TEE platforms such as Intel SGX, AMD SEV, ARM TrustZone and so on
    - TEEProverRegister is the interface for the attestation layer used by the operators and aggregator
- Aggregator
    - Aggregate the BLS signatures from operators and submit the aggregated state to AVS
    - Interact with the Automata attestation layer to check the validity of each prover(operator), those who failed to pass the attestation verification or liveness challenge will be rejected to handle tasks until they are valid again.
- Operator
    - Fetch state proofs from TEE prover and submit it to the aggregator
- TEE prover
    - TEE prover that prove the final state of a given task, for example, a prover of zk-rollup L2 will execute blocks inside TEE and produce the root state at specific block

## AVS Workflow
Below is a detailed diagram of the workflow

[![](https://mermaid.ink/img/pako:eNp1VMGOmzAQ_ZURl6Zq8gMcVqIsrVYKJd1E2wsXBybEDdjUNqnQav-9M4Zkk4XewH4z89688bwGhS4xCAOLfzpUBT5KURnR5ApAdE6rrtmj4T-AVhgnC9kK5SADYSFr0QinZ64jvo6qymA1D0hiRiSyQlWLHg3EWjkjCmen2JShaVc7uWqNPhM4etlOYWuG7ZIE1vKMCq295pzh58tHzqF1wkmt7urn6od2CL5UtlyHV6FAeqQlHIfkKls9PCRxCM-c5YQgVDkiKNDpGc4ccZvvQjVX9zWjeAYEiw0aqUtZiLruP-eq1rqFhCJ6qC-Yoi9qHAplIXynQ0qCd1KfsdXGXclsu30j3YAQdxD4K90R2m5_wn5hPezCZWO0PhCFNeX4wmxfiNqhn6mzOPubmwsKi2KKuwlLTaJI1XnoYWq2slI8dh9xP5Mr-8lVx_2L0bj4KKSaAnbxV39Ip6s1m-Y6o8gvTvdIwheD0FEUIVJhTuCOZMroxAXBsyO8Hix9j6TyuI8moCr9NE3mibp3lqoCbog3a3DrG7riCE7Yk28DNaToanZPKitL5OEeXIvGFHTmU3i4pZ4J0oSjVTvKw61ejVM6ZB8kwL4f5wvNU8kgr3iAXOR-snAWtSyl62F0cWr_wCabhrBsUTt4svDCR_5pbvxbGElR3C9Bk8fkxbgr8F2GHWil1wl9l3pFl3dwrC0Ctfp_BTN2_TfyQiBrgmXQoGmELGn9vfKOyANyscE8COmzxIOg95sHuXojKC_Dba-KIHSmw2XQtSXVH7dlEB4E1V4GSMq1SYeV6jfr2z84HOKq?type=jpg)](https://mermaid.live/edit#pako:eNp1VMGOmzAQ_ZURl6Zq8gMcVqIsrVYKJd1E2wsXBybEDdjUNqnQav-9M4Zkk4XewH4z89688bwGhS4xCAOLfzpUBT5KURnR5ApAdE6rrtmj4T-AVhgnC9kK5SADYSFr0QinZ64jvo6qymA1D0hiRiSyQlWLHg3EWjkjCmen2JShaVc7uWqNPhM4etlOYWuG7ZIE1vKMCq295pzh58tHzqF1wkmt7urn6od2CL5UtlyHV6FAeqQlHIfkKls9PCRxCM-c5YQgVDkiKNDpGc4ccZvvQjVX9zWjeAYEiw0aqUtZiLruP-eq1rqFhCJ6qC-Yoi9qHAplIXynQ0qCd1KfsdXGXclsu30j3YAQdxD4K90R2m5_wn5hPezCZWO0PhCFNeX4wmxfiNqhn6mzOPubmwsKi2KKuwlLTaJI1XnoYWq2slI8dh9xP5Mr-8lVx_2L0bj4KKSaAnbxV39Ip6s1m-Y6o8gvTvdIwheD0FEUIVJhTuCOZMroxAXBsyO8Hix9j6TyuI8moCr9NE3mibp3lqoCbog3a3DrG7riCE7Yk28DNaToanZPKitL5OEeXIvGFHTmU3i4pZ4J0oSjVTvKw61ejVM6ZB8kwL4f5wvNU8kgr3iAXOR-snAWtSyl62F0cWr_wCabhrBsUTt4svDCR_5pbvxbGElR3C9Bk8fkxbgr8F2GHWil1wl9l3pFl3dwrC0Ctfp_BTN2_TfyQiBrgmXQoGmELGn9vfKOyANyscE8COmzxIOg95sHuXojKC_Dba-KIHSmw2XQtSXVH7dlEB4E1V4GSMq1SYeV6jfr2z84HOKq)

Components:
- [Operator](./operator)
- [Aggregator](./aggregator)
- [MultiProver AVS](./contracts/src/core/MultiProverServiceManager.sol)
- [TEE Liveness Contract](./contracts/src/core/TEELivenessVerifier.sol)
- [Attestation Contract](https://github.com/automata-network/sgx-prover/blob/avs/verifier/contracts/AutomataDcapV3Attestation.sol)

The workflow is divided into two parts:
- Setup
    - Follow the [Eigenlayer's doc](https://docs.eigenlayer.xyz/eigenlayer/overview) to stake and register as operator of Multi-prover AVS
    - Generate attestation and register as TEE prover, attestation and its generating process differs depending on the TEE technology. For example, [dcap-v3-attestation](./contracts/dcap-v3-attestation/) is the contracts of verifying Dcap attestation of Intel SGX
- Working
    - Except what operators should do to handle tasks, they must complete liveness challenge periodically, otherwise they will be treated as invalid and their submission will be rejected by the aggregator
    - Operators fetch new task and finish the calculation inside TEE
    - Operators sign the final state and send it together with signature to aggregator
    - Aggregator will fetch operator's validity before accepting its submission
    - Aggregator aggregate all the BLS signature and submit to the AVS service manager

## TEE Committee and Quorum
![Committee and Quorum](./assets/committee-and-quorum.jpg)

**TEE Committee** is a set of quorum that is responsible to handle a specific type of task. For example, proving the root state of scroll at a particular block height. Operators do not need to actively choose a committee, but automatically belong to a committee by joining quorums. The introduction of `TEE Committee` facilitates a more organized structuring of operators and tasks. And lays the groundwork for future enhancements, including the rewarding mechanism and constraints of stake distribution across quorums.

The concept of a **TEE Quorum** aligns with the quorum definition utilized by Eigenlayer, but each quorum is associated with a TEE platform, such as Intel SGX. Each quorum belongs to a committee, operators can choose to join any quorum, But only the votes from operators possessing the requisite attestation will be accepted by the aggregator.
