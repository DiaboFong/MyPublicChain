package BLC

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"bytes"
)

//1.定义一个钱包结构:Wallet

type Wallet struct {
	//1.私钥
	PrivateKey ecdsa.PrivateKey

	//2.公钥
	PublicKey []byte //公钥
}

//2.产生一对密钥

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	/*
		1.根据椭圆曲线算法，产生随机私钥
		2.根据私钥，产生公钥
		椭圆：ellipse，
		曲线：curve，

		椭圆曲线加密：(ECC：ellipse curve Cryptography)，非对称加密
			加密：
			对称加密和非对称机密

			SECP256K1,算法

			x轴(32byte)，y轴(32byte)--->

		 */

	//椭圆加密
	curve := elliptic.P256() //根据椭圆加密算法，得到一个椭圆曲线值
	//得到私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	//产生公钥
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	return *privateKey, publicKey
}

//3.创建钱包对象
func NewWallet() *Wallet {
	privateKey, publicKey := newKeyPair()
	return &Wallet{privateKey, publicKey}
}

//定义版本号
const version = byte(0x00)
const addressCheckSumLen = 4

//4.根据公钥获取对应的地址 https://www.blockchain.com/btc/address/
//私钥 -- >公钥(原始公钥) -->sha256(hash) -->ripemd160(hash) -->hash值:公钥Hash :pubkeyHash
func (w *Wallet) GetAdrress() []byte {
	/*
	1. 原始公钥-->sha256 -->160 --->公钥Hash
	2. 版本号+公钥Hash -->校验码
	3. 版本号+公钥Hash + 校验码 -->Base58编码
	 */

	//1.得到公钥Hash
	pubKeyHash := PubKeyHash(w.PublicKey)
	//2. 添加版本号: 比特币中为0
	version_payload := append([]byte{version}, pubKeyHash...)
	//3.根据versioned_payload -- >两次sha256，取前四位，得到checksum
	checkSumByte := CheckSum(version_payload)
	//4. 拼接所有的数据
	full_payload := append(version_payload, checkSumByte...)
	//5.Base58编码
	address := Base58Encode(full_payload)

	return address
}

//原始公钥-->公钥Hash
// 1. sha256
// 2. ripemd160
func PubKeyHash(publicKey []byte) []byte {
	//1.sha256
	hasher256 := sha256.New()
	hasher256.Write(publicKey)
	hash256 := hasher256.Sum(nil)
	//2.ripemd160
	hasher160 := ripemd160.New()
	hasher160.Write(hash256)
	hash160 := hasher160.Sum(nil)
	//返回
	return hash160

}

//定义一个函数，用于产生校验码
//进行两次sha256 Hash
func CheckSum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:addressCheckSumLen]
}


//校验地址是否有效
func IsVaildAddress(address []byte)bool{

	//1. Base58解码
	// version+pubkeyHash + checksum
	full_payload := Base58Decode(address) //25
	//2. 获取地址中携带的CheckSum
	checkSumBytesOld := full_payload[len(full_payload)-addressCheckSumLen:] //[21:]
	version_payload := full_payload[:len(full_payload)-addressCheckSumLen] //[:21]

	//3.使用version_payload生成一次校验码
	checkSumBytesNew := CheckSum(version_payload)

	//4. 比较checkSumBytesOld与checkSumBytesNew
	return bytes.Compare(checkSumBytesOld,checkSumBytesNew) == 0

}