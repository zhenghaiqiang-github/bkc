package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
)

// BlockChain 2.1区块链管理文件

// 数据库名称
const dbName = "block.db"

// 表名称
const blockTableName = "blocks"

// BlockChain 2.1区块链的基本结构
type BlockChain struct {
	//Blocks []*Block // 区块的切片
	DB  *bolt.DB //数据库对象
	Tip []byte   // 保存最新区块的哈希值
}

// CreateBlockChainWithGenesisBlock 2.2初始化区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	//保存最新区块哈希值
	var blockHash []byte
	//2.2生成创世区块
	//block := CreateGenesisBlock([]byte("init blockchain"))
	// 5.1创建或者打开一个数据库
	// w r x 4 2 1
	db, err := bolt.Open(dbName, 0600, nil)
	if nil != err {
		log.Panicf("create db [%s] failed %v \n", dbName, err)
	}
	// 5.2创建桶,把生成的创世区块存到数据库中
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//没找到桶
			b, err = tx.CreateBucket([]byte(blockTableName))
		}
		if nil != err {
			log.Panicf("create bucket [%s] failed %v \n", blockTableName, err)
		}

		// 生成创世区块
		genesisBlock := CreateGenesisBlock([]byte("init blockchain"))
		//存储
		//1.key,value分别以什么数据代表--hash
		//2.如何把block结构存入到数据库中--序列化
		err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
		if nil != err {
			log.Panicf("insert the genesis block failed %v\n", err)
		}
		blockHash = genesisBlock.Hash
		//存储最新区块的哈希
		//1:latest
		err = b.Put([]byte("1"), genesisBlock.Hash)
		if nil != err {
			log.Panicf("save the hash of genesis block failed %v\n", err)
		}
		return nil

	})
	return &BlockChain{DB: db, Tip: blockHash}
}

// AddBlock 2.3添加区块到区块链中
func (bc *BlockChain) AddBlock(data []byte) {
	// 更新区块数据(insert)
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		//1.获取数据库桶
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			//2.获取最后插入的区块
			blockBytes := b.Get(bc.Tip)
			//3.区块数据反序列化
			latest_block := DeserializeBlock(blockBytes)
			//3.新建区块
			newBlock := NewBlock(latest_block.Height+1, latest_block.Hash, nil, data)
			//4. 存入数据库
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if nil != err {
				log.Panicf("insert the new block to db failed%v", err)
			}
			// 更新最新区块的哈希(数据库)
			err = b.Put([]byte("1"), newBlock.Hash)
			if nil != err {
				log.Panicf("update the latest block hash to db failed%v", err)
			}
			//更行区块链对象中的最新区块哈希
			bc.Tip = newBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Panicf("insert block to db failed %v", err)
	}
}

// 遍历数据库，输出所有区块信息
func (bc *BlockChain) PrintChain() {
	//读取数据库
	fmt.Println("区块链完整信息...")

	var curBlock *Block
	bcit := bc.Iterator() //获取迭代器对象
	//var currentHash []byte = bc.Tip

	// 循环读取
	// 退出条件
	for {
		fmt.Println("------------------------------------------")
		//bc.DB.View(func(tx *bolt.Tx) error {
		//	b := tx.Bucket([]byte(blockTableName))
		//	if b != nil {
		//		blockBytes := b.Get(currentHash)
		//		curBlock = DeserializeBlock(blockBytes)
		//		//输出区块详情
		//		fmt.Printf("\tHash:%x\n", curBlock.Hash)
		//
		//		fmt.Printf("\tPreBlockHash:%x\n", curBlock.PreBlockHash)
		//
		//		fmt.Printf("\tTimeStamp:%v\n", curBlock.TimeStamp)
		//
		//		fmt.Printf("\tData:%v\n", curBlock.Data)
		//
		//		fmt.Printf("\tHeight:%d\n", curBlock.Height)
		//
		//		fmt.Printf("\tNonce:%d\n", curBlock.Nonce)
		//
		//	}
		//	return nil
		//})
		curBlock = bcit.Next()
		fmt.Printf("\tHash:%x\n", curBlock.Hash)

		fmt.Printf("\tPreBlockHash:%x\n", curBlock.PreBlockHash)

		fmt.Printf("\tTimeStamp:%v\n", curBlock.TimeStamp)

		fmt.Printf("\tData:%v\n", curBlock.Data)

		fmt.Printf("\tHeight:%d\n", curBlock.Height)

		fmt.Printf("\tNonce:%d\n", curBlock.Nonce)
		// 退出条件
		// 转换为big.Int
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PreBlockHash)
		// 比较
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			//遍历到创世区块
			break
		}

		//更新当前要获取的区块哈希值
		//currentHash = curBlock.PreBlockHash
	}
}
