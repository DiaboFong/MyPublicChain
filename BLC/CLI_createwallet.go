package BLC

func (cli *CLI) CreateWallet() {
	wallets := GetWallets() //获取钱包集合对象
	/*
	1. 取出钱包集合对象
	2. 删除一个map集合
	3. 调用CreateNewWallets将新生成的wallet对象存入值map集合中
	4. 通过SaveFile存入本地文件
	delete为测试代码

	//每次调用一次createwallet后，都会变化，因为存储跟取的时候都是以指针的方式，对象并不一样了。
	map[1DFJmH6cDmvZCHPL1oX9BdSLR6cqpkFtnR:0xc42001b700]
	map[1DFJmH6cDmvZCHPL1oX9BdSLR6cqpkFtnR:0xc420099dc0 14Ze5NCHLcNiYrEmuXwoe9VFKaYxNsizfL:0xc4200d4140]
	map[1DFJmH6cDmvZCHPL1oX9BdSLR6cqpkFtnR:0xc42001be40 14Ze5NCHLcNiYrEmuXwoe9VFKaYxNsizfL:0xc4200be1c0 1J71MJHJ8yQtXbEvNnDkp5yqwwLLwma:0xc4200be240] 		      map[1DFJmH6cDmvZCHPL1oX9BdSLR6cqpkFtnR:0xc4200be240 4Ze5NCHLcNiYrEmuXwoe9VFKaYxNsizfL:0xc42001be40  11J71MJHJ8yQtXbEvNnDkp5yqwwLLwma:0xc4200be1c0
1Lu1RXh2QefEXBcQTKTbpk3SKCrbzKjMDH:0xc4200be2c0]
	验证网站
	https://www.blockchain.com/btc/address/1Lu1RXh2QefEXBcQTKTbpk3SKCrbzKjMDH
	 */
	//delete(wallets.WalletMap,"1MaU1Q1LFwxHAbx4gfkVYSpEs2qDBmbV7c")
	wallets.CreateNewWallets()
	//fmt.Println("钱包：", wallets.WalletMap)
}
