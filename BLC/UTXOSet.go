package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"encoding/hex"
	"fmt"
	"bytes"
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
			//fmt.Println("啦啦啦啦。。。。。")
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

//增加功能：
/*
添加一个方法，用于查询要转账的utxo
二狗，转账小花：5
	二狗：
		utxo
 */

func (utxoSet *UTXOSet) FindSpentableUTXOs(from string, amount int64, txs []*Transaction) (int64, map[string][]int) {
	var total int64
	//用于存储转账所使用utxo
	spentableUTXOMap := make(map[string][]int)
	//1.查询未打包可以使用的utxo：txs
	unPackageSpentableUTXOs := utxoSet.FindUnpackeSpentableUTXO(from, txs)

	for _, utxo := range unPackageSpentableUTXOs {
		total += utxo.Output.Value
		txIDStr := hex.EncodeToString(utxo.TxID)
		spentableUTXOMap[txIDStr] = append(spentableUTXOMap[txIDStr], utxo.Index)
		if total >= amount {
			return total, spentableUTXOMap
		}
	}



	//2.查询utxotable，查询utxo
	//已经存储的但是未花费的utxo
	err := utxoSet.BlockChian.DB.View(func(tx *bolt.Tx) error {
		//查询utxotable中，未花费的utxo
		b := tx.Bucket([]byte(utxosettable))
		if b != nil {
			//查询
			c := b.Cursor()
		dbLoop:
			for k, v := c.First(); k != nil; k, v = c.Next() {
				txOutputs := DeserializeTxOutputs(v)
				for _, utxo := range txOutputs.UTXOs {
					if utxo.Output.UnlockWithAddress(from) {
						total += utxo.Output.Value
						txIDStr := hex.EncodeToString(utxo.TxID)
						spentableUTXOMap[txIDStr] = append(spentableUTXOMap[txIDStr], utxo.Index)
						if total >= amount {
							break dbLoop
							//return nil
						}
					}
				}

			}

		}

		return nil

	})
	if err != nil {
		log.Panic(err)
	}

	return total, spentableUTXOMap
}


//查询未打包的tx中，可以使用的utxo
func (utxoSet *UTXOSet) FindUnpackeSpentableUTXO(from string, txs []*Transaction) []*UTXO {
	//存储可以使用的未花费utxo
	var unUTXOs []*UTXO

	//存储已经花费的input
	spentedMap := make(map[string][]int)

	for i := len(txs) - 1; i >= 0; i-- {
		//func caculate(tx *Transaction, address string, spentTxOutputMap map[string][]int, unSpentUTXOs []*UTXO) []*UTXO {
		unUTXOs = caculate(txs[i], from, spentedMap, unUTXOs)
	}

	return unUTXOs
}


/*
每次转账后，更新UTXOSet：
转账产生了新的区块：
	交易：
		Input：引用之前的output
		Ouptut：

step1：删除本次交易产生的input对应的utxo
step2：添加本次交易产生的新utxo
 */
func (utxoSet *UTXOSet) Update() {
	/*
	表：key：txID
		value：TxOutputs
			UTXOs []UTXO
	 */

	//1.获取最后(从后超前遍历)一个区块,遍历该区块中的所有tx
	newBlock := utxoSet.BlockChian.Iterator().Next()
	//2.获取所有的input
	inputs := [] *TxInput{}
	//遍历交易，获取所有的input
	for _, tx := range newBlock.Txs {
		if !tx.IsCoinBaseTransaction() {
			for _, in := range tx.Vins {
				inputs = append(inputs, in)
			}
		}
	}

	fmt.Println(len(inputs)) //5
	/*
	以上的内容：找出区块中所有的花费
	Inputs{
	Input{9898, 0, sign,二狗}
	Input{6474,0,sign，小花} //4
	Input{3636,0,sign，小花} //10
	Input{2727,1,sign，小花} //2
}
	 */

	//存储该区块中的，tx中的未花费
	outsMap := make(map[string]*TxOutputs)

	//3.获取所有的output
	for _, tx := range newBlock.Txs {
		utxos := []*UTXO{}
		//找出太交易中的未花费
		for index, output := range tx.Vouts {
			isSpent := false
			//遍历inputs的数组，比较是否有intput和该output对应，如果满足，表示花费了
			for _, input := range inputs {
				if bytes.Compare(tx.TxID, input.TxID) == 0 && index == input.Vout {
					if bytes.Compare(output.PubKeyHash, PubKeyHash(input.PublicKey)) == 0 {
						isSpent = true
					}
				}
			}
			if isSpent == false {
				//output未花
				utxo := &UTXO{tx.TxID, index, output}
				utxos = append(utxos, utxo)
			}
		}

		//utxos,
		if len(utxos) > 0 {
			txIDStr := hex.EncodeToString(tx.TxID)
			outsMap[txIDStr] = &TxOutputs{utxos}
		}

	}
	/*
	以上的遍历，为了找出所有的未花费，存入到map
	map：
	map[6474] = TxOutputs{
	   utxo:{6474,1,output{6，二狗}}
   }

map[9090] = Txoutputs{
	   utxo:{9090,0,Output{15, rose}}
	   utxo:{9090,1,Output{1, 小花}}
   }

map[8989] = Txoutputs{
	   utxo:{8989,0,Output{10,二狗}}
   }


	 */

	//删除花费了数据,添加未花费
	err := utxoSet.BlockChian.DB.Update(func(tx *bolt.Tx) error {
		/*
		数据库：
			txID--->txoutputs
			9898->txoutputs{
				utxo:10,二狗
				utxo：5,小花
				utxo：7,rose
		}
		input{9898,0,二狗}

		utxos：{}
				txoutputs.UTXOs-->utxo

		utxo:{
			utxo：5,小花
			utxo：7,rose
		}


		txoutputs{
			utxo：5,小花
				utxo：7,rose
		}
		b.Delete(9898)

		b.Put(9898,txoutputs)


		 */
		b := tx.Bucket([]byte(utxosettable))
		if b != nil {
			//遍历inputs，删除
			for _, input := range inputs {
				txOutputsBytes := b.Get(input.TxID)
				if len(txOutputsBytes) == 0 {
					continue
				}

				//反序列化
				txOutputs := DeserializeTxOutputs(txOutputsBytes)
				//是否需要被删除
				isNeedDelete := false

				//存储该txoutout中未花费utxo
				utxos := []*UTXO{}

				for _, utxo := range txOutputs.UTXOs {
					if bytes.Compare(utxo.Output.PubKeyHash, PubKeyHash(input.PublicKey)) == 0 && input.Vout == utxo.Index {
						isNeedDelete = true
					} else {
						utxos = append(utxos, utxo)
					}
				}

				if isNeedDelete == true {
					b.Delete(input.TxID)
					if len(utxos) > 0 {
						txOutputs := &TxOutputs{utxos}
						b.Put(input.TxID, txOutputs.Serialize())
					}
				}
			}

			//遍历map，添加
			for txIDStr, txOutputs := range outsMap {
				txID, _ := hex.DecodeString(txIDStr)
				b.Put(txID, txOutputs.Serialize())

			}

		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}
