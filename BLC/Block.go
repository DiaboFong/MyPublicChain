package BLC

import (
	"time"

	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	//字段属性
	//1.高度：区块在区块链中的编号，第一个区块页叫创世区块，为0
	Height int64
	//2.上一个区块的Hash值
	PrevBlockHash []byte
	//3.数据：data，交易数据
	Txs []*Transaction
	//4.时间戳
	TimeStamp int64
	//5.自己的hash
	Hash []byte

	//6.Nonce
	Nonce int64
}

func NewBlock(txs []*Transaction, prevBlockHash [] byte, height int64) *Block {
	//创建区块
	block := &Block{height, prevBlockHash, txs, time.Now().Unix(), nil,0}
	//设置hash
	//block.SetHash()
	pow:=NewProofOfWork(block)
	hash,nonce:=pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}



func CreateGenesisBlock(txs []*Transaction) *Block{

	return NewBlock(txs,make([]byte,32,32),0)
}

//定义block的方法，用于序列化该block对象，获取[]byte
func (block *Block) Serialize()[]byte{
	//1.创建一个buff
	var buf bytes.Buffer

	//2.创建一个编码器
	encoder:=gob.NewEncoder(&buf)

	//3.编码
	err:=encoder.Encode(block)
	if err != nil{
		log.Panic(err)
	}

	return buf.Bytes()
}

//定义一个函数，用于将[]byte，转为block对象，反序列化
func DeserializeBlock(blockBytes [] byte) *Block{
	var block Block
	//1.先创建一个reader
	reader:=bytes.NewReader(blockBytes)
	//2.创建解码器
	decoder:=gob.NewDecoder(reader)
	//3.解码
	err:=decoder.Decode(&block)
	if err != nil{
		log.Panic(err)
	}
	return &block
}


//提供一个方法，用于将block块中的txs转为[]byte数组

func (block *Block) HashTransactions()[]byte{
	/*
	//1.创建一个二维数组，存储每笔交易的txid
	var txshashes [][] byte
	//2.遍历
	for _,tx:=range block.Txs{

		tx1,tx2,tx3...
		[][]{tx1.ID,tx2.ID,tx3.ID...}

		合并-->[]--->sha256

		 txshashes  = append(txshashes,tx.TxID)
	}
	//3.生成hash
	txhash:=sha256.Sum256(bytes.Join(txshashes,[]byte{}))
	return txhash[:]
	*/
	var txs [][]byte
	for _,tx:=range block.Txs{
		txs = append(txs,tx.Serialize())
	}
	mTree:=NewMerkleTree(txs)
	return mTree.RootNode.DataHash

}