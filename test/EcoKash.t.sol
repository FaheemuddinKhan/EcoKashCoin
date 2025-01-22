// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {EcoKash} from "../src/EcoKash.sol";

contract EcoKashTest is Test {
    EcoKash public ecokash;
    address private owner;
    address private addr1;

    function setUp() public {
        owner = address(this);
        ecokash = new EcoKash(1000);
    }

    function testTransfer() public {
        ecokash.transfer(addr1, 100);
        assertEq(ecokash.balanceOf(addr1), 100);
        assertEq(ecokash.balanceOf(owner), 900);
        assertEq(ecokash.totalSupply(), 1000);
    }





}
