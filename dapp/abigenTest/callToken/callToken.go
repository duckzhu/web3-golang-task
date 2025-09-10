package callToken

import (
	"abigenTest/mytoken"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// remix部署好的一个新的erc20合约  地址 0x11B8847A9219dEfce6AcbC2ec3B151641d1cE500
// abi
// const erc20abi = `[
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "spender",
// 				"type": "address"
// 			},
// 			{
// 				"internalType": "uint256",
// 				"name": "value",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "approve",
// 		"outputs": [
// 			{
// 				"internalType": "bool",
// 				"name": "",
// 				"type": "bool"
// 			}
// 		],
// 		"stateMutability": "nonpayable",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "uint256",
// 				"name": "value",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "burn",
// 		"outputs": [],
// 		"stateMutability": "nonpayable",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "account",
// 				"type": "address"
// 			},
// 			{
// 				"internalType": "uint256",
// 				"name": "value",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "burnFrom",
// 		"outputs": [],
// 		"stateMutability": "nonpayable",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [],
// 		"stateMutability": "nonpayable",
// 		"type": "constructor"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "spender",
// 				"type": "address"
// 			},
// 			{
// 				"internalType": "uint256",
// 				"name": "allowance",
// 				"type": "uint256"
// 			},
// 			{
// 				"internalType": "uint256",
// 				"name": "needed",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "ERC20InsufficientAllowance",
// 		"type": "error"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "sender",
// 				"type": "address"
// 			},
// 			{
// 				"internalType": "uint256",
// 				"name": "balance",
// 				"type": "uint256"
// 			},
// 			{
// 				"internalType": "uint256",
// 				"name": "needed",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "ERC20InsufficientBalance",
// 		"type": "error"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "approver",
// 				"type": "address"
// 			}
// 		],
// 		"name": "ERC20InvalidApprover",
// 		"type": "error"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "receiver",
// 				"type": "address"
// 			}
// 		],
// 		"name": "ERC20InvalidReceiver",
// 		"type": "error"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "sender",
// 				"type": "address"
// 			}
// 		],
// 		"name": "ERC20InvalidSender",
// 		"type": "error"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "spender",
// 				"type": "address"
// 			}
// 		],
// 		"name": "ERC20InvalidSpender",
// 		"type": "error"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "owner",
// 				"type": "address"
// 			}
// 		],
// 		"name": "OwnableInvalidOwner",
// 		"type": "error"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "account",
// 				"type": "address"
// 			}
// 		],
// 		"name": "OwnableUnauthorizedAccount",
// 		"type": "error"
// 	},
// 	{
// 		"anonymous": false,
// 		"inputs": [
// 			{
// 				"indexed": true,
// 				"internalType": "address",
// 				"name": "owner",
// 				"type": "address"
// 			},
// 			{
// 				"indexed": true,
// 				"internalType": "address",
// 				"name": "spender",
// 				"type": "address"
// 			},
// 			{
// 				"indexed": false,
// 				"internalType": "uint256",
// 				"name": "value",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "Approval",
// 		"type": "event"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "to",
// 				"type": "address"
// 			},
// 			{
// 				"internalType": "uint256",
// 				"name": "amount",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "mint",
// 		"outputs": [],
// 		"stateMutability": "nonpayable",
// 		"type": "function"
// 	},
// 	{
// 		"anonymous": false,
// 		"inputs": [
// 			{
// 				"indexed": true,
// 				"internalType": "address",
// 				"name": "previousOwner",
// 				"type": "address"
// 			},
// 			{
// 				"indexed": true,
// 				"internalType": "address",
// 				"name": "newOwner",
// 				"type": "address"
// 			}
// 		],
// 		"name": "OwnershipTransferred",
// 		"type": "event"
// 	},
// 	{
// 		"inputs": [],
// 		"name": "renounceOwnership",
// 		"outputs": [],
// 		"stateMutability": "nonpayable",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "to",
// 				"type": "address"
// 			},
// 			{
// 				"internalType": "uint256",
// 				"name": "value",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "transfer",
// 		"outputs": [
// 			{
// 				"internalType": "bool",
// 				"name": "",
// 				"type": "bool"
// 			}
// 		],
// 		"stateMutability": "nonpayable",
// 		"type": "function"
// 	},
// 	{
// 		"anonymous": false,
// 		"inputs": [
// 			{
// 				"indexed": true,
// 				"internalType": "address",
// 				"name": "from",
// 				"type": "address"
// 			},
// 			{
// 				"indexed": true,
// 				"internalType": "address",
// 				"name": "to",
// 				"type": "address"
// 			},
// 			{
// 				"indexed": false,
// 				"internalType": "uint256",
// 				"name": "value",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "Transfer",
// 		"type": "event"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "from",
// 				"type": "address"
// 			},
// 			{
// 				"internalType": "address",
// 				"name": "to",
// 				"type": "address"
// 			},
// 			{
// 				"internalType": "uint256",
// 				"name": "value",
// 				"type": "uint256"
// 			}
// 		],
// 		"name": "transferFrom",
// 		"outputs": [
// 			{
// 				"internalType": "bool",
// 				"name": "",
// 				"type": "bool"
// 			}
// 		],
// 		"stateMutability": "nonpayable",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "newOwner",
// 				"type": "address"
// 			}
// 		],
// 		"name": "transferOwnership",
// 		"outputs": [],
// 		"stateMutability": "nonpayable",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "owner",
// 				"type": "address"
// 			},
// 			{
// 				"internalType": "address",
// 				"name": "spender",
// 				"type": "address"
// 			}
// 		],
// 		"name": "allowance",
// 		"outputs": [
// 			{
// 				"internalType": "uint256",
// 				"name": "",
// 				"type": "uint256"
// 			}
// 		],
// 		"stateMutability": "view",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [
// 			{
// 				"internalType": "address",
// 				"name": "account",
// 				"type": "address"
// 			}
// 		],
// 		"name": "balanceOf",
// 		"outputs": [
// 			{
// 				"internalType": "uint256",
// 				"name": "",
// 				"type": "uint256"
// 			}
// 		],
// 		"stateMutability": "view",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [],
// 		"name": "decimals",
// 		"outputs": [
// 			{
// 				"internalType": "uint8",
// 				"name": "",
// 				"type": "uint8"
// 			}
// 		],
// 		"stateMutability": "view",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [],
// 		"name": "getOwner",
// 		"outputs": [
// 			{
// 				"internalType": "address",
// 				"name": "",
// 				"type": "address"
// 			}
// 		],
// 		"stateMutability": "view",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [],
// 		"name": "name",
// 		"outputs": [
// 			{
// 				"internalType": "string",
// 				"name": "",
// 				"type": "string"
// 			}
// 		],
// 		"stateMutability": "view",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [],
// 		"name": "owner",
// 		"outputs": [
// 			{
// 				"internalType": "address",
// 				"name": "",
// 				"type": "address"
// 			}
// 		],
// 		"stateMutability": "view",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [],
// 		"name": "symbol",
// 		"outputs": [
// 			{
// 				"internalType": "string",
// 				"name": "",
// 				"type": "string"
// 			}
// 		],
// 		"stateMutability": "view",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [],
// 		"name": "testFun",
// 		"outputs": [
// 			{
// 				"internalType": "string",
// 				"name": "",
// 				"type": "string"
// 			}
// 		],
// 		"stateMutability": "pure",
// 		"type": "function"
// 	},
// 	{
// 		"inputs": [],
// 		"name": "totalSupply",
// 		"outputs": [
// 			{
// 				"internalType": "uint256",
// 				"name": "",
// 				"type": "uint256"
// 			}
// 		],
// 		"stateMutability": "view",
// 		"type": "function"
// 	}
// ]`

var client *ethclient.Client
var err error
var tokenInstance *mytoken.Mytoken

// token owner 0xb66344A6236646DF74Aed82dd64F360B1113591a
// token sender 0xe3269B4B78240f4E66E79a41bA0EfC2002c77CA7 sender的私钥sepolia
const privateKeyHex string = "78b7dbc13704a5ed37cd0f2468b9**44d399df1d9724ad2c9c6e6472b0d8df6481"

var chainID = big.NewInt(11155111) // Sepolia 测试链的 ChainID

func Init() {
	// 1. 连接到以太坊节点
	// 可以是本地节点（如 http://localhost:8545），或者第三方节点如 Infura、Alchemy
	client, err = ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/mWrjcZw9dhDh-**FvHb2YCI")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	//token合约地址
	tokenContractAddressStr := "0x11B8847A9219dEfce6AcbC2ec3B151641d1cE500"
	tokenContractAddress := common.HexToAddress(tokenContractAddressStr)

	// 实例化合约
	tokenInstance, err = mytoken.NewMytoken(tokenContractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

}

func QueryTest() {

	// 查询余额
	addressStr := "0x93dF7b9eaF035DD76c80F814c3db2C52d76Cb107"
	address := common.HexToAddress(addressStr)

	balance, err := tokenInstance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatalf("Failed to call BalanceOf: %v", err)
	}
	fmt.Printf("Token Balance of %s: %s\n", address.Hex(), balance.String())

	// 调用合约的只读方法
	owner, err := tokenInstance.Owner(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to call Owner: %v", err)
	}
	fmt.Printf("Token Owner: %s\n", owner.Hex())

	testResult, err := tokenInstance.TestFun(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to call TestFun: %v", err)
	}
	fmt.Printf("TestFun Result: %s\n", testResult)
}

// 调用合约的状态变更方法（需要发送交易）
func Transfer() {
	// 接收地址，我自己的钱包测试账户3
	toAddressStr := "0x93dF7b9eaF035DD76c80F814c3db2C52d76Cb107"
	toAddress := common.HexToAddress(toAddressStr)

	// 1调用 Transfer 方法（发送代币）
	// 创建 TransactOpts 用于发送交易
	// 发送者账户0xe3269B4B78240f4E66E79a41bA0EfC2002c77CA7的私钥sepolia
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("无效的私钥:", err)
	}
	// abibind 包提供了一个方便的方法来创建 TransactOpts,可以对比下ethclient包的创建交易
	// 区别是 ethclient 需要你手动构建交易数据，而 abibind 会帮你处理这些细节
	// 你需要传入私钥和 chainID 来创建 TransactOpts
	// TransactOpts 包含了发送交易所需的信息，比如签名者、gas 限制、gas 价格等
	// 这里的 chainID 是为了防止重放攻击，确保交易只能在特定的链上被接受
	// 如果你连接的是主网，chainID 应该是 1；如果是 Ropsten 测试网，chainID 应该是 3；如果是 Rinkeby 测试网，chainID 应该是 4；如果是 Goerli 测试网，chainID 应该是 5；如果是 Sepolia 测试网，chainID 应该是 11155111
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("创建 TransactOpts 失败:", err)
	}

	// 2调用 Transfer 方法,这里的交易已经被签名好了，无需手动签名，也不需要设置nonce、gas等参数，abibind 会帮你处理
	// 你只需要传入接收者地址和转账金额（单位是最小单位，比如 wei 或 token 的最小单位）
	tx, err := tokenInstance.Transfer(
		auth,                            // 使用 TransactOpts
		toAddress,                       // 接收者地址
		big.NewInt(1000000000000000000), // 1 Token
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Transfer 交易已构建，TxHash:", tx.Hash().Hex())

	// 3最后你还需要用 client 发送签名交易（如果 Transfer 方法没有自动发送）
	// 但通常 abigen 生成的 Transfer 方法返回的是 *types.Transaction，你需要自己：
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
}
