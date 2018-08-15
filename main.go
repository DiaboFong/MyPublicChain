package main

import (
	"MyPublicChain/BLC"
	"fmt"
)

func main() {

	wallet := BLC.NewWallet()
	fmt.Println(wallet.PublicKey)
	fmt.Println(wallet.PrivateKey)
}
