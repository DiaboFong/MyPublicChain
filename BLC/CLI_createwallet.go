package BLC

import "fmt"

func (cli *CLI) CreateWallet() {
	_, wallets := GetWallets() //获取钱包集合对象
	/*
	1. 取出钱包集合对象
	2. 删除一个map集合
	3. 调用CreateNewWallets将新生成的wallet对象存入值map集合中
	4. 通过SaveFile存入本地文件
	delete为测试代码
	 */
	//delete(wallets.WalletMap,"1MaU1Q1LFwxHAbx4gfkVYSpEs2qDBmbV7c")
	wallets.CreateNewWallets()
	fmt.Println("钱包：", wallets.WalletMap)
}
