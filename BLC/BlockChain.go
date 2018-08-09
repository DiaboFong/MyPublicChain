package BLC

import (
	"github.com/boltdb/bolt"
	"os"
	"fmt"
	"log"
	"time"
	"math/big"
)

//定义一个区块链(区块的数组)
type BlockChain struct {
	//Blocks []*Block
	DB  *bolt.DB //对应的数据库对象
	Tip []byte   // 存储区块中最后一个块的Hash值
}

//创建一个区块链，包含创世区块====>拆解成两块内容(提供给命令行功能使用)
/*
1.创建创世区块CreateGenesisBlockToDB(data string)
(1) 如果数据库存在，返回:创世区块已存在，退出程序

(2) 如果数据库不存在，则创建创世区块



2.获取BlockChain对象GetBlockChainObject()
(1) 如果数据库存在，则从数据库中取数据并返回BlockChain对象
(2) 如果数据库不存在，则返回创世区块不存在，请通过xxx命令行创建创世区块


 */
func CreateGenesisBlockToDB(data string) {

	if dbExists() {
		fmt.Println("创世区块已存在，你可以继续添加新的区块")
		printUsage()
		os.Exit(1)
	}

	fmt.Println("创世区块不存在，开始创建")
	/*如果数据库不存在
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

}

//定义一个函数，用于获取BlockChain对象

func GetBlockChainObject() *BlockChain {

	//如果数据库存在，直接获取数据库中l对应的最新hash
	if dbExists() {
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

	return nil
}

//添加区块到区块链中(存储至Boltdb)
func (bc *BlockChain) AddBlockToBlockChain(data string) {

	err := bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		if bucket != nil {
			//获取bc的Tip(最新的Hash),从数据库中读取最后一个Block,获取其Hash,Height
			blockBytes := bucket.Get(bc.Tip)
			lastBlock := DeSerializeBlock(blockBytes)
			//创建新的区块
			newBlock := NewBlock(data, lastBlock.Hash, lastBlock.Height+1)
			err := bucket.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//更新bc的TIP，以及数据库中l的值
			bc.Tip = newBlock.Hash
			bucket.Put([]byte("l"), newBlock.Hash)

		}
		return nil

	})
	if err != nil {
		log.Panic(err)
	}
}

//定义一个方法，用于判断数据库是否存在
func dbExists() bool {
	//获取文件对应的信息
	if _, err := os.Stat(DBName); os.IsNotExist(err) {
		return false
	}
	return true

}

//定义一个用于遍历数据库的方法，打印所有区块
func (bc *BlockChain) PrintChains() {
	/*
	1. bc.DB.View()
	根据Hash,获取Block数据
	反序列化
	打印输出
	 */
	//(1) 获取迭代器
	iterator := bc.Iterator()
	block := new(Block)
	//(2) 根据迭代器的Next()方法获取Block对象
	for {
		block = iterator.Next()
		fmt.Printf("第%d个区块信息如下:\n", block.Height+1)
		fmt.Printf("区块高度:%d\n", block.Height)
		fmt.Printf("上一个区块哈希:%x\n", block.PrevBlockHash)
		fmt.Printf("区块哈希:%x\n", block.Hash)
		fmt.Printf("区块交易:%s\n", block.Data)
		fmt.Printf("区块时间戳:%s\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("区块随机数%d\n", block.Nonce)
		fmt.Println()
		//2.判断区块的prevBlockHash是否为0，

		// 为0  : 表示该Block是创世区块,结束循环
		hashBigInt := new(big.Int)
		hashBigInt.SetBytes(block.PrevBlockHash)
		if hashBigInt.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("这是创世区块，数据查询结束")
			break
		}

	}

}

//获取blockchainitetor的对象
func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{DB: bc.DB, CurrentHash: bc.Tip}
}
