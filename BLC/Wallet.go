package BLC

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

//1.定义一个钱包结构:Wallet

type Wallet struct {
	//1.私钥
	PrivateKey ecdsa.PrivateKey

	//2.公钥
	PublicKey []byte
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
				对称加密和非对称机密啊

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
