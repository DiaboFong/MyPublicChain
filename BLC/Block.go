package BLC

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//1.定义一个Block

type Block struct {
	//定义字段属性
	//高度：区块在区块链中的编号，第一个区块叫做创世区块，高度为0
	Height int64
	//上一个区块的Hash值
	PrevBlockHash []byte
	//交易数据
	Data []byte
	//时间戳
	TimeStamp int64
	//区块自己的Hash值
	Hash []byte
	//随机数 Nonce
	Nonce int64
}

//2. 定义一个函数用于创建一个区块
func NewBlock(data string, prevBlock []byte, height int64) *Block {
	//创建区块
	block := &Block{Height: height, PrevBlockHash: prevBlock, Data: []byte(data), TimeStamp: time.Now().Unix()}
	//设置区块Hash ===> 通过POW方法计算出Hash值
	/*	block.SetHash()
		return block*/

	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

//3. 设置区块的Hash值
func (block *Block) SetHash() {
	//可以通过当前的block属性值来生成Hash，保存为[]byte
	//1. 转Height
	heightBytes := IntToHex(block.Height)
	//2. 转TimeStamp(另外一种方式)
	timeStampString := strconv.FormatInt(block.TimeStamp, 2)
	timeStampBytes := []byte(timeStampString)
	//3.通过join拼接所有的[]byte
	//join(s [][]byte, sep []byte) []byte
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeStampBytes,
	}, []byte{})
	//设置到Block上
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]

}

//4.生成创世区块
func CreateGenesisBlock(data string) *Block {

	return NewBlock(data, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0)

}

//序列化-编码,将block转换成buff
func (block *Block) Serialize() []byte {
	//1.创建一个Buff对象
	var buf bytes.Buffer
	//2. 创建一个编码器
	encoder := gob.NewEncoder(&buf)
	//3.对block进行编码
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}

//发序列化-解码,将blockBytes转换成block
func DeSerializeBlock(blockBytes []byte) *Block {
	var block Block
	//1.创建一个reader对象
	reader := bytes.NewReader(blockBytes)
	//2.创建解码器
	decoder := gob.NewDecoder(reader)
	//3.解码
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block

}
