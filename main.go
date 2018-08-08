package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte("mybucket"))
		//创建游标对数据进行遍历
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			fmt.Printf("key:%s, value:%s\n", k, v)
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}
