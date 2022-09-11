package BLC

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// 区块基本结构与功能管理文件

// 1.实现一个最基本本的区块结构

type Block struct {
	TimeStamp    int64  // 时间戳
	Hash         []byte // 当前区块HASH
	PreBlockHash []byte // 前区块哈希
	Height       int64  // 区块高度
	Data         []byte //交易数据
}

// 2.新建一个区块

func NewBlock(height int64, prevBlockHash []byte, hash []byte, data []byte) *Block {
	//声明区块对象
	var block Block

	block = Block{
		TimeStamp:    time.Now().Unix(),
		Hash:         hash,
		PreBlockHash: prevBlockHash,
		Height:       height,
		Data:         data,
	}
	// 3.生成哈希
	block.SetHash()

	return &block
}

// function method
// 3.生成哈希-计算区块哈希

func (b *Block) SetHash() {
	// 调用sha256实现哈希生成
	// 实现int64->hash
	timeStampBytes := IntToHex(b.TimeStamp)
	heightBytes := IntToHex(b.Height)
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		b.PreBlockHash,
		b.Data,
	}, []byte{})

	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}
