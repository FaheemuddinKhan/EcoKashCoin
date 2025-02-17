// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

// For building our own token (currency in this case) we will use ERC20 standard
// https://eips.ethereum.org/EIPS/eip-20 
interface IERC20 {
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