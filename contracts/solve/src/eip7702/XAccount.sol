// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

// import { SignatureCheckerLib } from "solady/src/utils/SignatureCheckerLib.sol";
import { SignatureChecker } from "@openzeppelin/contracts/utils/cryptography/SignatureChecker.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";

contract XAccount is Initializable {
    error CallFailed();
    error AlreadySpent();
    error InvalidSignature();
    error Unauthorized();

    struct Call {
        address to;
        bytes data;
        uint256 value;
    }

    address public immutable executor;

    constructor(address _executor) {
        executor = _executor;
        _disableInitializers();
    }

    struct Storage {
        mapping(bytes32 orderId => bool) spent;
    }

    function _getStorage() internal pure returns (Storage storage $) {
        bytes32 slot = keccak256("omni.xaccount.storage");
        assembly ("memory-safe") {
            $.slot := slot
        }
    }

    function execute(bytes32 orderId, Call[] calldata calls, bytes calldata signature) external payable {
        if (msg.sender != executor) revert Unauthorized();

        Storage storage $ = _getStorage();

        if ($.spent[orderId]) revert AlreadySpent();
        $.spent[orderId] = true;

        // for some reason this does not work
        // if (!SignatureCheckerLib.isValidSignatureNow(address(this), authCallsDigest(orderId, calls), signature)) {
        //     revert InvalidSignature();
        // }

        // oz version does
        if (SignatureChecker.isValidSignatureNow(address(this), authCallsDigest(orderId, calls), signature)) {
            revert InvalidSignature();
        }

        for (uint256 i = 0; i < calls.length; i++) {
            Call memory call = calls[i];
            (bool success,) = call.to.call{ value: call.value }(call.data);
            if (!success) revert CallFailed();
        }
    }

    function authCallsDigest(bytes32 orderId, Call[] calldata calls) public view returns (bytes32) {
        return keccak256(abi.encode(orderId, calls, block.chainid, address(this)));
    }
}
