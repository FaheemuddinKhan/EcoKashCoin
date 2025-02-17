// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
import {ERC20} from "../src/erc20/ERC20.sol";


contract EcoKash is ERC20 {
    constructor(uint initialSupply) ERC20("Eco Kash","EKC", initialSupply) {
    }
}