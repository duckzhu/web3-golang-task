package main

import (
	"fmt"
	"strings"
	"time"

	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ERC20 ABI你自己用 ​Solidity 写了一个合约​
// 你使用 ​Remix、Hardhat、Foundry、Truffle​ 等工具部署了它
// 如何获取 ABI：
// ​如果你用 Remix：​​
// 编译合约后，在 ​​“Compile” 标签页​ 下，选择你的合约
// 点击右侧的 ​​“ABI” 按钮，就可以复制 ABI（JSON 格式）
// ​如果你用 Hardhat / Truffle：​​
// 编译后，在 artifacts/contracts/YourContract.sol/YourContract.json 文件中
// 里面有一个 "abi": [...] 字段，那就是该合约的 ABI
// 你可以直接读取该 JSON 文件并加载到你的程序中
// ​如果你用 Foundry：​​
// 可以通过 forge inspect 命令导出 ABI
const erc20ABI = `[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "approve",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "allowance",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "needed",
				"type": "uint256"
			}
		],
		"name": "ERC20InsufficientAllowance",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "balance",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "needed",
				"type": "uint256"
			}
		],
		"name": "ERC20InsufficientBalance",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "approver",
				"type": "address"
			}
		],
		"name": "ERC20InvalidApprover",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			}
		],
		"name": "ERC20InvalidReceiver",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			}
		],
		"name": "ERC20InvalidSender",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			}
		],
		"name": "ERC20InvalidSpender",
		"type": "error"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Approval",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "transfer",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "transferFrom",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			}
		],
		"name": "allowance",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "account",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "decimals",
		"outputs": [
			{
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "name",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "symbol",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "totalSupply",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`

const privateKeyHex = "78b7dbc13704a5ed37cd0f2468b944d399df1d9724ad2c9c6e6**472b0d8df6481" // 0xe3269B4B78240f4E66E79a41bA0EfC2002c77CA7

func main() {

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/mWrjcZw9dh**Dh-FvHb2YCI")
	if err != nil {
		log.Fatal("连接节点失败:", err)
	}
	defer client.Close()
	//tx  0x2cbbfe229b044f69e03592c98ef408f525db1f62a29d7075c006208105e264da
	fmt.Println("已部署的token合约地址: 0xdF7e2082B32Ed7955A91a973D7cEBa5AcE647976,mint代币给  1000000")
	// token合约地址
	tokenContractAddressStr := "0xdF7e2082B32Ed7955A91a973D7cEBa5AcE647976"
	tokenContractAddress := common.HexToAddress(tokenContractAddressStr)
	// 接收地址，我自己的钱包测试账户3
	toAddressStr := "0x93dF7b9eaF035DD76c80F814c3db2C52d76Cb107"
	toAddress := common.HexToAddress(toAddressStr)
	// 3. 设置转账的代币数量，比如 100 个代币（注意是小数位！）
	// 假设代币有 6 位小数（比如 USDT），你要转 100 个 => 100 * 10^6
	// decimals := big.NewInt(18)
	// decimals := big.NewInt(4)
	amount := big.NewInt(100000) // 转 100 个代币
	// amount = amount.Mul(amount, big.NewInt(0).Exp(big.NewInt(10), decimals, nil)) // amount = 100 * 10^decimals
	fmt.Printf("准备转账的代币数量（最小单位）: %v\n", amount)
	// 4. 发送者地址
	fromAddressStr := "0xe3269B4B78240f4E66E79a41bA0EfC2002c77CA7"
	fromAddress := common.HexToAddress(fromAddressStr)

	// 5. 加载 ERC20 ABI
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		log.Fatal("解析 ABI 失败:", err)
	}

	//查询发送方代币余额
	balance, err := callBalanceOf(client, parsedABI, tokenContractAddress, fromAddress)
	if err != nil {
		log.Fatal("调用 balanceOf 失败:", err)
	}
	fmt.Printf("地址 %s 在合约 %s 中的交易前的代币余额为: %v\n", fromAddressStr, tokenContractAddressStr, balance)

	//调用合约的 transfer 方法进行代币转账
	signedTx := callTransfer(amount, client, parsedABI, tokenContractAddress, fromAddress, toAddress)

	//等待交易被打包
	time.Sleep(10 * time.Second)
	// 查询交易状态
	_, isPending, err := client.TransactionByHash(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal("查询交易状态失败:", err)
	}
	fmt.Printf("交易是否还在待处理队列中: %v\n", isPending)
	//再次查询发送方的余额
	balance, err = callBalanceOf(client, parsedABI, tokenContractAddress, fromAddress)
	if err != nil {
		log.Fatal("调用 balanceOf 失败:", err)
	}
	fmt.Printf("地址 %s 在合约 %s 中的交易后的代币余额为: %v\n", fromAddressStr, tokenContractAddressStr, balance)
}

// 使用 ethclient 直接调用合约的 balanceOf 函数
func callBalanceOf(client *ethclient.Client, parsedABI abi.ABI, contractAddr, ownerAddr common.Address) (*big.Int, error) {

	// 构造函数调用数据
	data, err := parsedABI.Pack("balanceOf", ownerAddr)
	if err != nil {
		return nil, fmt.Errorf("构造 balanceOf 调用数据失败: %v", err)
	}

	// 构造调用请求
	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}

	// 调用合约（不发送交易，只查询）
	result, err := client.CallContract(context.Background(), msg, nil) // nil 表示最新区块
	if err != nil {
		return nil, fmt.Errorf("调用合约失败: %v", err)
	}

	// 解析返回值：balanceOf 返回的是 uint256（即 []byte 表示的大整数）
	var balance big.Int
	if len(result) < 32 {
		return nil, fmt.Errorf("返回数据长度不足")
	}
	balance.SetBytes(result[12:32]) // 第一个 12 字节是 padding，实际数据从第 12 字节开始（对于 uint256）

	return &balance, nil
}

// 在以太坊上，​代币（如 USDT、USDC、自定义 ERC20 Token）并不是直接通过 send ETH 的方式转账的，而是通过调用该代币合约的 transfer 方法来实现的。​​
// 所以，​代币转账 ≠ ETH 转账，它需要：
// ​知道代币合约地址​
// ​调用该合约的 transfer(address to, uint256 amount) 方法​
// ​构造一笔调用合约方法的交易（不是普通转账）​​
// ​对这笔交易进行签名并发送

func callTransfer(amount *big.Int, client *ethclient.Client, parsedABI abi.ABI, tokenContractAddress common.Address, fromAddress common.Address, toAddress common.Address) *types.Transaction {
	// 1. 构造 transfer 函数的调用数据
	data, err := parsedABI.Pack("transfer", toAddress, amount)
	if err != nil {
		log.Fatal("构造调用数据失败:", err)
	}
	fmt.Println("调用数据（hex）:", fmt.Sprintf("0x%x", data))

	// 2. 构造交易参数
	value := big.NewInt(0) // 代币转账不需要 ETH，所以 value 为 0
	//  gas limit
	gasLimit := uint64(100000) // 通常代币转账 gas limit 在 60k ~ 100k
	// 获取发送方地址的 Nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("获取 Nonce 失败:", err)
	}
	fmt.Printf("Nonce: %d\n", nonce)
	// 获取当前建议的 Gas Price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获取 Gas Price 失败:", err)
	}
	fmt.Printf("Gas Price: %v Wei\n", gasPrice)
	// chain id Sepolia 测试网的 chain ID
	chainID := big.NewInt(11155111)

	if err != nil {
		log.Fatal("无效的私钥:", err)
	}

	tx := types.NewTransaction(
		nonce,
		tokenContractAddress, // 合约地址，不是接收者地址！
		value,
		gasLimit,
		gasPrice,
		data, //接收者的地址位于data中
	)

	// 3. 用私钥对交易进行签名
	// 发送者账户0xe3269B4B78240f4E66E79a41bA0EfC2002c77CA7的私钥sepolia
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("无效的私钥:", err)
	}
	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		log.Fatal("交易签名失败:", err)
	}

	// 4. 发送已签名的交易到以太坊网络或测试网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("发送交易失败:", err)
	}
	// 5. 打印交易哈希，可以在 Etherscan 查询
	fmt.Printf("✅ 交易已成功发送！Tx Hash: %s\n", signedTx.Hash().Hex())
	return signedTx

}
