// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// XAccountCall is an auto generated low-level Go binding around an user-defined struct.
type XAccountCall struct {
	To    common.Address
	Data  []byte
	Value *big.Int
}

// XAccountMetaData contains all meta data concerning the XAccount contract.
var XAccountMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_executor\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authCallsDigest\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structXAccount.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structXAccount.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadySpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610b05380380610b0583398101604081905261002f916100b0565b6001600160a01b03811660805261004461004a565b506100e0565b63409feecd19805460018116156100695763f92ee8a96000526004601cfd5b6001600160401b03808260011c146100ab578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b6000602082840312156100c257600080fd5b81516001600160a01b03811681146100d957600080fd5b9392505050565b608051610a05610100600039600081816093015260d80152610a056000f3fe6080604052600436106100345760003560e01c80634ac0fcbe146100395780635b8d81c81461004e578063c34c08e514610081575b600080fd5b61004c6100473660046105b6565b6100cd565b005b34801561005a57600080fd5b5061006e61006936600461065a565b6102c3565b6040519081526020015b60405180910390f35b34801561008d57600080fd5b506100b57f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610078565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610115576040516282b42960e81b815260040160405180910390fd5b60008581527fb5eea60ec80a31e42f46494999054c4330c842046d162a25a4073414687d9a1f602081905260409091205460ff16156101675760405163d4d9b0fd60e01b815260040160405180910390fd5b6000868152602082905260409020805460ff191660011790556101cb3061018f8888886102c3565b85858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506102fe92505050565b156101e957604051638baa579f60e01b815260040160405180910390fd5b60005b848110156102ba576000868683818110610208576102086106a6565b905060200281019061021a91906106bc565b61022390610768565b9050600081600001516001600160a01b03168260400151836020015160405161024c9190610843565b60006040518083038185875af1925050503d8060008114610289576040519150601f19603f3d011682016040523d82523d6000602084013e61028e565b606091505b50509050806102b057604051633204506f60e01b815260040160405180910390fd5b50506001016101ec565b50505050505050565b600083838346306040516020016102de959493929190610855565b6040516020818303038152906040528051906020012090505b9392505050565b6000836001600160a01b03163b6000036103605760008061031f8585610372565b509092509050600081600381111561033957610339610966565b1480156103575750856001600160a01b0316826001600160a01b0316145b925050506102f7565b61036b8484846103bf565b90506102f7565b600080600083516041036103ac5760208401516040850151606086015160001a61039e8882858561049b565b9550955095505050506103b8565b50508151600091506002905b9250925092565b6000806000856001600160a01b031685856040516024016103e192919061097c565b60408051601f198184030181529181526020820180516001600160e01b0316630b135d3f60e11b179052516104169190610843565b600060405180830381855afa9150503d8060008114610451576040519150601f19603f3d011682016040523d82523d6000602084013e610456565b606091505b509150915081801561046a57506020815110155b801561049157508051630b135d3f60e11b9061048f90830160209081019084016109b6565b145b9695505050505050565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156104d65750600091506003905082610560565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561052a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661055657506000925060019150829050610560565b9250600091508190505b9450945094915050565b60008083601f84011261057c57600080fd5b50813567ffffffffffffffff81111561059457600080fd5b6020830191508360208260051b85010111156105af57600080fd5b9250929050565b6000806000806000606086880312156105ce57600080fd5b85359450602086013567ffffffffffffffff808211156105ed57600080fd5b6105f989838a0161056a565b9096509450604088013591508082111561061257600080fd5b818801915088601f83011261062657600080fd5b81358181111561063557600080fd5b89602082850101111561064757600080fd5b9699959850939650602001949392505050565b60008060006040848603121561066f57600080fd5b83359250602084013567ffffffffffffffff81111561068d57600080fd5b6106998682870161056a565b9497909650939450505050565b634e487b7160e01b600052603260045260246000fd5b60008235605e198336030181126106d257600080fd5b9190910192915050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610715576107156106dc565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610744576107446106dc565b604052919050565b80356001600160a01b038116811461076357600080fd5b919050565b60006060823603121561077a57600080fd5b6107826106f2565b61078b8361074c565b815260208084013567ffffffffffffffff808211156107a957600080fd5b9085019036601f8301126107bc57600080fd5b8135818111156107ce576107ce6106dc565b6107e0601f8201601f1916850161071b565b915080825236848285010111156107f657600080fd5b808484018584013760009082018401529183019190915250604092830135928101929092525090565b60005b8381101561083a578181015183820152602001610822565b50506000910152565b600082516106d281846020870161081f565b6000608080830188845260206080818601528188835260a08601905060a08960051b87010192508960005b8a81101561094357878503609f190183528135368d9003605e190181126108a657600080fd5b8c0160606001600160a01b036108bb8361074c565b16875285820135601e198336030181126108d457600080fd5b8201868101903567ffffffffffffffff8111156108f057600080fd5b8036038213156108ff57600080fd5b82888a015280838a015280828b8b013760008982018b0152604093840135938901939093525050601f01601f19169094018501939183019190830190600101610880565b50505050604084018690526001600160a01b038516606085015291506104919050565b634e487b7160e01b600052602160045260246000fd5b82815260406020820152600082518060408401526109a181606085016020870161081f565b601f01601f1916919091016060019392505050565b6000602082840312156109c857600080fd5b505191905056fea2646970667358221220dd5cb7b993b4776ba45f6601f7fb4800ecdd2ee2cd38bb7f18835ea7fe9348b464736f6c63430008180033",
}

// XAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use XAccountMetaData.ABI instead.
var XAccountABI = XAccountMetaData.ABI

// XAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use XAccountMetaData.Bin instead.
var XAccountBin = XAccountMetaData.Bin

// DeployXAccount deploys a new Ethereum contract, binding an instance of XAccount to it.
func DeployXAccount(auth *bind.TransactOpts, backend bind.ContractBackend, _executor common.Address) (common.Address, *types.Transaction, *XAccount, error) {
	parsed, err := XAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(XAccountBin), backend, _executor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XAccount{XAccountCaller: XAccountCaller{contract: contract}, XAccountTransactor: XAccountTransactor{contract: contract}, XAccountFilterer: XAccountFilterer{contract: contract}}, nil
}

// XAccount is an auto generated Go binding around an Ethereum contract.
type XAccount struct {
	XAccountCaller     // Read-only binding to the contract
	XAccountTransactor // Write-only binding to the contract
	XAccountFilterer   // Log filterer for contract events
}

// XAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type XAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XAccountSession struct {
	Contract     *XAccount         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XAccountCallerSession struct {
	Contract *XAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// XAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XAccountTransactorSession struct {
	Contract     *XAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// XAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type XAccountRaw struct {
	Contract *XAccount // Generic contract binding to access the raw methods on
}

// XAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XAccountCallerRaw struct {
	Contract *XAccountCaller // Generic read-only contract binding to access the raw methods on
}

// XAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XAccountTransactorRaw struct {
	Contract *XAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXAccount creates a new instance of XAccount, bound to a specific deployed contract.
func NewXAccount(address common.Address, backend bind.ContractBackend) (*XAccount, error) {
	contract, err := bindXAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XAccount{XAccountCaller: XAccountCaller{contract: contract}, XAccountTransactor: XAccountTransactor{contract: contract}, XAccountFilterer: XAccountFilterer{contract: contract}}, nil
}

// NewXAccountCaller creates a new read-only instance of XAccount, bound to a specific deployed contract.
func NewXAccountCaller(address common.Address, caller bind.ContractCaller) (*XAccountCaller, error) {
	contract, err := bindXAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XAccountCaller{contract: contract}, nil
}

// NewXAccountTransactor creates a new write-only instance of XAccount, bound to a specific deployed contract.
func NewXAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*XAccountTransactor, error) {
	contract, err := bindXAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XAccountTransactor{contract: contract}, nil
}

// NewXAccountFilterer creates a new log filterer instance of XAccount, bound to a specific deployed contract.
func NewXAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*XAccountFilterer, error) {
	contract, err := bindXAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XAccountFilterer{contract: contract}, nil
}

// bindXAccount binds a generic wrapper to an already deployed contract.
func bindXAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := XAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XAccount *XAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XAccount.Contract.XAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XAccount *XAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XAccount.Contract.XAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XAccount *XAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XAccount.Contract.XAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XAccount *XAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XAccount *XAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XAccount *XAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XAccount.Contract.contract.Transact(opts, method, params...)
}

// AuthCallsDigest is a free data retrieval call binding the contract method 0x5b8d81c8.
//
// Solidity: function authCallsDigest(bytes32 orderId, (address,bytes,uint256)[] calls) view returns(bytes32)
func (_XAccount *XAccountCaller) AuthCallsDigest(opts *bind.CallOpts, orderId [32]byte, calls []XAccountCall) ([32]byte, error) {
	var out []interface{}
	err := _XAccount.contract.Call(opts, &out, "authCallsDigest", orderId, calls)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AuthCallsDigest is a free data retrieval call binding the contract method 0x5b8d81c8.
//
// Solidity: function authCallsDigest(bytes32 orderId, (address,bytes,uint256)[] calls) view returns(bytes32)
func (_XAccount *XAccountSession) AuthCallsDigest(orderId [32]byte, calls []XAccountCall) ([32]byte, error) {
	return _XAccount.Contract.AuthCallsDigest(&_XAccount.CallOpts, orderId, calls)
}

// AuthCallsDigest is a free data retrieval call binding the contract method 0x5b8d81c8.
//
// Solidity: function authCallsDigest(bytes32 orderId, (address,bytes,uint256)[] calls) view returns(bytes32)
func (_XAccount *XAccountCallerSession) AuthCallsDigest(orderId [32]byte, calls []XAccountCall) ([32]byte, error) {
	return _XAccount.Contract.AuthCallsDigest(&_XAccount.CallOpts, orderId, calls)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_XAccount *XAccountCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XAccount.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_XAccount *XAccountSession) Executor() (common.Address, error) {
	return _XAccount.Contract.Executor(&_XAccount.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_XAccount *XAccountCallerSession) Executor() (common.Address, error) {
	return _XAccount.Contract.Executor(&_XAccount.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x4ac0fcbe.
//
// Solidity: function execute(bytes32 orderId, (address,bytes,uint256)[] calls, bytes signature) payable returns()
func (_XAccount *XAccountTransactor) Execute(opts *bind.TransactOpts, orderId [32]byte, calls []XAccountCall, signature []byte) (*types.Transaction, error) {
	return _XAccount.contract.Transact(opts, "execute", orderId, calls, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x4ac0fcbe.
//
// Solidity: function execute(bytes32 orderId, (address,bytes,uint256)[] calls, bytes signature) payable returns()
func (_XAccount *XAccountSession) Execute(orderId [32]byte, calls []XAccountCall, signature []byte) (*types.Transaction, error) {
	return _XAccount.Contract.Execute(&_XAccount.TransactOpts, orderId, calls, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x4ac0fcbe.
//
// Solidity: function execute(bytes32 orderId, (address,bytes,uint256)[] calls, bytes signature) payable returns()
func (_XAccount *XAccountTransactorSession) Execute(orderId [32]byte, calls []XAccountCall, signature []byte) (*types.Transaction, error) {
	return _XAccount.Contract.Execute(&_XAccount.TransactOpts, orderId, calls, signature)
}

// XAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the XAccount contract.
type XAccountInitializedIterator struct {
	Event *XAccountInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XAccountInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XAccountInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XAccountInitialized represents a Initialized event raised by the XAccount contract.
type XAccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_XAccount *XAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*XAccountInitializedIterator, error) {

	logs, sub, err := _XAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &XAccountInitializedIterator{contract: _XAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_XAccount *XAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *XAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _XAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XAccountInitialized)
				if err := _XAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_XAccount *XAccountFilterer) ParseInitialized(log types.Log) (*XAccountInitialized, error) {
	event := new(XAccountInitialized)
	if err := _XAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
