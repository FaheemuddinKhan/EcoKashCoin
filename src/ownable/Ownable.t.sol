// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "./Ownable.sol";

// Mock contract to instantiate Ownable
contract OwnableMock is Ownable {}

contract OwnableTest is Test {
    OwnableMock ownable;
    address owner;
    address addr1;

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1);
        ownable = new OwnableMock(); // Now we can instantiate it
    }

    function testInitialOwner() public view {
        assertEq(ownable.owner(), owner, "Owner should be deployer");
    }

    function testTransferOwnership() public {
        ownable.transferOwnership(addr1);
        assertEq(ownable.owner(), addr1, "Ownership transfer failed");
    }

    function testOnlyOwnerCanTransfer() public {
        vm.prank(addr1); // Simulates a call from addr1
        vm.expectRevert(Ownable.IsNotOwner.selector);
        ownable.transferOwnership(owner);
    }

    function testTransferToZeroAddress() public {
        vm.expectRevert(InvalidAddress.selector);
        ownable.transferOwnership(address(0));
    }
}
