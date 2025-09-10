package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//eth转账必要条件--发送方地址，接收方地址，ETH余额，发送方私钥
	// 没有私钥 = 无法对交易签名 = 无法发送 ETH！
	// 1. 连接到以太坊节点（比如 Infura / Alchemy / 本地 Geth）
	// rpcURL := "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID" // 或 http://127.0.0.1:8545
	// client, err := ethclient.Dial(rpcURL)
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/mWrjcZw9dhD**h-FvHb2YCI")
	if err != nil {
		log.Fatal("连接节点失败:", err)
	}
	defer client.Close()

	// 2. 设置发送方地址和接收方地址
	fromAddressStr := "0x93dF7b9eaF035DD76c80F814c3db2C52d76Cb107"
	toAddressStr := "0xe3269B4B78240f4E66E79a41bA0EfC2002c77CA7"

	fromAddress := common.HexToAddress(fromAddressStr)
	toAddress := common.HexToAddress(toAddressStr)

	// 3. 设置私钥（十六进制，带 0x 前缀）—— 请务必保密！！！
	// 通常来自 MetaMask 导出、助记词生成等
	privateKeyHex := "7f202274633793f95f72b0d44cc07665482251936ba459**fc0612143d77678ce0"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("无效的私钥:", err)
	}

	// 4. 获取发送方地址的 Nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("获取 Nonce 失败:", err)
	}
	fmt.Printf("Nonce: %d\n", nonce)

	// 5. 设置转账金额（比如 0.01 ETH）
	amountEth := big.NewFloat(0.01) // 转账 0.01 ETH
	amountWei := new(big.Int)
	amountWei.SetString(amountEth.Mul(amountEth, big.NewFloat(1e18)).String(), 10) // 1 ETH = 1e18 Wei
	fmt.Printf("转账金额: %v Wei\n", amountWei)

	// 6. 获取当前建议的 Gas Price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获取 Gas Price 失败:", err)
	}
	fmt.Printf("Gas Price: %v Wei\n", gasPrice)

	// 7. 构造交易
	tx := types.NewTransaction(
		nonce,
		toAddress,
		amountWei,
		21000, // 标准 Gas Limit（普通转账）
		gasPrice,
		nil, // 不含额外数据
	)

	// 8. 设置 Chain ID（主网为 1，Sepolia 为 11155111，根据你的目标网络设置）
	chainID := big.NewInt(11155111) // Sepolia

	// 9. 用私钥对交易进行签名
	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		log.Fatal("交易签名失败:", err)
	}

	// 10. 发送已签名的交易到以太坊网络或测试网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("发送交易失败:", err)
	}

	// 11. 打印交易哈希，可以在 Etherscan 查询
	fmt.Printf("✅ 交易已成功发送！Tx Hash: %s\n", signedTx.Hash().Hex())
}
