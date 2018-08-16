package BLC

import "fmt"

//定义一个钱包的集合，存储多个钱包对象
type Wallets struct {
	WalletMap map[string]*Wallet
}

//定义一个函数，用于创建一个钱包的集合

func NewWallets() *Wallets {
	wallets := &Wallets{}
	wallets.WalletMap = make(map[string]*Wallet)
	return wallets

}

func (ws *Wallets) CreateNewWallet() {
	wallet := NewWallet()
	address := wallet.GetAdrress()
	fmt.Printf("创建的钱包地址为:%s\n", address)
	ws.WalletMap[string(address)] = wallet
}
