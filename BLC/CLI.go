package BLC

import (
	"flag"
	"os"
	"log"
	"fmt"
)

type CLI struct {
	//BlockChain *BlockChain
}

func (cli *CLI) Run() {

	/*
	Usage:
		addblock -data DATA
		printchain


	./bc printchain
		-->执行打印的功能

	 ./bc send -from '["wangergou"]' -to '["lixiaohua"]' -amount '["4"]'
	./bc send -from '["wangergou","rose"]' -to '["lixiaohua","jace"]' -amount '["4","5"]'


	 */
	isValidArgs()

	//1.创建flagset命令对象
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	getAddresslistsCmd:=flag.NewFlagSet("getaddresslists",flag.ExitOnError)
	CreateBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	testMethodCmd:=flag.NewFlagSet("test",flag.ExitOnError)

	//2.设置命令后的参数对象
	flagCreateBlockChainData := CreateBlockChainCmd.String("address", "GenesisBlock", "创世区块的信息")

	flagSendFromData := sendCmd.String("from", "", "转账源地址")
	flagSendToData := sendCmd.String("to", "", "转账目标地址")
	flagSendAmountData := sendCmd.String("amount", "", "转账金额")

	flagGetBalanceData := getBalanceCmd.String("address", "", "要查询余额的账户")

	//3.解析
	switch os.Args[1] {
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := CreateBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getaddresslists":
		err := getAddresslistsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "test":
		err := testMethodCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)

	}
	//4.根据终端输入的命令执行对应的功能
	if sendCmd.Parsed() {
		//fmt.Println("添加区块。。。",*flagAddBlockData)
		if *flagSendFromData == "" || *flagSendToData == "" || *flagSendAmountData == "" {
			fmt.Println("转账信息有误。。")
			printUsage()
			os.Exit(1)
		}
		//添加区块
		//cli.AddBlockToBlockChain([]*Transaction{})
		//from:=*flagSendFromData
		//to:=*flagSendToData
		//amount:=*flagSendAmountData
		from := JSONToArray(*flagSendFromData)     //[]string
		to := JSONToArray(*flagSendToData)         //[]string
		amount := JSONToArray(*flagSendAmountData) //[]string
		for i := 0; i < len(from); i++ {
			if !IsValidAddress([]byte(from[i])) || !IsValidAddress([]byte(to[i])) {
				fmt.Println("地址无效，无法转账。。")
				printUsage()
				os.Exit(1)
			}
		}
		//fmt.Println(from)
		//fmt.Println(to)
		//fmt.Println(amount)
		cli.Send(from, to, amount)
	}

	if printChainCmd.Parsed() {
		//fmt.Println("打印区块。。。")
		//cli.BlockChain.PrintChains()
		cli.PrintChains()
	}

	//添加创世区块的创建
	if CreateBlockChainCmd.Parsed() {
		//if *flagCreateBlockChainData == "" {
		if !IsValidAddress([]byte(*flagCreateBlockChainData)) {
			fmt.Println("地址无效，无法创建创世前区块。。")
			printUsage()
			os.Exit(1)
		}
		cli.CreateBlockChain(*flagCreateBlockChainData)
	}

	if getBalanceCmd.Parsed() {
		//if *flagGetBalanceData == "" {
		if !IsValidAddress([]byte(*flagGetBalanceData)) {
			fmt.Println("查询地址有误。。")
			printUsage()
			os.Exit(1)
		}
		cli.GetBalance(*flagGetBalanceData)
	}

	if createWalletCmd.Parsed() {
		//创建钱包--->交易地址
		cli.CreateWallet()
	}


	if getAddresslistsCmd.Parsed(){
		cli.GetAddressLists()
	}

	if testMethodCmd.Parsed(){
		cli.TestMethod()
	}

}

//判断终端输入的参数的长度
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

//添加程序运行的说明
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreatewallet -- 创建钱包")
	fmt.Println("\tgetaddresslists -- 获取所有的钱包地址")
	fmt.Println("\tcreateblockchain -address DATA -- 创建创世区块")
	fmt.Println("\tsend -from From -to To -amount Amount -- 转账交易")
	fmt.Println("\tprintchain -- 打印区块")
	fmt.Println("\tgetbalance -address Data -- 查询余额")
	fmt.Println("\ttest -- 重置")
}

//func (cli *CLI) AddBlockToBlockChain(txs []*Transaction) {
//	//cli.BlockChain.AddBlockToBlockChain(data)
//	bc := GetBlockChainObject()
//	if bc == nil {
//		fmt.Println("没有BlockChain，无法添加新的区块。。")
//		os.Exit(1)
//	}
//	defer bc.DB.Close()
//	bc.AddBlockToBlockChain(txs)
//}
