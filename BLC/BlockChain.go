package BLC

//定义一个区块链(区块的数组)
type BlockChain struct {
	Blocks []*Block
}

//创建一个区块链，包含创世区块
func CreateBlockChainWithGenesisBlock(data string) *BlockChain {
	//1.创建创世区块
	genesisBlock := CreateGenesisBlock(data)
	//2.创建区块链对象并返回
	return &BlockChain{Blocks: []*Block{genesisBlock}}
}

//添加区块到区块链中
func (bc *BlockChain) AddBlockToBlockChain(data string, prevBlockHash []byte, height int64) {

	//1.根据参数的数据，创建Block
	newBlock := NewBlock(data, prevBlockHash, height)

	//2.将block加入blockchain
	bc.Blocks = append(bc.Blocks, newBlock)

}
