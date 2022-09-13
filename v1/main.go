package main

import (
	"bkc/v1/BLC"
	"fmt"
)

// 1.5启动
func main() {
	block := BLC.NewBlock(1, nil, nil, []byte("the first block testing"))
	fmt.Printf("the first block testing : %v\n", block)
}
