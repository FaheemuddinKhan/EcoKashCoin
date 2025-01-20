// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {EcoKash} from "../src/EcoKash.sol";

contract EcoKashTest is Test {
    EcoKash public ecokash;

    function setUp() public {
        ecokash = new EcoKash();
        ecokash.setNumber(0);
    }


}
