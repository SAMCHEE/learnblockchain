package main

import (
	"fmt"

	"github.com/szlove/learnblockchain/blockchain"
)

func AddBlocks(c *blockchain.Chain) {
	for _, v := range []uint8{1, 2, 3} {
		data := fmt.Sprintf("i am just a block%d..", v)
		c.AddBlock(data)
	}
}

func ListBlocks(c *blockchain.Chain) {
	fmt.Println()
	for i, b := range c.Blocks {
		fmt.Printf("Index    : %d\n", i)
		fmt.Printf("PrevHash : %x\n", b.PrevHash)
		fmt.Printf("Data     : %s\n", b.Data)
		fmt.Printf("Hash     : %x\n", b.Hash)
		fmt.Printf("Nonce    : %d\n", b.Nonce)
		pow := blockchain.NewProof(b)
		fmt.Printf("PoW      : %t\n", pow.Valid())
		fmt.Println()
	}
}

func main() {
	chain := blockchain.InitChain()
	AddBlocks(chain)
	ListBlocks(chain)
}
