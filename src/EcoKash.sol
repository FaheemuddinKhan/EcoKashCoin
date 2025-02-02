// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;


// For building our own token (currency in this case) we will use ERC20 standard
// https://eips.ethereum.org/EIPS/eip-20 
interface ERC20 {
    function transfer(address to, uint256 value) external returns (bool);
    function balanceOf(address owner) external view returns (uint256);
    function totalSupply() external view returns (uint256);
    function approve(address spender, uint256 value) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function transferFrom(address from, address to, uint256 value) external returns (bool);

    //for notifying transfers
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract EcoKash is ERC20 {
    uint256 private _totalSupply;
    mapping (address => uint256) private balances;
    mapping (address => mapping (address => uint256)) private allowances;

    address private owner;

    string public name = "EcoKash";
    string public symbol = "EKC";
    uint8 public decimals = 18;


    uint256 public constant MAX_SUPPLY = 1_000_000 * 10**18;

    constructor(uint initialSupply) {
        owner = msg.sender;
        mint(owner, initialSupply);
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can call this function");
        _;
    }

    //function to transer tokens from sender
    function transfer(address to, uint256 value) public returns (bool){
        require(balances[msg.sender] >= value, "Insufficient balance");
        require(to != address(0), "Cannot transfer to zero address");

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

    function approve(address spender, uint256 value)  public returns (bool) {
        require (spender != address(0), "Cannot approve for zero address");
        require(allowances[msg.sender][spender] == 0, "Increase or decrease the existing allowances");
        allowances[msg.sender][spender] = value;

        emit Approval(msg.sender, spender, value);
        return true;
    }

    function allowance(address _owner, address spender) public view returns (uint256) {
        return allowances[_owner][spender];
    }

    function increaseAllowances(address spender, uint256 addedValue) public returns (bool) {
        require(addedValue >= 0, "Added value should be greater than zero.");
        require(spender != address(0),"Spender address cannot be zero.");
        allowances[msg.sender][spender] += addedValue;

        emit Approval(msg.sender, spender, allowances[msg.sender][spender]);
        return true;
    }

    function decreaseAllowances(address spender, uint256 subtractedValue) public returns (bool) {
        require(subtractedValue >= 0, "value should be greater than zero.");
        require(spender != address(0),"spender address cannot be zero.");
        uint256 currentAllowance = allowances[msg.sender][spender];
        require(currentAllowance >= subtractedValue, "Decreased allowance below zero.");
        allowances[msg.sender][spender] -= subtractedValue;

        emit Approval(msg.sender, spender, allowances[msg.sender][spender]);
        return true;
    }

    function transferFrom(address from, address to, uint256 value) public returns (bool) {
        require(balances[from] >= value,"Insufficient balance");
        require(allowances[from][msg.sender] >= value,"Allowance exceeded");
        require(to != address(0), "Transfer to zero address");
        balances[to] += value;
        balances[from] -= value;
        allowances[from][msg.sender] -= value;

        emit Transfer(from, to, value);
        return true;
    }

     // Minting function
    function mint(address to, uint256 amount) public onlyOwner {
        require(to != address(0), "Cannot mint to zero address");
        _totalSupply += amount; // Increase total supply
        balances[to] += amount; // Add tokens to the recipient's balance

        emit Transfer(address(0), to, amount); // Emit the minting event
    }
}



//Improvements
// MAX_SUPPLY, decimals, 