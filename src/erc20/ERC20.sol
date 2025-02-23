// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
import {IERC20} from "../interface/IERC20.sol";
import {Ownable} from "../ownable/ownable.sol";
import {InvalidAddress} from "../errors/error.sol";


abstract contract ERC20 is IERC20, Ownable {
    uint256 private _totalSupply;
    mapping(address=>uint256) private _balances;
    mapping(address=>mapping(address=>uint256)) private _allowances;

    string private _name;
    string private _symbol;
    uint8 private _decimals;

    error InsufficientAllowance();
    error InsufficientBalance();
    error ExceedsInitialSupplyThreshold();

    constructor(string memory name, string memory symbol, uint256 initialSupply ){
        _name = name;
        _symbol = symbol;
        if (initialSupply <= 1000 *10**18) {
            _updateBalances(address(0), msg.sender, initialSupply);
        } else {
            revert ExceedsInitialSupplyThreshold();
        }
    }

    function totalSupply() external view returns (uint256){
        return _totalSupply;
    }

    function balanceOf(address account) external view returns (uint256){
        return _balances[account];
    }

    function allowance(address owner, address spender) public view returns (uint256) {
        return _allowances[owner][spender];
    }

    function transfer(address to, uint256 value) public returns (bool)  {
        address from = msg.sender;
        _transfer(from, to, value);
        return true;
    }

    function transferFrom(address from, address to, uint256 value) public returns (bool){
        address spender = msg.sender;
        _spendAllowance(from, spender, value);
        _transfer(from, to, value);
        return true;
    }

    function approve(address spender, uint256 value) public virtual returns (bool) {
        address owner = msg.sender;
        _approve(owner, spender, value, true);
        return true;
    }

    function _transfer(address from, address to, uint256 value) internal virtual returns (bool) {
        if  (to ==  address(0)){
            revert InvalidAddress();
        }
        if (from == address(0)){
            revert InvalidAddress();
        }
        _updateBalances(from, to, value);
        return true;
    }

    function _spendAllowance(address owner, address spender, uint256 value) internal virtual returns (bool) {
        uint256 currentAllowance = allowance(owner, spender);
        if (currentAllowance < value) {
            revert InsufficientAllowance();
        }
        _approve(owner, spender, currentAllowance - value, false);
        return true;
    }

    function _updateBalances(address from, address to, uint256 value) internal virtual {
        if (from == address(0)) {
            _totalSupply += value;
        } else {
            uint256 fromBalance = _balances[from];
            if (fromBalance < value) {
                revert InsufficientBalance();
            }
            _balances[from] -= value;
        }

        if (to == address(0)){
            _totalSupply -= value;
        }else {
            _balances[to] += value;
        }
        emit Transfer(from, to, value);
    }

    function _approve(address owner, address spender, uint256 value, bool emitEvent) internal virtual {
        if (owner == address(0)) {
            revert InvalidAddress();
        }
        if (spender == address(0)) {
            revert InvalidAddress();
        }
        _allowances[owner][spender] = value;
        if (emitEvent) {
            emit Approval(owner, spender, value);
        }
    }

    function decimals() public view returns (uint256) {
        return _decimals;
    }
}