package main

import (
	"context"
	"fmt"
	"log"

	store "task5.1/go-bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	privateKey := "a3189fa9381849f460e4ed8b04c0d5c055df81d9b16b6eabe0a597bbf93ef6f9"

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/cef15f4b47934f068d6d15ff6ce629e8")
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}

	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("私钥解析失败: %v", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("获取ChainID失败: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		log.Fatalf("构造交易器失败: %v", err)
	}

	// 4. 加载合约（替换为实际地址）
	contractAddr := common.HexToAddress("0x7408f39e8d59489f59b5435c1d897318bd6158fa")
	instance, err := store.NewStore(contractAddr, client)
	if err != nil {
		log.Fatalf("合约加载失败: %v", err)
	}

	fmt.Println("=== 调用开始 ===")

	value, err := instance.GetCounter(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}
	fmt.Printf("当前计数器值: %d\n", value)

	tx, err := instance.Increment(auth)
	if err != nil {
		log.Fatalf("交易失败: %v", err)
	}
	fmt.Printf("交易已发送，哈希: %s\n", tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("等待确认失败: %v", err)
	}
	fmt.Printf("交易已确认，区块: %d\n", receipt.BlockNumber)

	newValue, err := instance.GetCounter(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}
	fmt.Printf("新计数器值: %d\n", newValue)
}
