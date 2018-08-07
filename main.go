package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func main() {
	/*	//1.创建一个带有创世区块的区块链
		bc := BLC.CreateBlockChainWithGenesisBlock("Create Block Chain With GenesisBlock")
		//2.通过POW挖出第一个区块
		bc.AddBlockToBlockChain("Send $B100 To brucefeng", bc.Blocks[len(bc.Blocks)-1].Hash, bc.Blocks[len(bc.Blocks)-1].Height+1)
		//3.验证区块是否合法
		pow := BLC.NewProofOfWork(bc.Blocks[1])
		isValid := pow.IsValid()
		if isValid {
			fmt.Println("区块合法")
			return
		}
		fmt.Println("区块不合法")*/

	/*
	安装配置BoltDB
	1.安装 		go get "github.com/boltdb/bolt"
	2.打开数据库
	 */
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	/*
db.View() 查看数据库
db.Update() 读写
 */
	//3.存储一个数据[]byte
	db.Update(func(tx *bolt.Tx) error {
		//存储数据
		//1. 创建表
		bucket, err := tx.CreateBucket([]byte("mybucket"))
		if err != nil {
			log.Panic(err) //数据表创建有误

		}
		//2.存储数据
		if bucket !=nil {
			err := bucket.Put([]byte("k"),[]byte("send 100 to brucefeng"))
			if err !=nil {
				fmt.Println("数据存储有误")
			}
		}
		return nil

	})

}
