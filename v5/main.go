package main

import (
	"bkc/v3/BLC"
	"fmt"
)

// 1.5启动
// 2.4启动
func main() {
	bc := BLC.CreateBlockChainWithGenesisBlock()
	fmt.Printf("blockchain:%v\n", bc.Blocks[0])
	//上链
	bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Height+1,
		bc.Blocks[len(bc.Blocks)-1].Hash,
		[]byte("a send 10 btc to b"))
	bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Height+1,
		bc.Blocks[len(bc.Blocks)-1].Hash,
		[]byte("c send 5 btc to d"))
	for _, block := range bc.Blocks {
		//fmt.Printf("block : %v \n", block)
		fmt.Printf("prevBlockHash : %x, currentHash : %x \n", block.PreBlockHash, block.Hash)
	}
}
