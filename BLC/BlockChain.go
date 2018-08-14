package BLC

import (
	"github.com/boltdb/bolt"
	"os"
	"fmt"
	"log"
	"time"
	"math/big"
	"strconv"
	"encoding/hex"
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
func CreateGenesisBlockToDB(address string) {

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
	//1.创建一个txs -->Coinbase
	txCoinBase := NewCoinBaseTransaction(address)

	genesisBlock := CreateGenesisBlock([]*Transaction{txCoinBase})
	db, err := bolt.Open(DBName, 0600, nil)
	defer db.Close()
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
func (bc *BlockChain) AddBlockToBlockChain(txs []*Transaction) {

	err := bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		if bucket != nil {
			//获取bc的Tip(最新的Hash),从数据库中读取最后一个Block,获取其Hash,Height
			blockBytes := bucket.Get(bc.Tip)
			lastBlock := DeSerializeBlock(blockBytes)
			//创建新的区块
			newBlock := NewBlock(txs, lastBlock.Hash, lastBlock.Height+1)
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
		//fmt.Printf("区块交易:%s\n", block.Data)
		//Data =>txs ,遍历数组获取所有的交易信息
		fmt.Println("交易信息")
		for _, tx := range block.Txs {
			fmt.Printf("\t\t交易ID:%x\n", tx.TxID)
			fmt.Println("\t\tVins:")
			for _, txInput := range tx.Vins { //每个TxInput:TxID,vout,解锁脚本
				fmt.Printf("\t\t\tTxID:%x\n", txInput.TxID)
				fmt.Printf("\t\t\tVout:%d\n", txInput.Vout)
				fmt.Printf("\t\t\tScriptSiq:%s\n", txInput.ScriptSiq)
			}
			fmt.Println("\t\tVouts:")
			for _, txOutput := range tx.Vouts { //每个TxOutput:value,锁定脚本
				fmt.Printf("\t\t\tValue:%d\n", txOutput.Value)
				fmt.Printf("\t\t\tScriptPubkey:%s\n", txOutput.ScriptPubKey)

			}
		}

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

//新增功能：通过转账，创建区块
func (bc *BlockChain) MineNewBlock(from, to, amount []string) {
	/*
	1. 新建交易
	2. 新建区块
		读取数据库，获取最后一块block
	3. 存入到数据库中
	 */
	//1. 新建交易
	var txs []*Transaction

	//amount[0] -->int
	amountInt, _ := strconv.ParseInt(amount[0], 10, 64)
	tx := NewSimpleTransaction(from[0], to[0], amountInt, bc)
	txs = append(txs, tx)

	//2. 新建区块
	newBlock := new(Block)
	err := bc.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		if bucket != nil {
			//读取数据库
			blockBytes := bucket.Get(bc.Tip)
			lastBlock := DeSerializeBlock(blockBytes)
			newBlock = NewBlock(txs, lastBlock.Hash, lastBlock.Height+1)

		}

		return nil

	})
	if err != nil {
		log.Panic(err)
	}

	//3. 存入到数据库中
	err = bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		if bucket != nil {
			//将新block存入到数据中
			bucket.Put(newBlock.Hash, newBlock.Serialize())
			//更新l对应的值
			bucket.Put([]byte("l"), newBlock.Hash)
			//更新Tip
			bc.Tip = newBlock.Hash

		}
		return nil

	})
	if err != nil {
		log.Panic(err)
	}

}

//4. 查询余额,即查询某个账户的未花费的Output
func (bc *BlockChain) GetBalance(address string) int64 {

	//txOutputs := bc.UnSpent(address)
	unSpentUTXOs := bc.UnSpent(address)
	var total int64
	for _, utxo := range unSpentUTXOs {
		total += utxo.Output.Value
	}
	return total

}

//定义方法，用于获取指定用户的所有未花费(UnSpent)的Txoutput
/*
UTXO模型
Unspent Transaction TxOutput
 */
func (bc *BlockChain) UnSpent(address string) []*UTXO {

	/*
	1. 遍历数据库，获取每个Block的Txs
	2. 遍历所有交易：
	Inputs ,将数据记录为已经花费
	Outputs ,将每个output和input进行对比，如果有关联，则表示已经花费，如果没有任何关联，表示为未花费

	 */

	//存储未花费的TxOutput
	var unSpentUTXOs []*UTXO
	//存储已经花费的TxOutput。map[TxID][]int{vout...}
	spentTxOutPutMap := make(map[string][]int) //map[TxID]=[]int{vout}

	iterator := bc.Iterator()
	for {
		//1.获取每个Block
		block := iterator.Next()

		//2.遍历该Block的txs
		for _, tx := range block.Txs {
			//遍历每个Tx:TxID,Vins,Vouts
			//遍历TxInput ,已经花费的
			if !tx.IsCoinBaseTransaction() { //tx不是CoinBase交易，则遍历其Vints
				for _, txInput := range tx.Vins {
					//in ： TxID,Vout,SciptSiq
					if txInput.UnlockWithAddress(address) { //如果Input中的解锁脚本是传入的address，则将
						//txInput的解锁脚本(用户名)如果和要查询余额的用户名一致相同
						key := hex.EncodeToString(txInput.TxID)
						spentTxOutPutMap[key] = append(spentTxOutPutMap[key], txInput.Vout)
						/*
						map[key] --> value
						map[key] --> []int
						 */

					}

				}
			}
			//遍历TxOutput,未花费的
		outputs:
			for index, txOutput := range tx.Vouts {
				if txOutput.UnlockWithAddress(address) {
					if len(spentTxOutPutMap) != 0 {
						//遍历map
						var isSpentOutput bool //false
						for txID, indexArray := range spentTxOutPutMap {
							//遍历记录已经花费的下标的数组
							for _, i := range indexArray {
								if i == index && hex.EncodeToString(tx.TxID) == txID {
									isSpentOutput = true //标记当前的txOutput是否花费
									continue outputs
								}
							}

						}
						if !isSpentOutput {
							//unSpentTxOutput = append(unSpentTxOutput, txOutput)
							//根据未花费的Output创建UTXO对象 -- >数组
							utxo := &UTXO{
								TxID:   tx.TxID,
								Index:  index,
								Output: txOutput,
							}
							unSpentUTXOs = append(unSpentUTXOs, utxo)

						}

					} else { //如果map长度为0,那么证明还没有花费记录，那么表示未花费，output无需花费
						utxo := &UTXO{
							TxID:   tx.TxID,
							Index:  index,
							Output: txOutput,
						}
						unSpentUTXOs = append(unSpentUTXOs, utxo)

					}
				}
			}
		}

		//3.判断退出
		hashInt := new(big.Int)
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(hashInt) == 0 {
			break
		}
	}

	return unSpentUTXOs
}

//提供一个方法，用于一次转账的交易中，可以使用的未花费的UTXO

func (bc *BlockChain) FindSpentableUTXOs(from string, amount int64) (int64, map[string][]int) {
	/*
	1. 根据from获取到所有的UTXO
	2. 遍历UTXOs,累加金额，判断是否大于等于需要转账的金额

	返回:map[TxID] -->[]int{下标...} -->output
	 */


	var total int64
	var spentableMap = make(map[string][]int)
	//获取所有的UTXO:10
	utxos := bc.UnSpent(from)
	//2.找即将使用的UTXO
	for _, utxo := range utxos {
		total += utxo.Output.Value
		txIDstr := hex.EncodeToString(utxo.TxID)
		spentableMap[txIDstr] = append(spentableMap[txIDstr], utxo.Index)
		if total > amount {
			break
		}

	}
	//3.金额不足，无法转账
	if total < amount {
		fmt.Printf("%s,余额不足，无法转账\n", from)
		os.Exit(1)
	}
	return total, spentableMap

}
