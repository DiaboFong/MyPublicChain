package main

import (
	"MyPublicChain/BLC"
	"fmt"
)

func main() {

	wallet := BLC.NewWallet()
	fmt.Println(wallet.PublicKey)
	fmt.Println(wallet.PrivateKey)

	address := wallet.GetAdrress()
	fmt.Println(address)
	fmt.Println(string(address))
}
