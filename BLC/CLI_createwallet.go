package BLC

import "fmt"

func (cli *CLI) CreateWallet(){
	wallets:=NewWallets() //获取钱包集合
	wallets.CreateNewWallet()//创建钱包对象
	fmt.Println("钱包：",wallets.WalletMap)
}
