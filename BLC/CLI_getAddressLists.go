package BLC

import "fmt"

func (cli *CLI) GetAddressLists() {
	fmt.Println("打印所有的钱包地址")
	wallets :=NewWallets()
	for address,_ := range wallets.WalletMap {
		fmt.Println("address:",address)
	}
}
