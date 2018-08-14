package BLC

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

//工作量证明：

const TargetBit = 16 //目标哈希的0个个数,16,20,24,28

type ProofOfWork struct {
	Block  *Block   //要验证的block
	Target *big.Int //目标hash
}

func NewProofOfWork(block *Block) *ProofOfWork {
	//1.创建pow对象
	pow := &ProofOfWork{}
	//2.设置属性值
	pow.Block = block
	target := big.NewInt(1)           // 目标hash，初始值为1
	target.Lsh(target, 256-TargetBit) //左移256-16
	pow.Target = target
	/*
	hash：256bit
	16进制：32个
	4个0--->16个0
	0000 0000 0000 0000 1000 0000 0000 0000    256
	 */

	/*
	0000 0001
	0010 0000

	8-2
	256-16
	 */
	 return pow

}

//设计一个函数，得到有效hash，nonce
func (pow *ProofOfWork) Run() ([]byte, int64) {
	//挖矿--->更改nonce的值，计算hash，直到小于目标hash。
	/*
	思路：
	1.设置nonce值：0,1,2.......
	2.block-->拼接数组，产生hash
	3.比较实际hash和pow的目标hash，
	 */
	var nonce int64 = 0
	var hash [32]byte
	for {
		//1.根据nonce获取数据
		data := pow.prepareData(nonce)
		//2.生成hash
		hash = sha256.Sum256(data) //[32]byte
		fmt.Printf("\r%d,%x",nonce,hash)
		//3.验证：和目标hash比较
		/*
		func (x *Int) Cmp(y *Int) (r int)
		Cmp compares x and y and returns:

		   -1 if x <  y
			0 if x == y
		   +1 if x >  y

		目的：target > hashInt,成功
		 */
		hashInt := new(big.Int)
		hashInt.SetBytes(hash[:])

		if pow.Target.Cmp(hashInt) == 1 {
			break
		}

		//if hashInt.Cmp(pow.Target) == -1{
		//
		//}

		nonce++
	}
	fmt.Println()
	return hash[:], nonce

}

//根据nonce，获取pow中要验证的block拼接成的数组的数据
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	//1.根据nonce，生成pow中要验证的block的数组
	data := bytes.Join([][]byte{
		IntToHex(pow.Block.Height),
		pow.Block.PrevBlockHash,
		IntToHex(pow.Block.TimeStamp),
		pow.Block.HashTransactions(),
		IntToHex(nonce),
		IntToHex(TargetBit),
	}, []byte{})
	return data

}




//提供一个方法：
func (pow *ProofOfWork) IsValid()bool{
	hashInt :=new(big.Int)
	hashInt.SetBytes(pow.Block.Hash)
	return pow.Target.Cmp(hashInt) == 1
}
