package BLC

import (
	"github.com/boltdb/bolt"
	"os"
	"fmt"
	"log"
)

//定义一个区块链(区块的数组)
type BlockChain struct {
	//Blocks []*Block
	DB  *bolt.DB //对应的数据库对象
	Tip []byte   // 存储区块中最后一个块的Hash值
}

//创建一个区块链，包含创世区块
func CreateBlockChainWithGenesisBlock(data string) *BlockChain {

	//数据库如果存在，则表示创世区块已经创建 ，取值即可
	if dbExists() {
		fmt.Println("数据库已经存在...")
		//打开数据库
		db, err := bolt.Open(DBName, 0600, nil)
		if err != nil {
			log.Panic(err)
		}
		var blockchain *BlockChain

		err = db.View(func(tx *bolt.Tx) error {
			//打开Bucket，读取l对应的最新Hash值
			bucket := tx.Bucket([]byte(BucketName))
			if bucket != nil {
				//读取最新的hash
				hash := bucket.Get([]byte("l"))
				blockchain = &BlockChain{DB: db, Tip: hash}

			}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		return blockchain

	}

	//数据库不存在，创建创世区块，并存入数据库中
	fmt.Println("数据库不存在")
	/*
	1.创建创世区块
	2.存入到数据库中
	 */
	//1.创建创世区块
	genesisBlock := CreateGenesisBlock(data)
	db, err := bolt.Open(DBName, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		//创世区块序列化后存入到数据库中
		bucket, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			log.Panic(err)
		}

		if bucket != nil {
			err := bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			bucket.Put([]byte("l"), genesisBlock.Hash)
		}
		return nil

	})
	if err != nil {
		log.Panic(err)
	}
	return &BlockChain{DB: db, Tip: genesisBlock.Hash}

}

//添加区块到区块链中
func (bc *BlockChain) AddBlockToBlockChain(data string, prevBlockHash []byte, height int64) {

/*	//1.根据参数的数据，创建Block
	newBlock := NewBlock(data, prevBlockHash, height)

	//2.将block加入blockchain
	bc.Blocks = append(bc.Blocks, newBlock)
*/
}

//定义一个方法，用于判断数据库是否存在
func dbExists() bool {
	//获取文件对应的信息
	if _, err := os.Stat(DBName); os.IsNotExist(err) {
		return false
	}
	return true

}
