package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"encoding/hex"
	"fmt"
)

/*
持久化：
	数据库：blockchain.db
		数据表(bucket) blocks
			存储所有的block

		数据表(bucket) utxoset
			存储所有的未花费utxo


查询余额，转账
 */
type UTXOSet struct {
	BlockChian *BlockChain
}

const utxosettable = "utxoset"

//提供一个重置的功能：获取blockchain中所有的未花费utxo

/*
查询block块中所有的未花费utxo：执行FindUnspentUTXOMap--->map

 */
func (utxoset *UTXOSet) ResetUTXOSet() {
	err := utxoset.BlockChian.DB.Update(func(tx *bolt.Tx) error {
		//1.utxoset表存在，删除
		b := tx.Bucket([]byte(utxosettable))
		if b != nil {
			err := tx.DeleteBucket([]byte(utxosettable))
			if err != nil {
				log.Panic("重置时，删除表失败。。")
			}
		}

		//2.创建utxoset
		b, err := tx.CreateBucket([]byte(utxosettable))
		if err != nil {
			log.Panic("重置时，创建表失败。。")
		}
		if b != nil {
			//3.将map数据--->表
			unUTXOMap := utxoset.BlockChian.FindUnspentUTXOMap()
			/*
			map:
				key:[string]-->[]byte
				value:*Txoutputs{[]UTXO}



			 */
			for txIDStr, outs := range unUTXOMap {
				txID, _ := hex.DecodeString(txIDStr) //[]byte
				b.Put(txID, outs.Serialize())
			}
			fmt.Println("啦啦啦啦。。。。。")
		}

		return nil

	})

	if err != nil {
		log.Panic(err)
	}

}

//查询余额
func (utxoSet *UTXOSet) GetBalance(address string) int64 {
	utxos := utxoSet.FindUnspentUTXOsByAddress(address)

	var total int64

	for _, utxo := range utxos {
		total += utxo.Output.Value
	}
	return total
}

//根据指定的地址，找出所有的utxo
func (utxoSet *UTXOSet) FindUnspentUTXOsByAddress(address string) []*UTXO {
	var utxos []*UTXO
	err := utxoSet.BlockChian.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxosettable))
		if b != nil {
			/*
			获取表中的所有的数据
			key,value
			key:TxID
			value：TxOuputs
			 */
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				txOutputs := DeserializeTxOutputs(v)
				for _, utxo := range txOutputs.UTXOs { //txid, index,output
					if utxo.Output.UnlockWithAddress(address) {
						utxos = append(utxos, utxo)
					}
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return utxos
}
