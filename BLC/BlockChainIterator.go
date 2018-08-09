package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//定义区块链的迭代器，用于迭代遍历该区块链对应的数据库中的Block对象

type BlockChainIterator struct {
	DB          *bolt.DB
	CurrentHash []byte
}

func (bcIterator *BlockChainIterator) Next() *Block {
	var block *Block
	err := bcIterator.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		if bucket != nil {
			//根据current获取对应的区块数据
			blockBytes := bucket.Get(bcIterator.CurrentHash)
			//反序列化后得到Block对象
			block = DeSerializeBlock(blockBytes)
			bcIterator.CurrentHash = block.PrevBlockHash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return block

}
