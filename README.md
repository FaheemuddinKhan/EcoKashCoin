# Project README

## **Foundry Commands**

### 1. **Initialize a New Project**

To initialize a new Foundry project, use:

```bash
forge init <project-name>
```

This will create a new project structure with necessary directories like `src`, `lib`, and `test`.

### 2. **Test Contracts**

To test your smart contracts, use:

Run all tests in the test folder:

```bash
forge test
```

Run a specific test function:

```bash
forge test --match-test <test-name>  # Run a specific test function
# Example: forge test --match-test testMinting
```

### 3. **Build Contracts**

To build your smart contracts, use:

```bash
forge build
```

This command compiles the Solidity contracts and generates the artifacts in the `out` folder, which will include the bytecode and ABI.

**Note:** If you are using `solc` as an alternative, refer to the [solc commands section below](#solc-commands) for how to manually compile contracts.

### 3. **Start Local Blockchain**

Start Anvil: To start Anvil, you can run the following command:

```bash
anvil
```

This will start a local Ethereum node with a default configuration, which listens on http://127.0.0.1:8545. You can change the port or other parameters by passing arguments to Anvil.

This will start a local Ethereum network at `http://127.0.0.1:8545`, where you can interact with your contracts.

### 4. **Run (Build + Deploy on the go) a Contract Using a Script**

To run a contract using a deployment script, run:

```bash
forge script <script-file> --rpc-url <rpc-url> --private-key <your-private-key> --broadcast
```

### 5. Check the Contract’s Bytecode

Inspect the bytecode of the compiled contract:

```bash
forge inspect <contract-name> --bytecode
# Example: forge inspect EcoKash --bytecode
```

### 6. **Interactive Console**

To interact with the local blockchain and execute Solidity code directly from the command line, run:

```bash
forge console --rpc-url http://127.0.0.1:8545
```

### 7. Clean the Project

Remove the build artifacts from the out folder to ensure a fresh build:

```bash
forge clean  # Deletes the build folder
# Example: forge clean (from the root directory of the project)
```

In the console, you can deploy contracts, query contract states, and send transactions.

---

## **solc Commands**

If you prefer to use **solc** (the Solidity compiler) directly for compiling and deploying contracts, you can do so by using the following commands.

### 1. **Compiling Contracts**

To compile a Solidity contract, use:

```bash
solc --bin --abi <contract-file.sol> -o <output-folder>
```

This generates the compiled bytecode (`.bin`) and ABI (`.abi`) files in the specified output folder.

### 2. **Compiling with Optimization for Production**

To compile the contract with optimization enabled (for production environments), run:

```bash
solc --optimize --bin --abi <contract-file.sol> -o <output-folder>
```

### 3. **Manual Contract Deployment**

After compiling with `solc`, you can deploy the contract using frameworks like **ethers.js**. Here's an example of a deployment script using **ethers.js**:

1. **Install ethers.js:**

   ```bash
   npm install ethers
   ```

2. **Write the Deployment Script:**

   ```javascript
   const { ethers } = require("ethers");
   const fs = require("fs");

   const provider = new ethers.JsonRpcProvider("http://127.0.0.1:8545");
   const signer = new ethers.Wallet("YOUR_PRIVATE_KEY", provider);

   const abi = JSON.parse(fs.readFileSync("out/EcoKash.abi"));
   const bytecode = fs.readFileSync("out/EcoKash.bin").toString();

   async function deploy() {
     const factory = new ethers.ContractFactory(abi, bytecode, signer);
     const contract = await factory.deploy(1000000); // Initial supply as constructor argument
     console.log(`Contract deployed to: ${contract.address}`);
   }

   deploy();
   ```

### 4. **Deploying with Truffle (Optional)**

You can also deploy contracts using **Truffle** if you prefer working with a higher-level framework. Use the Truffle deployment commands after compiling with `solc`.

---

### Summary of Commands

#### **Foundry Commands:**

- **Initialize a new project:**

  ```bash
  forge init <project-name>
  ```

- **Build contracts:**

  ```bash
  forge build
  ```

- **Start local blockchain:**

  ```bash
  forge network start
  ```

- **Deploy a contract:**

  ```bash
  forge script <script-file> --rpc-url <rpc-url> --private-key <your-private-key> --broadcast
  ```

- **Run interactive console:**

  ```bash
  forge console --rpc-url http://127.0.0.1:8545
  ```

#### **solc Commands:**

- **Compile contracts (generate bytecode and ABI):**

  ```bash
  solc --bin --abi <contract-file.sol> -o <output-folder>
  ```

- **Compile with optimization (for production):**

  ```bash
  solc --optimize --bin --abi <contract-file.sol> -o <output-folder>
  ```

### **Forge**

1. **`forge build`**
   - Compiles the smart contract files in the project.

2. **`forge test`**
   - Runs the unit tests for your smart contracts.

3. **`forge deploy`**
   - Deploys the smart contract to a specified network.

4. **`forge console`**
   - Opens an interactive console with access to your deployed contracts.

5. **`forge create <ContractName>`**
   - Deploys a contract and provides the deployed contract’s address.

6. **`forge snapshot`**
   - Saves the state of a smart contract at a specific point.

7. **`forge verify-contract`**
   - Verifies a contract on an Ethereum network (e.g., Etherscan).

---

### **Anvil**

1. **`anvil`**
   - Starts a local Ethereum test node.

2. **`anvil --fork`**
   - Starts a local node forked from an existing network (e.g., mainnet, ropsten).

3. **`anvil accounts`**
   - Lists the Ethereum accounts available on the local node.

4. **`anvil balance <address>`**
   - Retrieves the balance of an Ethereum address on the local testnet.

5. **`anvil mine`**
   - Manually mines a block on the local node (useful for testing transactions).

6. **`anvil block <blockNumber>`**
   - Fetches information about a specific block number.

---

### **Caste**

1. **`caste call <contractAddress> <functionSignature> <arguments>`**

   - Calls a contract function on a specified address without publishing a transaction (read-only operation).
   Example:

   ```bash
   caste call <contractAddress> "balanceOf(address)" "0x123..."
   ```

2. **`caste send <contractAddress> <functionSignature> <arguments>`**

   - Sends a transaction to execute a contract function (writes data to the blockchain).

   Example:

   ```bash
   caste send <contractAddress> "transfer(address,uint256)" "0x456... 100"
   ```

3. **`caste estimate <transactionData>`**
   - Estimates the gas cost of a transaction.

4. **`caste block`**
   - Gets information about the latest block in the chain.

5. **`caste receipt <transactionHash>`**
   - Retrieves the receipt of a transaction.

6. **`caste decode <data>`**
   - Decodes ABI-encoded data from a transaction.

7. **`caste gas-price`**
   - Displays the current gas price on the network.

8. **`caste logs <address>`**

    - Retrieves logs for a specific address.

---
