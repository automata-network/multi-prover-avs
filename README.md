## Table of Contents
- [About Multi-Prover AVS](#about-multi-prover-avs)
- [Directory Structure](#directory-structure)
- [AVS Task Description](#avs-task-description)
- [AVS Architecture](#avs-architecture)
- [AVS Workflow](#avs-workflow)
- [Quick Start](#quick-start)

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

```mermaid
sequenceDiagram
  autonumber
    participant O as Operator
    participant A as Aggregator
    participant EC as Eigenlayer Contracts
    participant M as Multi-prover AVS
    participant L as TEE Liveness Contract
    participant AC as Attestation Contracts
 
Note over O,L: Operator registration
O->>EC: Restake and register to Multi-prover AVS
O->>L: Operator Liveness
Note over O,AC: Operator Liveness (Periodically)
loop Every liveness cycle
O->>O: Generate Attestation Report
O->>L: Submit Attesatation Report with pubkey(submitLivenessProof)
L->>+AC: Verify Attestation Report(verifyAttestation)
AC->>AC: Verify MrEncalve and MrSigner
AC->>AC: Verify QE Report
AC->>AC: Verify Quote CertChain
AC->>AC: Verify TCB
AC-->>-L: Return reportData(pubkey)
L->>L: Mark the operator(pubkey) as attested within the liveness cycle
end
 
Note over O,L: Proving state
O-->>O: Fetch task and calculate inside TEE
O->>A: Provide state and signature(submitTask)
A-->>EC: Fetch pubkey by OperaterId
A->>L: Fetch operator's validity (verifyLivenessProof)
L-->>A: Operator's validity
alt Is Valid TEE Prover
A-->>A: Wait and aggregate signatures
A->>M: Submit state and aggregated signatures
else Not Valid TEE Prover
A-->>O: Reject
end
```

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

## Quick Start

Steps for running the multi-prover-avs in simulation node(without the SGX prover):

1. Start a execution node
```
> anvil --fork-url ${holesky_rpc_endpoint} --fork-block-number 1218851
```

2. Prepare the environment
```
> cp .env.example .env
> # vim .env
```

3. Deploy the avs contracts
```
> ./script/deploy.sh init_all --simulation
> # this script will update the configure on config/operator.json and config/aggregator.json
```

4. Run the aggregator
```
go run ./cmd/aggregator
```

5. Run the operator
```
go run ./cmd/operator
```

### With Docker compose

1. build image
```
> ./script/docker.sh build
```

2. init state
```
> ./scripts/docker.sh init_state --simulation
```

3. start aggregator & operator
```
> ./scripts/docker.sh run
```

3. stop all
```
> ./scripts/docker.sh stop
```

