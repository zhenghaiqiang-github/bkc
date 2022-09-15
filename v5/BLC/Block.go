package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

// Block 区块基本结构与功能管理文件
//
// 1.1实现一个最基本本的区块结构
type Block struct {
	TimeStamp    int64  // 时间戳
	Hash         []byte // 当前区块HASH
	PreBlockHash []byte // 前区块哈希
	Height       int64  // 区块高度
	Data         []byte //交易数据
	Nonce        int64  //在运行pow时生成的哈希变化值，也代表pow运行时动态修改的数据
}

// NewBlock 1.2新建一个区块
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
	// 1.3生成哈希
	block.SetHash()
	// 3.1 同过POW生成哈希值
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = int64(nonce)
	return &block
}

// SetHash function method
// 1.3生成哈希-计算区块哈希
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

// CreateGenesisBlock 生成创世区块
func CreateGenesisBlock(data []byte) *Block {
	return NewBlock(1, nil, nil, data)
}

// 区块结构化序列
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	//新建编码对象
	encoder := gob.NewEncoder(&buffer)
	//编码(序列化)
	if err := encoder.Encode(block); err != nil {
		log.Panicf("serialize the block to []byte failed %v\n", err)
	}

	return buffer.Bytes()
}

// 区块数据反序列化
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&block); err != nil {
		log.Panicf("deserialize the []byte to block failed! %v\n", err)
	}

	return &block
}
