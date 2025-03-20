// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecokash

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

// EcokashMetaData contains all meta data concerning the Ecokash contract.
var EcokashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExceedsInitialSupplyThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IsNotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnereshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600860065f6101000a81548160ff021916908360ff16021790555060065f9054906101000a900460ff16600a61003a919061059d565b60ff166103e861004a91906105d5565b61ffff16600660016101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550348015610091575f5ffd5b5060405161173038038061173083398181016040528101906100b39190610648565b6040518060400160405280600881526020017f45636f204b6173680000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f454b43000000000000000000000000000000000000000000000000000000000081525082335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826004908161016e91906108a7565b50816005908161017e91906108a7565b50600660019054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681116101cc576101c75f338361020760201b60201c565b6101fe565b6040517fed7433b900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050610a04565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610257578060015f82825461024b9190610976565b92505081905550610327565b5f60025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050818110156102d2576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461031e91906109a9565b92505081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610377578060015f82825461036b91906109a9565b925050819055506103cb565b8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546103c39190610976565b925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161042891906109eb565b60405180910390a3505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b60018511156104b75780860481111561049357610492610435565b5b60018516156104a25780820291505b80810290506104b085610462565b9450610477565b94509492505050565b5f826104cf576001905061058a565b816104dc575f905061058a565b81600181146104f257600281146104fc5761052b565b600191505061058a565b60ff84111561050e5761050d610435565b5b8360020a91508482111561052557610524610435565b5b5061058a565b5060208310610133831016604e8410600b84101617156105605782820a90508381111561055b5761055a610435565b5b61058a565b61056d848484600161046e565b9250905081840481111561058457610583610435565b5b81810290505b9392505050565b5f60ff82169050919050565b5f6105a782610591565b91506105b283610591565b92506105c060ff84846104c0565b905092915050565b5f61ffff82169050919050565b5f6105df826105c8565b91506105ea836105c8565b92508282026105f8816105c8565b915080821461060a57610609610435565b5b5092915050565b5f5ffd5b5f819050919050565b61062781610615565b8114610631575f5ffd5b50565b5f815190506106428161061e565b92915050565b5f6020828403121561065d5761065c610611565b5b5f61066a84828501610634565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806106ee57607f821691505b602082108103610701576107006106aa565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026107637fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610728565b61076d8683610728565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6107a86107a361079e84610615565b610785565b610615565b9050919050565b5f819050919050565b6107c18361078e565b6107d56107cd826107af565b848454610734565b825550505050565b5f5f905090565b6107ec6107dd565b6107f78184846107b8565b505050565b5b8181101561081a5761080f5f826107e4565b6001810190506107fd565b5050565b601f82111561085f5761083081610707565b61083984610719565b81016020851015610848578190505b61085c61085485610719565b8301826107fc565b50505b505050565b5f82821c905092915050565b5f61087f5f1984600802610864565b1980831691505092915050565b5f6108978383610870565b9150826002028217905092915050565b6108b082610673565b67ffffffffffffffff8111156108c9576108c861067d565b5b6108d382546106d7565b6108de82828561081e565b5f60209050601f83116001811461090f575f84156108fd578287015190505b610907858261088c565b86555061096e565b601f19841661091d86610707565b5f5b828110156109445784890151825560018201915060208501945060208101905061091f565b86831015610961578489015161095d601f891682610870565b8355505b6001600288020188555050505b505050505050565b5f61098082610615565b915061098b83610615565b92508282019050808211156109a3576109a2610435565b5b92915050565b5f6109b382610615565b91506109be83610615565b92508282039050818111156109d6576109d5610435565b5b92915050565b6109e581610615565b82525050565b5f6020820190506109fe5f8301846109dc565b92915050565b610d1f80610a115f395ff3fe608060405234801561000f575f5ffd5b5060043610610091575f3560e01c806370a082311161006457806370a08231146101315780638da5cb5b14610161578063a9059cbb1461017f578063dd62ed3e146101af578063f2fde38b146101df57610091565b8063095ea7b31461009557806318160ddd146100c557806323b872dd146100e3578063313ce56714610113575b5f5ffd5b6100af60048036038101906100aa9190610adc565b6101fb565b6040516100bc9190610b34565b60405180910390f35b6100cd610218565b6040516100da9190610b5c565b60405180910390f35b6100fd60048036038101906100f89190610b75565b610221565b60405161010a9190610b34565b60405180910390f35b61011b61024a565b6040516101289190610b5c565b60405180910390f35b61014b60048036038101906101469190610bc5565b610262565b6040516101589190610b5c565b60405180910390f35b6101696102a8565b6040516101769190610bff565b60405180910390f35b61019960048036038101906101949190610adc565b6102cf565b6040516101a69190610b34565b60405180910390f35b6101c960048036038101906101c49190610c18565b6102eb565b6040516101d69190610b5c565b60405180910390f35b6101f960048036038101906101f49190610bc5565b61036d565b005b5f5f33905061020d8185856001610518565b600191505092915050565b5f600154905090565b5f5f3390506102318582856106d1565b5061023d85858561073c565b5060019150509392505050565b5f60065f9054906101000a900460ff1660ff16905090565b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f5f3390506102df81858561073c565b50600191505092915050565b5f60035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103f2576040517f65b023fd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610457576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f499fc83ca1314bbc7cbdcd55d36bbb85930986478f8c210e09fdd709f783d1bc60405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff160361057d576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036105e2576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555080156106cb578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516106c29190610b5c565b60405180910390a35b50505050565b5f5f6106dd85856102eb565b905082811015610719576040517f13be252b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107308585858461072a9190610c83565b5f610518565b60019150509392505050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036107a2576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610807576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61081284848461081d565b600190509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361086d578060015f8282546108619190610cb6565b9250508190555061093d565b5f60025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050818110156108e8576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109349190610c83565b92505081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361098d578060015f8282546109819190610c83565b925050819055506109e1565b8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109d99190610cb6565b925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610a3e9190610b5c565b60405180910390a3505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a7882610a4f565b9050919050565b610a8881610a6e565b8114610a92575f5ffd5b50565b5f81359050610aa381610a7f565b92915050565b5f819050919050565b610abb81610aa9565b8114610ac5575f5ffd5b50565b5f81359050610ad681610ab2565b92915050565b5f5f60408385031215610af257610af1610a4b565b5b5f610aff85828601610a95565b9250506020610b1085828601610ac8565b9150509250929050565b5f8115159050919050565b610b2e81610b1a565b82525050565b5f602082019050610b475f830184610b25565b92915050565b610b5681610aa9565b82525050565b5f602082019050610b6f5f830184610b4d565b92915050565b5f5f5f60608486031215610b8c57610b8b610a4b565b5b5f610b9986828701610a95565b9350506020610baa86828701610a95565b9250506040610bbb86828701610ac8565b9150509250925092565b5f60208284031215610bda57610bd9610a4b565b5b5f610be784828501610a95565b91505092915050565b610bf981610a6e565b82525050565b5f602082019050610c125f830184610bf0565b92915050565b5f5f60408385031215610c2e57610c2d610a4b565b5b5f610c3b85828601610a95565b9250506020610c4c85828601610a95565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610c8d82610aa9565b9150610c9883610aa9565b9250828203905081811115610cb057610caf610c56565b5b92915050565b5f610cc082610aa9565b9150610ccb83610aa9565b9250828201905080821115610ce357610ce2610c56565b5b9291505056fea26469706673582212208a0358e78ef2cff97d858306135a0725081c07ff73e326aea0d18ae2b7bfb95e64736f6c634300081d0033",
}

// EcokashABI is the input ABI used to generate the binding from.
// Deprecated: Use EcokashMetaData.ABI instead.
var EcokashABI = EcokashMetaData.ABI

// EcokashBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EcokashMetaData.Bin instead.
var EcokashBin = EcokashMetaData.Bin

// DeployEcokash deploys a new Ethereum contract, binding an instance of Ecokash to it.
func DeployEcokash(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *Ecokash, error) {
	parsed, err := EcokashMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EcokashBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ecokash{EcokashCaller: EcokashCaller{contract: contract}, EcokashTransactor: EcokashTransactor{contract: contract}, EcokashFilterer: EcokashFilterer{contract: contract}}, nil
}

// Ecokash is an auto generated Go binding around an Ethereum contract.
type Ecokash struct {
	EcokashCaller     // Read-only binding to the contract
	EcokashTransactor // Write-only binding to the contract
	EcokashFilterer   // Log filterer for contract events
}

// EcokashCaller is an auto generated read-only Go binding around an Ethereum contract.
type EcokashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcokashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EcokashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcokashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EcokashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcokashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EcokashSession struct {
	Contract     *Ecokash          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EcokashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EcokashCallerSession struct {
	Contract *EcokashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EcokashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EcokashTransactorSession struct {
	Contract     *EcokashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EcokashRaw is an auto generated low-level Go binding around an Ethereum contract.
type EcokashRaw struct {
	Contract *Ecokash // Generic contract binding to access the raw methods on
}

// EcokashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EcokashCallerRaw struct {
	Contract *EcokashCaller // Generic read-only contract binding to access the raw methods on
}

// EcokashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EcokashTransactorRaw struct {
	Contract *EcokashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcokash creates a new instance of Ecokash, bound to a specific deployed contract.
func NewEcokash(address common.Address, backend bind.ContractBackend) (*Ecokash, error) {
	contract, err := bindEcokash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ecokash{EcokashCaller: EcokashCaller{contract: contract}, EcokashTransactor: EcokashTransactor{contract: contract}, EcokashFilterer: EcokashFilterer{contract: contract}}, nil
}

// NewEcokashCaller creates a new read-only instance of Ecokash, bound to a specific deployed contract.
func NewEcokashCaller(address common.Address, caller bind.ContractCaller) (*EcokashCaller, error) {
	contract, err := bindEcokash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EcokashCaller{contract: contract}, nil
}

// NewEcokashTransactor creates a new write-only instance of Ecokash, bound to a specific deployed contract.
func NewEcokashTransactor(address common.Address, transactor bind.ContractTransactor) (*EcokashTransactor, error) {
	contract, err := bindEcokash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EcokashTransactor{contract: contract}, nil
}

// NewEcokashFilterer creates a new log filterer instance of Ecokash, bound to a specific deployed contract.
func NewEcokashFilterer(address common.Address, filterer bind.ContractFilterer) (*EcokashFilterer, error) {
	contract, err := bindEcokash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EcokashFilterer{contract: contract}, nil
}

// bindEcokash binds a generic wrapper to an already deployed contract.
func bindEcokash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EcokashMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecokash *EcokashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecokash.Contract.EcokashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecokash *EcokashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecokash.Contract.EcokashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecokash *EcokashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecokash.Contract.EcokashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecokash *EcokashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecokash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecokash *EcokashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecokash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecokash *EcokashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecokash.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ecokash *EcokashCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ecokash.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ecokash *EcokashSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ecokash.Contract.Allowance(&_Ecokash.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ecokash *EcokashCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ecokash.Contract.Allowance(&_Ecokash.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ecokash *EcokashCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ecokash.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ecokash *EcokashSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ecokash.Contract.BalanceOf(&_Ecokash.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ecokash *EcokashCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ecokash.Contract.BalanceOf(&_Ecokash.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Ecokash *EcokashCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ecokash.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Ecokash *EcokashSession) Decimals() (*big.Int, error) {
	return _Ecokash.Contract.Decimals(&_Ecokash.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Ecokash *EcokashCallerSession) Decimals() (*big.Int, error) {
	return _Ecokash.Contract.Decimals(&_Ecokash.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ecokash *EcokashCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecokash.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ecokash *EcokashSession) Owner() (common.Address, error) {
	return _Ecokash.Contract.Owner(&_Ecokash.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ecokash *EcokashCallerSession) Owner() (common.Address, error) {
	return _Ecokash.Contract.Owner(&_Ecokash.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ecokash *EcokashCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ecokash.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ecokash *EcokashSession) TotalSupply() (*big.Int, error) {
	return _Ecokash.Contract.TotalSupply(&_Ecokash.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ecokash *EcokashCallerSession) TotalSupply() (*big.Int, error) {
	return _Ecokash.Contract.TotalSupply(&_Ecokash.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Ecokash *EcokashTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ecokash.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Ecokash *EcokashSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ecokash.Contract.Approve(&_Ecokash.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Ecokash *EcokashTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ecokash.Contract.Approve(&_Ecokash.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Ecokash *EcokashTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ecokash.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Ecokash *EcokashSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ecokash.Contract.Transfer(&_Ecokash.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Ecokash *EcokashTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ecokash.Contract.Transfer(&_Ecokash.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Ecokash *EcokashTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ecokash.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Ecokash *EcokashSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ecokash.Contract.TransferFrom(&_Ecokash.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Ecokash *EcokashTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ecokash.Contract.TransferFrom(&_Ecokash.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ecokash *EcokashTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ecokash.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ecokash *EcokashSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ecokash.Contract.TransferOwnership(&_Ecokash.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ecokash *EcokashTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ecokash.Contract.TransferOwnership(&_Ecokash.TransactOpts, newOwner)
}

// EcokashApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ecokash contract.
type EcokashApprovalIterator struct {
	Event *EcokashApproval // Event containing the contract specifics and raw log

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
func (it *EcokashApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcokashApproval)
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
		it.Event = new(EcokashApproval)
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
func (it *EcokashApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcokashApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcokashApproval represents a Approval event raised by the Ecokash contract.
type EcokashApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ecokash *EcokashFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EcokashApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ecokash.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EcokashApprovalIterator{contract: _Ecokash.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ecokash *EcokashFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EcokashApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ecokash.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcokashApproval)
				if err := _Ecokash.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ecokash *EcokashFilterer) ParseApproval(log types.Log) (*EcokashApproval, error) {
	event := new(EcokashApproval)
	if err := _Ecokash.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcokashOwnereshipTransferredIterator is returned from FilterOwnereshipTransferred and is used to iterate over the raw logs and unpacked data for OwnereshipTransferred events raised by the Ecokash contract.
type EcokashOwnereshipTransferredIterator struct {
	Event *EcokashOwnereshipTransferred // Event containing the contract specifics and raw log

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
func (it *EcokashOwnereshipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcokashOwnereshipTransferred)
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
		it.Event = new(EcokashOwnereshipTransferred)
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
func (it *EcokashOwnereshipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcokashOwnereshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcokashOwnereshipTransferred represents a OwnereshipTransferred event raised by the Ecokash contract.
type EcokashOwnereshipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnereshipTransferred is a free log retrieval operation binding the contract event 0x499fc83ca1314bbc7cbdcd55d36bbb85930986478f8c210e09fdd709f783d1bc.
//
// Solidity: event OwnereshipTransferred(address indexed from, address indexed to)
func (_Ecokash *EcokashFilterer) FilterOwnereshipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EcokashOwnereshipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ecokash.contract.FilterLogs(opts, "OwnereshipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EcokashOwnereshipTransferredIterator{contract: _Ecokash.contract, event: "OwnereshipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnereshipTransferred is a free log subscription operation binding the contract event 0x499fc83ca1314bbc7cbdcd55d36bbb85930986478f8c210e09fdd709f783d1bc.
//
// Solidity: event OwnereshipTransferred(address indexed from, address indexed to)
func (_Ecokash *EcokashFilterer) WatchOwnereshipTransferred(opts *bind.WatchOpts, sink chan<- *EcokashOwnereshipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ecokash.contract.WatchLogs(opts, "OwnereshipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcokashOwnereshipTransferred)
				if err := _Ecokash.contract.UnpackLog(event, "OwnereshipTransferred", log); err != nil {
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

// ParseOwnereshipTransferred is a log parse operation binding the contract event 0x499fc83ca1314bbc7cbdcd55d36bbb85930986478f8c210e09fdd709f783d1bc.
//
// Solidity: event OwnereshipTransferred(address indexed from, address indexed to)
func (_Ecokash *EcokashFilterer) ParseOwnereshipTransferred(log types.Log) (*EcokashOwnereshipTransferred, error) {
	event := new(EcokashOwnereshipTransferred)
	if err := _Ecokash.contract.UnpackLog(event, "OwnereshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcokashTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ecokash contract.
type EcokashTransferIterator struct {
	Event *EcokashTransfer // Event containing the contract specifics and raw log

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
func (it *EcokashTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcokashTransfer)
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
		it.Event = new(EcokashTransfer)
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
func (it *EcokashTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcokashTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcokashTransfer represents a Transfer event raised by the Ecokash contract.
type EcokashTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ecokash *EcokashFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EcokashTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ecokash.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EcokashTransferIterator{contract: _Ecokash.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ecokash *EcokashFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EcokashTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ecokash.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcokashTransfer)
				if err := _Ecokash.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ecokash *EcokashFilterer) ParseTransfer(log types.Log) (*EcokashTransfer, error) {
	event := new(EcokashTransfer)
	if err := _Ecokash.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
