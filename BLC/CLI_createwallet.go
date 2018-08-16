package BLC

import "fmt"

func (cli *CLI) CreateWallet() {
	wallets := NewWallets()
	wallets.CreateNewWallet()
	fmt.Println("钱包:", wallets.WalletMap)

}
