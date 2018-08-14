package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//定义区块链的迭代器，专门用于迭代遍历该区块链对应的数据库中block对象
type BlockChainIterator struct {
	DB          *bolt.DB
	CurrentHash []byte
}

func (bcIterator *BlockChainIterator) Next() *Block {
	block := new(Block)
	//1.根据bcIterator，操作DB对象，读取数据库
	err := bcIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucketName))
		if b != nil {
			//根据current获取对应的区块的数据
			blockBytes := b.Get(bcIterator.CurrentHash)
			//反序列化后得到block对象
			block = DeserializeBlock(blockBytes)
			//更新currenthash
			bcIterator.CurrentHash = block.PrevBlockHash
		}
		return nil

	})
	if err != nil {
		log.Panic(err)
	}
	return block
}
