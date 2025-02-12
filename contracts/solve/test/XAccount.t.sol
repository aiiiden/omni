// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Test } from "forge-std/Test.sol";
import { Vm } from "forge-std/Vm.sol";
import { XAccount } from "../src/eip7702/XAccount.sol";

contract XAccountTest is Test {
    function test_execute() public {
        address executor = makeAddr("executor");
        Receiver receiver = new Receiver();
        XAccount xaccount = new XAccount(executor);

        bytes32 orderId = keccak256(abi.encodePacked("orderId"));

        XAccount.Call[] memory calls = new XAccount.Call[](1);
        calls[0] = XAccount.Call({ to: address(receiver), data: bytes(""), value: 0.5 ether });

        (address user, uint256 pk) = makeAddrAndKey("user");

        vm.signAndAttachDelegation(address(xaccount), pk);

        bytes32 digest = XAccount(user).authCallsDigest(orderId, calls);

        bytes memory sig = _sign(digest, pk);

        vm.deal(executor, 1 ether);
        vm.prank(executor);
        XAccount(user).execute{ value: 1 ether }(orderId, calls, sig);

        assertEq(address(receiver).balance, 0.5 ether);
        assertEq(user.balance, 0.5 ether);
    }

    function _sign(bytes32 digest, uint256 pk) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        return abi.encodePacked(r, s, v);
    }
}

contract Receiver {
    fallback() external payable { }
    receive() external payable { }
}
