// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;


// For building our own token (currency in this case) we will use ERC20 standard
// https://eips.ethereum.org/EIPS/eip-20 
interface ERC20 {
    function transfer(address _to, uint256 _value) external returns (bool);
    function balanceOf(address _owner) external view returns (uint256);
    function totalSupply() external view returns (uint256);

    //for notifying transfers
    event Transfer(address indexed from, address indexed to, uint256 value);
}

contract EcoKash is ERC20 {
    uint256 private _totalSupply;
    mapping (address => uint256) private balances;

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
    function totalSupply() public view returns (uint256){
        return _totalSupply;
    }
    //function to check balance
    function balanceOf(address account) public view returns (uint256){
        return balances[account];
    }


     // Minting function
    function mint(address to, uint256 amount) internal onlyOwner {
        require(to != address(0), "Cannot mint to zero address");
        _totalSupply += amount; // Increase total supply
        balances[to] += amount; // Add tokens to the recipient's balance

        emit Transfer(address(0), to, amount); // Emit the minting event
    }
}



//Improvements
// MAX_SUPPLY, decimals, 