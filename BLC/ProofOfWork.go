package BLC

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

//工作量证明
//1. 构建POW模型

/*
Hash:256bit == > 16进制:32个字符
16个0       == >  4个0
0000 0000 0000 0000 0000 1000 0000 。。。。
转成16进制
0000 1000
 */

const TargetBit = 16 //目标Hash的0的个数(256位16个0，16进制:4个0)

type ProofOfWork struct {
	Block  *Block   // 要验证的Block
	Target *big.Int //目标hash
}

//2. 提供一个POW对象
func NewProofOfWork(block *Block) *ProofOfWork {

	//1.创建POW对象
	pow := &ProofOfWork{}
	//2.设置属性值
	pow.Block = block

	target := big.NewInt(1)           //目标hash,初始值为1, 0000 0001  （目标2个0，需要左移(8-2)6位）  0010 0000
	target.Lsh(target, 256-TargetBit) //左移256-16
	pow.Target = target
	return pow

}

//3.定义一个函数，用于计算出有效的Hash
//挖矿 ===>不断更改Nonce的值，计算Hash,直到小于目标Hash
func (pow *ProofOfWork) Run() ([]byte, int64) {
	/*
	1. 设置一个Nonce值，默认值为0
	2. 拼接block属性数组，产生hash
	3. 比较实际Hash和POW的目标Hash
	 */
	var nonce int64 = 0
	var hash [32]byte
	for {
		//(1)根据nonce获取拼接数组
		dataByte := pow.prepareData(nonce)
		//(2)生成hash
		hash = sha256.Sum256(dataByte)
		//循环打印出计算的Hash值，不换行
		fmt.Printf("\rNonce:%d Hash:%x", nonce, hash)
		//(3)将生成的Hash转换成big.Int类型
		hashBigInt := new(big.Int)
		hashBigInt.SetBytes(hash[:])
		//(4)比较生成的hash与目标Hash进行对比，需要满足Target > hashBigInt
		if pow.Target.Cmp(hashBigInt) == 1 {
			fmt.Println("挖矿成功")

			break
		}
		nonce ++

	}
	return hash[:], nonce

}

//根据Nonce，获取POW中要验证的block拼接成的数组的数据
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	//1. 根据nonce,生成POW中要验证的block的数组
	data := bytes.Join([][]byte{
		IntToHex(pow.Block.Height),
		pow.Block.PrevBlockHash,
		IntToHex(pow.Block.TimeStamp),
		pow.Block.Data,
		IntToHex(nonce),
		IntToHex(TargetBit),
	}, []byte{})
	return data
}

//定义一个方法用于验证区块是否合法
func (pow *ProofOfWork) IsValid() bool {
	hashBigInt := new(big.Int)
	hashBigInt.SetBytes(pow.Block.Hash)
	if pow.Target.Cmp(hashBigInt) == 1 {
		return true
	}
	return false

}
