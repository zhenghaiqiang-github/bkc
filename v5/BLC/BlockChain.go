package BLC

// BlockChain 2.1区块链管理文件
// 2.1区块链的基本结构
type BlockChain struct {
	Blocks []*Block // 区块的切片

}

// CreateBlockChainWithGenesisBlock 2.2初始化区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	//2.2生成创世区块
	block := CreateGenesisBlock([]byte("init blockchain"))
	return &BlockChain{[]*Block{block}}
}

// AddBlock 2.3添加区块到区块链中
func (bc *BlockChain) AddBlock(height int64, preBlockHash []byte, data []byte) {
	//var newBlock *Block
	newBlock := NewBlock(height, preBlockHash, nil, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}
