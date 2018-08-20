package BLC

import (
	"flag"
	"os"
	"log"
	"fmt"
)

type CLI struct {

}

func (cli *CLI) Run() {

/*
Usage:
        createwallet
                        -- 创建钱包
        getaddresslists
                        -- 获取所有的钱包地址
        createblockchain -address address
                        -- 创建创世区块
        send -from SourceAddress -to DestAddress -amount Amount
                        -- 转账交易
        printchain
                        -- 打印区块
        getbalance -address address
                        -- 查询余额

*/
	isValidArgs()

	//1.创建flagset命令对象
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	getAddresslistsCmd := flag.NewFlagSet("getaddresslists", flag.ExitOnError)
	CreateBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	testMethodCmd := flag.NewFlagSet("test", flag.ExitOnError)

	//2.设置命令后的参数对象
	flagCreateBlockChainData := CreateBlockChainCmd.String("address", "GenesisBlock", "创世区块的信息")
	flagSendFromData := sendCmd.String("from", "", "转账源地址")
	flagSendToData := sendCmd.String("to", "", "转账目标地址")
	flagSendAmountData := sendCmd.String("amount", "", "转账金额")
	flagGetBalanceData := getBalanceCmd.String("address", "", "要查询余额的账户")

	//3.解析命令对象
	switch os.Args[1] {
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
	case "createblockchain":
		err := CreateBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
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
	//4.1 创建钱包--->交易地址
	if createWalletCmd.Parsed() {
		cli.CreateWallet()
	}
	//4.2 获取钱包地址
	if getAddresslistsCmd.Parsed() {
		cli.GetAddressLists()
	}
	//4.3 创建创世区块
	if CreateBlockChainCmd.Parsed() {
		if !IsValidAddress([]byte(*flagCreateBlockChainData)) {
			fmt.Println("地址无效，无法创建创世前区块")
			printUsage()
			os.Exit(1)
		}
		cli.CreateBlockChain(*flagCreateBlockChainData)
	}
	//4.4 转账交易
	if sendCmd.Parsed() {
		if *flagSendFromData == "" || *flagSendToData == "" || *flagSendAmountData == "" {
			fmt.Println("转账信息有误")
			printUsage()
			os.Exit(1)
		}
		//添加区块
		from := JSONToArray(*flagSendFromData)     //[]string
		to := JSONToArray(*flagSendToData)         //[]string
		amount := JSONToArray(*flagSendAmountData) //[]string
		for i := 0; i < len(from); i++ {
			if !IsValidAddress([]byte(from[i])) || !IsValidAddress([]byte(to[i])) {
				fmt.Println("地址无效，无法转账")
				printUsage()
				os.Exit(1)
			}
		}

		cli.Send(from, to, amount)
	}
	//4.5 查询余额
	if getBalanceCmd.Parsed() {
		if !IsValidAddress([]byte(*flagGetBalanceData)) {
			fmt.Println("查询地址有误")
			printUsage()
			os.Exit(1)
		}
		cli.GetBalance(*flagGetBalanceData)
	}

	//4.6 打印区块信息
	if printChainCmd.Parsed() {
		cli.PrintChains()
	}

	if testMethodCmd.Parsed() {
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
	fmt.Println("\tcreatewallet\n\t\t\t-- 创建钱包")
	fmt.Println("\tgetaddresslists\n\t\t\t-- 获取所有的钱包地址")
	fmt.Println("\tcreateblockchain -address address\n\t\t\t-- 创建创世区块")
	fmt.Println("\tsend -from SourceAddress -to DestAddress -amount Amount\n\t\t\t-- 转账交易")
	fmt.Println("\tprintchain\n\t\t\t-- 打印区块")
	fmt.Println("\tgetbalance -address address\n\t\t\t-- 查询余额")

	//fmt.Println("\ttest -- 重置")
}
