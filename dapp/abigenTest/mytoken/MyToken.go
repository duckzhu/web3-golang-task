// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mytoken

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

// MytokenMetaData contains all meta data concerning the Mytoken contract.
var MytokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testFun\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50336040518060400160405280600981526020017f54657374546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600581526020017f54534d544b000000000000000000000000000000000000000000000000000000815250816003908161008c9190610424565b50806004908161009c9190610424565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361010f575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016101069190610532565b60405180910390fd5b61011e8161012460201b60201c565b5061054b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061026257607f821691505b6020821081036102755761027461021e565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026102d77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261029c565b6102e1868361029c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61032561032061031b846102f9565b610302565b6102f9565b9050919050565b5f819050919050565b61033e8361030b565b61035261034a8261032c565b8484546102a8565b825550505050565b5f5f905090565b61036961035a565b610374818484610335565b505050565b5b818110156103975761038c5f82610361565b60018101905061037a565b5050565b601f8211156103dc576103ad8161027b565b6103b68461028d565b810160208510156103c5578190505b6103d96103d18561028d565b830182610379565b50505b505050565b5f82821c905092915050565b5f6103fc5f19846008026103e1565b1980831691505092915050565b5f61041483836103ed565b9150826002028217905092915050565b61042d826101e7565b67ffffffffffffffff811115610446576104456101f1565b5b610450825461024b565b61045b82828561039b565b5f60209050601f83116001811461048c575f841561047a578287015190505b6104848582610409565b8655506104eb565b601f19841661049a8661027b565b5f5b828110156104c15784890151825560018201915060208501945060208101905061049c565b868310156104de57848901516104da601f8916826103ed565b8355505b6001600288020188555050505b505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61051c826104f3565b9050919050565b61052c81610512565b82525050565b5f6020820190506105455f830184610523565b92915050565b6112f5806105585f395ff3fe608060405234801561000f575f5ffd5b5060043610610109575f3560e01c8063715018a6116100a05780638da5cb5b1161006f5780638da5cb5b1461029157806395d89b41146102af578063a9059cbb146102cd578063dd62ed3e146102fd578063f2fde38b1461032d57610109565b8063715018a61461022f57806379cc679014610239578063855f302c14610255578063893d20e81461027357610109565b8063313ce567116100dc578063313ce567146101a957806340c10f19146101c757806342966c68146101e357806370a08231146101ff57610109565b806306fdde031461010d578063095ea7b31461012b57806318160ddd1461015b57806323b872dd14610179575b5f5ffd5b610115610349565b6040516101229190610f43565b60405180910390f35b61014560048036038101906101409190610ff4565b6103d9565b604051610152919061104c565b60405180910390f35b6101636103fb565b6040516101709190611074565b60405180910390f35b610193600480360381019061018e919061108d565b610404565b6040516101a0919061104c565b60405180910390f35b6101b1610432565b6040516101be91906110f8565b60405180910390f35b6101e160048036038101906101dc9190610ff4565b61043a565b005b6101fd60048036038101906101f89190611111565b610450565b005b6102196004803603810190610214919061113c565b610464565b6040516102269190611074565b60405180910390f35b6102376104a9565b005b610253600480360381019061024e9190610ff4565b6104bc565b005b61025d6104dc565b60405161026a9190610f43565b60405180910390f35b61027b610519565b6040516102889190611176565b60405180910390f35b610299610527565b6040516102a69190611176565b60405180910390f35b6102b761054f565b6040516102c49190610f43565b60405180910390f35b6102e760048036038101906102e29190610ff4565b6105df565b6040516102f4919061104c565b60405180910390f35b6103176004803603810190610312919061118f565b610601565b6040516103249190611074565b60405180910390f35b6103476004803603810190610342919061113c565b610683565b005b606060038054610358906111fa565b80601f0160208091040260200160405190810160405280929190818152602001828054610384906111fa565b80156103cf5780601f106103a6576101008083540402835291602001916103cf565b820191905f5260205f20905b8154815290600101906020018083116103b257829003601f168201915b5050505050905090565b5f5f6103e3610707565b90506103f081858561070e565b600191505092915050565b5f600254905090565b5f5f61040e610707565b905061041b858285610720565b6104268585856107b3565b60019150509392505050565b5f6012905090565b6104426108a3565b61044c828261092a565b5050565b61046161045b610707565b826109a9565b50565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6104b16108a3565b6104ba5f610a28565b565b6104ce826104c8610707565b83610720565b6104d882826109a9565b5050565b60606040518060400160405280600b81526020017f7465737473756363657373000000000000000000000000000000000000000000815250905090565b5f610522610527565b905090565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606004805461055e906111fa565b80601f016020809104026020016040519081016040528092919081815260200182805461058a906111fa565b80156105d55780601f106105ac576101008083540402835291602001916105d5565b820191905f5260205f20905b8154815290600101906020018083116105b857829003601f168201915b5050505050905090565b5f5f6105e9610707565b90506105f68185856107b3565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b61068b6108a3565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036106fb575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016106f29190611176565b60405180910390fd5b61070481610a28565b50565b5f33905090565b61071b8383836001610aeb565b505050565b5f61072b8484610601565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156107ad578181101561079e578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016107959392919061122a565b60405180910390fd5b6107ac84848484035f610aeb565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610823575f6040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260040161081a9190611176565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610893575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161088a9190611176565b60405180910390fd5b61089e838383610cba565b505050565b6108ab610707565b73ffffffffffffffffffffffffffffffffffffffff166108c9610527565b73ffffffffffffffffffffffffffffffffffffffff1614610928576108ec610707565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161091f9190611176565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361099a575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016109919190611176565b60405180910390fd5b6109a55f8383610cba565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a19575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610a109190611176565b60405180910390fd5b610a24825f83610cba565b5050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b5b575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610b529190611176565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610bcb575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610bc29190611176565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610cb4578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610cab9190611074565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610d0a578060025f828254610cfe919061128c565b92505081905550610dd8565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610d93578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610d8a9392919061122a565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e1f578060025f8282540392505081905550610e69565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610ec69190611074565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610f1582610ed3565b610f1f8185610edd565b9350610f2f818560208601610eed565b610f3881610efb565b840191505092915050565b5f6020820190508181035f830152610f5b8184610f0b565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f9082610f67565b9050919050565b610fa081610f86565b8114610faa575f5ffd5b50565b5f81359050610fbb81610f97565b92915050565b5f819050919050565b610fd381610fc1565b8114610fdd575f5ffd5b50565b5f81359050610fee81610fca565b92915050565b5f5f6040838503121561100a57611009610f63565b5b5f61101785828601610fad565b925050602061102885828601610fe0565b9150509250929050565b5f8115159050919050565b61104681611032565b82525050565b5f60208201905061105f5f83018461103d565b92915050565b61106e81610fc1565b82525050565b5f6020820190506110875f830184611065565b92915050565b5f5f5f606084860312156110a4576110a3610f63565b5b5f6110b186828701610fad565b93505060206110c286828701610fad565b92505060406110d386828701610fe0565b9150509250925092565b5f60ff82169050919050565b6110f2816110dd565b82525050565b5f60208201905061110b5f8301846110e9565b92915050565b5f6020828403121561112657611125610f63565b5b5f61113384828501610fe0565b91505092915050565b5f6020828403121561115157611150610f63565b5b5f61115e84828501610fad565b91505092915050565b61117081610f86565b82525050565b5f6020820190506111895f830184611167565b92915050565b5f5f604083850312156111a5576111a4610f63565b5b5f6111b285828601610fad565b92505060206111c385828601610fad565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061121157607f821691505b602082108103611224576112236111cd565b5b50919050565b5f60608201905061123d5f830186611167565b61124a6020830185611065565b6112576040830184611065565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61129682610fc1565b91506112a183610fc1565b92508282019050808211156112b9576112b861125f565b5b9291505056fea264697066735822122001e189005b90ded645cc0044c2e508d63c3ff2de1baf688d4ed3b62d5ad0865a64736f6c634300081e0033",
}

// MytokenABI is the input ABI used to generate the binding from.
// Deprecated: Use MytokenMetaData.ABI instead.
var MytokenABI = MytokenMetaData.ABI

// MytokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MytokenMetaData.Bin instead.
var MytokenBin = MytokenMetaData.Bin

// DeployMytoken deploys a new Ethereum contract, binding an instance of Mytoken to it.
func DeployMytoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mytoken, error) {
	parsed, err := MytokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MytokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mytoken{MytokenCaller: MytokenCaller{contract: contract}, MytokenTransactor: MytokenTransactor{contract: contract}, MytokenFilterer: MytokenFilterer{contract: contract}}, nil
}

// Mytoken is an auto generated Go binding around an Ethereum contract.
type Mytoken struct {
	MytokenCaller     // Read-only binding to the contract
	MytokenTransactor // Write-only binding to the contract
	MytokenFilterer   // Log filterer for contract events
}

// MytokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MytokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MytokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MytokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MytokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MytokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MytokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MytokenSession struct {
	Contract     *Mytoken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MytokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MytokenCallerSession struct {
	Contract *MytokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MytokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MytokenTransactorSession struct {
	Contract     *MytokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MytokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MytokenRaw struct {
	Contract *Mytoken // Generic contract binding to access the raw methods on
}

// MytokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MytokenCallerRaw struct {
	Contract *MytokenCaller // Generic read-only contract binding to access the raw methods on
}

// MytokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MytokenTransactorRaw struct {
	Contract *MytokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMytoken creates a new instance of Mytoken, bound to a specific deployed contract.
func NewMytoken(address common.Address, backend bind.ContractBackend) (*Mytoken, error) {
	contract, err := bindMytoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mytoken{MytokenCaller: MytokenCaller{contract: contract}, MytokenTransactor: MytokenTransactor{contract: contract}, MytokenFilterer: MytokenFilterer{contract: contract}}, nil
}

// NewMytokenCaller creates a new read-only instance of Mytoken, bound to a specific deployed contract.
func NewMytokenCaller(address common.Address, caller bind.ContractCaller) (*MytokenCaller, error) {
	contract, err := bindMytoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MytokenCaller{contract: contract}, nil
}

// NewMytokenTransactor creates a new write-only instance of Mytoken, bound to a specific deployed contract.
func NewMytokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MytokenTransactor, error) {
	contract, err := bindMytoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MytokenTransactor{contract: contract}, nil
}

// NewMytokenFilterer creates a new log filterer instance of Mytoken, bound to a specific deployed contract.
func NewMytokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MytokenFilterer, error) {
	contract, err := bindMytoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MytokenFilterer{contract: contract}, nil
}

// bindMytoken binds a generic wrapper to an already deployed contract.
func bindMytoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MytokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mytoken *MytokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mytoken.Contract.MytokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mytoken *MytokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mytoken.Contract.MytokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mytoken *MytokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mytoken.Contract.MytokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mytoken *MytokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mytoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mytoken *MytokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mytoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mytoken *MytokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mytoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mytoken *MytokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mytoken *MytokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mytoken.Contract.Allowance(&_Mytoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mytoken *MytokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mytoken.Contract.Allowance(&_Mytoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mytoken *MytokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mytoken *MytokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mytoken.Contract.BalanceOf(&_Mytoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mytoken *MytokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mytoken.Contract.BalanceOf(&_Mytoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mytoken *MytokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mytoken *MytokenSession) Decimals() (uint8, error) {
	return _Mytoken.Contract.Decimals(&_Mytoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mytoken *MytokenCallerSession) Decimals() (uint8, error) {
	return _Mytoken.Contract.Decimals(&_Mytoken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Mytoken *MytokenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Mytoken *MytokenSession) GetOwner() (common.Address, error) {
	return _Mytoken.Contract.GetOwner(&_Mytoken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Mytoken *MytokenCallerSession) GetOwner() (common.Address, error) {
	return _Mytoken.Contract.GetOwner(&_Mytoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mytoken *MytokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mytoken *MytokenSession) Name() (string, error) {
	return _Mytoken.Contract.Name(&_Mytoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mytoken *MytokenCallerSession) Name() (string, error) {
	return _Mytoken.Contract.Name(&_Mytoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mytoken *MytokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mytoken *MytokenSession) Owner() (common.Address, error) {
	return _Mytoken.Contract.Owner(&_Mytoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mytoken *MytokenCallerSession) Owner() (common.Address, error) {
	return _Mytoken.Contract.Owner(&_Mytoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mytoken *MytokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mytoken *MytokenSession) Symbol() (string, error) {
	return _Mytoken.Contract.Symbol(&_Mytoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mytoken *MytokenCallerSession) Symbol() (string, error) {
	return _Mytoken.Contract.Symbol(&_Mytoken.CallOpts)
}

// TestFun is a free data retrieval call binding the contract method 0x855f302c.
//
// Solidity: function testFun() pure returns(string)
func (_Mytoken *MytokenCaller) TestFun(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "testFun")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TestFun is a free data retrieval call binding the contract method 0x855f302c.
//
// Solidity: function testFun() pure returns(string)
func (_Mytoken *MytokenSession) TestFun() (string, error) {
	return _Mytoken.Contract.TestFun(&_Mytoken.CallOpts)
}

// TestFun is a free data retrieval call binding the contract method 0x855f302c.
//
// Solidity: function testFun() pure returns(string)
func (_Mytoken *MytokenCallerSession) TestFun() (string, error) {
	return _Mytoken.Contract.TestFun(&_Mytoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mytoken *MytokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mytoken *MytokenSession) TotalSupply() (*big.Int, error) {
	return _Mytoken.Contract.TotalSupply(&_Mytoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mytoken *MytokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Mytoken.Contract.TotalSupply(&_Mytoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Mytoken *MytokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Mytoken *MytokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Approve(&_Mytoken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Mytoken *MytokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Approve(&_Mytoken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_Mytoken *MytokenTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_Mytoken *MytokenSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Burn(&_Mytoken.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_Mytoken *MytokenTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Burn(&_Mytoken.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_Mytoken *MytokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_Mytoken *MytokenSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.BurnFrom(&_Mytoken.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_Mytoken *MytokenTransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.BurnFrom(&_Mytoken.TransactOpts, account, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Mytoken *MytokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Mytoken *MytokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Mint(&_Mytoken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Mytoken *MytokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Mint(&_Mytoken.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mytoken *MytokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mytoken *MytokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mytoken.Contract.RenounceOwnership(&_Mytoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mytoken *MytokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mytoken.Contract.RenounceOwnership(&_Mytoken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Mytoken *MytokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Mytoken *MytokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Transfer(&_Mytoken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Mytoken *MytokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Transfer(&_Mytoken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Mytoken *MytokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Mytoken *MytokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.TransferFrom(&_Mytoken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Mytoken *MytokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.TransferFrom(&_Mytoken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mytoken *MytokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mytoken *MytokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mytoken.Contract.TransferOwnership(&_Mytoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mytoken *MytokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mytoken.Contract.TransferOwnership(&_Mytoken.TransactOpts, newOwner)
}

// MytokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mytoken contract.
type MytokenApprovalIterator struct {
	Event *MytokenApproval // Event containing the contract specifics and raw log

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
func (it *MytokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MytokenApproval)
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
		it.Event = new(MytokenApproval)
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
func (it *MytokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MytokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MytokenApproval represents a Approval event raised by the Mytoken contract.
type MytokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mytoken *MytokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MytokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mytoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MytokenApprovalIterator{contract: _Mytoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mytoken *MytokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MytokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mytoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MytokenApproval)
				if err := _Mytoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Mytoken *MytokenFilterer) ParseApproval(log types.Log) (*MytokenApproval, error) {
	event := new(MytokenApproval)
	if err := _Mytoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MytokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mytoken contract.
type MytokenOwnershipTransferredIterator struct {
	Event *MytokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MytokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MytokenOwnershipTransferred)
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
		it.Event = new(MytokenOwnershipTransferred)
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
func (it *MytokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MytokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MytokenOwnershipTransferred represents a OwnershipTransferred event raised by the Mytoken contract.
type MytokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mytoken *MytokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MytokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mytoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MytokenOwnershipTransferredIterator{contract: _Mytoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mytoken *MytokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MytokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mytoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MytokenOwnershipTransferred)
				if err := _Mytoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mytoken *MytokenFilterer) ParseOwnershipTransferred(log types.Log) (*MytokenOwnershipTransferred, error) {
	event := new(MytokenOwnershipTransferred)
	if err := _Mytoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MytokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mytoken contract.
type MytokenTransferIterator struct {
	Event *MytokenTransfer // Event containing the contract specifics and raw log

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
func (it *MytokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MytokenTransfer)
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
		it.Event = new(MytokenTransfer)
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
func (it *MytokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MytokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MytokenTransfer represents a Transfer event raised by the Mytoken contract.
type MytokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mytoken *MytokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MytokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mytoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MytokenTransferIterator{contract: _Mytoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mytoken *MytokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MytokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mytoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MytokenTransfer)
				if err := _Mytoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Mytoken *MytokenFilterer) ParseTransfer(log types.Log) (*MytokenTransfer, error) {
	event := new(MytokenTransfer)
	if err := _Mytoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
