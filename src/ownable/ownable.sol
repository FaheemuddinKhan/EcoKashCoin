// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
import {Error} from "../errors/error.sol";

abstract contract Ownable is Error {
    address private _owner;

    event OwnereshipTransferred(address indexed from, address indexed to);

    error IsNotOwner();
  

    constructor(){
        _owner = msg.sender;
    }

    function owner() external view  returns (address){
        return _owner;
    }

    modifier onlyOwner {
        if (msg.sender != _owner) {
            revert IsNotOwner();
        } _;
    }

    function transferOwnership(address newOwner) external onlyOwner {
        if (newOwner != address(0)) {
            revert InvalidAddress();
        }
        address previousOwner = _owner;
        _owner = newOwner;
        emit OwnereshipTransferred(previousOwner, newOwner);
    }
}