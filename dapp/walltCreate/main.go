package main

import (
	"fmt"
	"log"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//创建钱包
	// 1. 生成一个新的私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}
	// 2. 从私钥导出公钥
	// 公钥用于生成以太坊地址，地址是公钥的哈希值的一部分。公钥可以公开分享，用于验证交易签名。
	// 私钥用于签署交易，必须保密。拥有私钥的人可以控制与该私钥对应的以太坊地址上的资产。
	// 地址是用户在以太坊网络上的身份标识，用于接收和发送以太币及其他代币。
	// 什么场景用到公钥，接收和发送以太币时已经有账户地址了为何还要公钥
	// 签名验证：在进行交易时，公钥用于验证交易的签名，确保交易是由拥有私钥的用户发起的。
	// 共享身份：公钥可以公开分享，其他用户可以通过公钥向对应的以太坊地址发送资产。
	// 多重签名：在某些场景下，可能需要多个公钥来共同签署一笔交易，以提高安全性。

	publicKey := privateKey.Public()
	fmt.Printf("Public Key: %s\n", publicKey)
	// 生成以太坊地址，这个地址是公钥的哈希值的一部分，可以用来接收和发送以太币及其他代币
	address := crypto.PubkeyToAddress(*publicKey.(*ecdsa.PublicKey))
	fmt.Printf("Ethereum Address: %s\n", address.Hex())
	// 这里的地址是0x开头的40个十六进制字符，表示以太坊网络上的一个账户地址
	//这个地址跟keystore创建的地址是一样的吗
	fmt.Printf("Private Key: %x\n", crypto.FromECDSA(privateKey))
	// 这里的私钥是16进制字符串，必须保密，不能泄露给他人

	// 3. 创建钱包
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	// 4. 创建一个新的账户（钱包地址）可以用来接收和发送以太币及其他代币
	password := "123456"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatalf("Failed to create new account: %v", err)
	}
	fmt.Printf("New account created: %s\n", account.Address.Hex())
	// 5. 将私钥导入到钱包中
	importedAccount, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		log.Fatalf("Failed to import private key: %v", err)
	}
	fmt.Printf("Imported account: %s\n", importedAccount.Address.Hex())
	// 6. 钱包创建完成
	fmt.Println("Wallet created successfully!")
}
