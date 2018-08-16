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
	fmt.Println("地址是否有效:", BLC.IsVaildAddress(address))
	fmt.Println("地址是否有效:", BLC.IsVaildAddress([]byte("brucefeng")))
	//测试钱包集合

	wallets := BLC.NewWallets()
	fmt.Println(wallets)
	wallets.CreateNewWallet()
	wallets.CreateNewWallet()
	fmt.Println(wallets)



}
