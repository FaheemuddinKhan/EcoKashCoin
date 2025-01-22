// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {EcoKash} from "../src/EcoKash.sol";

contract CounterScript is Script {
    EcoKash public ecokash;

    function setUp() public {
        ecokash = new EcoKash(1000);
        console.log(address(this));
    }
    function run() external{
        setUp();
        // Example: Interact with contract (e.g., checking balance)
        uint256 balance = ecokash.balanceOf(address(this));
        console.log("Balance of contract deployer:", balance);

        ecokash.transfer(address(0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65), 100);
        balance = ecokash.balanceOf(address(this));
        console.log("Balance of contract deployer:", balance);
    }
}
