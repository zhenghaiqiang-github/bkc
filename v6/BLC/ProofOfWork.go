package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 共识算法管理文件
// 实现POW实例以及相关功能

// 目标难度值
const targetBit = 16

// ProofOfWork 3.1工作量证明的结构
type ProofOfWork struct {
	// 需要共识验证的区块
	Block *Block
	// 目标难度的哈希
	target *big.Int
}

// NewProofOfWork 创建一个Pow对象
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	// 数据长度为8位
	// 需求：需要满足前两位为0，才能解决问题
	// 1 * 2 << (8-2) = 64
	// 0100 0000
	// 0011 1111 = 63
	// 32 * 8
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{
		Block:  block,
		target: target,
	}
}

// Run 3.3 哈希碰撞 执行pow 比较哈希
// 返回哈希值，以及碰撞次数
func (proofOfWork *ProofOfWork) Run() ([]byte, int) {
	// 碰撞次数
	var nonce = 0
	var hashInt big.Int
	var hash [32]byte // 生成的哈希值
	//无限循环，生成符合条件的哈希值
	for {
		//生成准备数据
		dataBytes := proofOfWork.prepareData(int64(nonce))
		hash = sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		//检测生成的哈希值是否符合条件
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			// 找到了符合条件的哈希值，中断循环
			break
		}
		nonce++
	}
	fmt.Printf("\n碰撞次数:%d\n", nonce)
	return hash[:], nonce
}

// 3.4生成准备数据
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	var data []byte
	// 拼接区块属性，进行哈希计算
	timeStampBytes := IntToHex(pow.Block.TimeStamp)
	heightBytes := IntToHex(pow.Block.Height)
	data = bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		pow.Block.PreBlockHash,
		pow.Block.Data,
		IntToHex(nonce),
		IntToHex(targetBit),
	}, []byte{})
	return data
}
