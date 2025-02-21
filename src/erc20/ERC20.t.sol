// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {ERC20} from "../ERC20/ERC20.sol";
import {Ownable} from "../ownable/ownable.sol";
import "../errors/error.sol";
import "forge-std/Test.sol";

contract MockERC20 is ERC20 {
    constructor(string memory name, string memory symbol, uint256 initialSupply)
        ERC20(name, symbol, initialSupply) {}
}


contract ERC20Test is Test {
    MockERC20 public token;
    address public owner;
    address public addr1;
    address public addr2;

    uint256 public initialSupply = 100000000;
    uint256 public transferAmount = 100;

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1);
        addr2 = address(0x2);

        // Deploy the MockERC20 token
        token = new MockERC20("Test Token", "TST", initialSupply);
    }

    function testInitialSupply() public view {
        assertEq(token.totalSupply(), initialSupply, "Initial supply should be correct");
    }

    function testBalanceOf() public view {
        assertEq(token.balanceOf(owner), initialSupply, "Owner should have the total supply initially");
    }

    function testTransfer() public {
        token.transfer(addr1, transferAmount);
        assertEq(token.balanceOf(addr1), transferAmount, "Addr1 should receive transferred amount");
        assertEq(token.balanceOf(owner), initialSupply - transferAmount, "Owner's balance should be reduced correctly");
    }

    function testTransferToZeroAddress() public {
        vm.expectRevert(InvalidAddress.selector);
        token.transfer(address(0), transferAmount);
    }

    function testAllowance() public  {
        token.approve(addr1, 100);
        assertEq(token.allowance(owner, addr1), 100);
    }

    function testTransferFrom() public {
        uint256 amount = 100;
        token.approve(addr1, amount);
        console.log("Allowance before transfer: ", token.allowance(owner, addr1));

        vm.prank(addr1);
        console.log("Sender before transfer: ", msg.sender);
        token.transferFrom(owner, addr2, amount);

        // Assertions
        assertEq(token.balanceOf(addr2), amount, "Addr2 should receive the transfer from Addr1");
        assertEq(token.balanceOf(owner), initialSupply - amount, "Owner's balance should be reduced correctly");
    }


    function testTransferFromInsufficientAllowance() public {
        vm.expectRevert(ERC20.InsufficientAllowance.selector);
        vm.prank(addr1);
        token.transferFrom(owner, addr2, transferAmount); // addr1 has no allowance to spend
    }

    function testApprove() public {
        token.approve(addr1, transferAmount);
        assertEq(token.allowance(owner, addr1), transferAmount, "Allowance should be set correctly");
    }

    function testApproveToZeroAddress() public {
        vm.expectRevert(InvalidAddress.selector);
        token.approve(address(0), transferAmount);
    }

    function testExceedsInitialSupplyThreshold() public {
        uint256 largeSupply = 2000 * 10 ** 18; // Greater than 1000 * 10^18
        vm.expectRevert(ERC20.ExceedsInitialSupplyThreshold.selector);
        new MockERC20("Large Supply Token", "LST", largeSupply);
    }

    function testInsufficientBalance() public {
        uint256 excessiveAmount = initialSupply + 1;
        vm.expectRevert(ERC20.InsufficientBalance.selector);
        token.transfer(addr1, excessiveAmount); // Trying to transfer more than the owner's balance
    }
}
