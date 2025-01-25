// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;


// For building our own token (currency in this case) we will use ERC20 standard
// https://eips.ethereum.org/EIPS/eip-20 
interface ERC20 {
    function transfer(address to, uint256 value) external returns (bool);
    function balanceOf(address account) external view returns (uint256);
    function totalSupply() external view returns (uint256);
    function approve(address spender, uint256 value) external returns (bool);
    function transferFrom(address from, address to, uint256 value) external returns (bool);
    //for notifying transfers
    event Transfer(address indexed from, address indexed to, uint256 value);
}

contract EcoKash is ERC20 {
    uint256 private _totalSupply;

    mapping (address => uint256) private balances;
    mapping(address => mapping(address => uint256)) private allowances;

    address private _owner;

    constructor(uint initialSupply) {
        _owner = msg.sender;
        mint(msg.sender, initialSupply);
    }

    modifier onlyOwner() {
        require(msg.sender == _owner, "Only the owner can call this function");
        _;
    }

    //function to transer tokens from sender
    function transfer(address to, uint256 value) public returns (bool){
        require(balances[msg.sender] >= value, "Insufficient balance");
        balances[msg.sender] = balances[msg.sender] - value;
        balances[to] = balances[to] + value;

        emit Transfer(msg.sender, to, value);
        return true;
    }
    //function to check total supply
    function totalSupply() external view returns (uint256){
        return _totalSupply;
    }
    //function to check balance
    function balanceOf(address account) external view returns (uint256){
        return balances[account];
    }


     // Minting function
    function mint(address to, uint256 amount) internal onlyOwner {
        require(to != address(0), "Cannot mint to zero address");
        _totalSupply += amount; // Increase total supply
        balances[to] += amount; // Add tokens to the recipient's balance

        emit Transfer(address(0), to, amount); // Emit the minting event
    }


    function approve(address spender, uint256 value) external returns (bool) {
        require(spender != address(0), "Cannot approve to zero address");
        require(value <= balances[msg.sender], "Insufficient balance");
        allowances[msg.sender][spender] = value;
        return true;
    }

    function transferFrom(address from, address to, uint256 value) external returns (bool) {
        require(balances[from] >= value, "Insufficient balance");
        require(from != address(0), "Cannot transfer from zero address");
        require(to != address(0), "Cannot transfer to zero address");

        allowances[from][msg.sender] -= value;
        balances[from] -= value;
        balances[to] += value;

        emit Transfer(from, to, value);
        return true;
    }
}



// Improvements
// MAX_SUPPLY for limiing the supply, making it scarce
// decimals, is used for the precision of the token, how many parts can it be divided into ex: 1 USD = 100 cents, 1 GBP = 100 pence, 1 INR 100 paise.
// burn function to burn tokens, only owner can burn, for making it scarce


// Problem with approve

// When you use approve in ERC20 tokens, you tell the contract how many tokens someone else (a "spender") can use on your behalf. However, if you update the allowance without resetting it to 0 first, bad things can happen.

// Example Scenario
// You approve Bob to spend 100 tokens for you.
// Bob spends 50 tokens, so now Bob has 50 tokens left to spend.
// You decide to let Bob spend 200 tokens, so you call approve(Bob, 200).
// The Problem
// While you're updating Bob's allowance:

// Bob can sneak in a transaction to use the remaining 50 tokens before your new allowance of 200 takes effect.
// Now, Bob has 50 + 200 = 250 tokens to spend, even though you only wanted him to have 200 tokens.


// Example (OpenZeppelin Implementation)

// contract MyToken is ERC20, ERC20Permit {
//     constructor() ERC20("MyToken", "MTK") ERC20Permit("MyToken") {}
// }
// Here, ERC20Permit is imported from OpenZeppelin, a popular library of audited smart contract templates.

// The ERC20Permit contract implements the permit function and integrates seamlessly with ERC20.

// ERC20Permit saves Gas?
// No Separate approve Transaction:

// In the standard process, you first pay gas for approve, then pay gas for the actual operation (e.g., swap or transferFrom).
// With ERC20Permit, you skip the approve step entirely.
// One Transaction Instead of Two:

// Instead of:
// approve (gas-heavy).
// Token operation (e.g., transferFrom, also gas-heavy).
// You do only:
// A single transaction (e.g., permit + transferFrom).

