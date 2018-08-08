package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"MyPublicChain/BLC"
)

func main() {
	block := BLC.NewBlock("create a block for boltdb", make([]byte, 3, 3), 0)

	//测试Block存入到数据库Boltdb中
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	//将block存入到数据库中

	err = db.Update(func(tx *bolt.Tx) error {
		//创建一个bucket
		bucket, err := tx.CreateBucketIfNotExists([]byte("blocks"))
		if err != nil {
			log.Panic(err)
		}
		if bucket != nil {
			err = bucket.Put([]byte("l"), block.Serialize())
			if err != nil {
				fmt.Println("数据存储失败")
			}

		}

		return nil
	})

	if err != nil {
		fmt.Println("更新数据库失败")
	}

	err = db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte("blocks"))
		if bucket != nil {

			dataBytes := bucket.Get([]byte("l"))
			block := BLC.DeSerializeBlock(dataBytes)
			fmt.Println(block)

		}

		return nil

	})
	if err != nil {
		log.Panic(err)
	}
}
