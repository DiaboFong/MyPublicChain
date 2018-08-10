package BLC

import (
	"os"
	"fmt"
	"flag"
	"log"
)

type CLI struct {
}

func (cli CLI) Run() {

	//1. 创建FlagSet命令对象
	createBlockChainCmd := flag.NewFlagSet("creategenesisblock", flag.ExitOnError)
	//addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	//2.设置命令后的参数对象
	flagCreateBlockChainData := createBlockChainCmd.String("address", "GenesisBlock", "创建带有创世区块的区块链")
	flagSendFromData := sendCmd.String("from", "", "转账源地址")
	flagSendToData := sendCmd.String("to", "", "转账目标地址")
	flagSendAmountData := sendCmd.String("amount", "", "转账金额")
	flagGetBalanceData := getBalanceCmd.String("address", "", "需要查询的余额账户")

	//3.解析
	switch os.Args[1] {
	case "creategenesisblock":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
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
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		printUsage()
		os.Exit(1)
	}
	//4.根据终端输入的命令执行对应的功能
	if createBlockChainCmd.Parsed() {
		fmt.Println()
		if *flagCreateBlockChainData == "" {
			printUsage()
		}
		cli.CreateBlockChainWithGenesisBlock(*flagCreateBlockChainData)

	}

	if sendCmd.Parsed() {

		if *flagSendFromData == "" || *flagSendToData == "" || *flagSendAmountData == "" {
			fmt.Println("转账信息有误")
			printUsage()
			os.Exit(1)
		}
		//添加区块
		//cli.AddBlockToBlockChain([]*Transaction{})
		fromData := JSONToArray(*flagSendFromData)
		toData := JSONToArray(*flagSendToData)
		amountData := JSONToArray(*flagSendAmountData)
		//测试 ./bc send -from '["brucefeng"]' 			   -to '["jimenghao"]'  -amount '["10"]'
		//测试 ./bc send -from '["brucefeng","jimenghao"]'  -to '["jack","tom"]' -amount '["10","20"]'

		//fmt.Printf("send -from %s -to %s -amount %s\n", fromData, toData, amountData)
		cli.Send(fromData, toData, amountData)

	}
	if printChainCmd.Parsed() {
		fmt.Println("打印区块信息")
		cli.PrintChains()
	}

	if getBalanceCmd.Parsed() {
		if *flagGetBalanceData == "" {
			fmt.Println("查询地址为空")
			printUsage()
			os.Exit(1)
		}
		cli.GetBalance(*flagGetBalanceData)
	}

}

/*
	Usage:
		addblock -data DATA
		printchain

	./bc printchain  //执行打印的功能
	./bc addblock -data "send 100BTC To Brucefeng"
	 */

//判断终端输入的参数长度
func IsValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

//添加程序运行说明

func printUsage() {
	/*
		Usage:
		addblock -data DATA
		printchain
	 */
	fmt.Println("Uasge:")
	fmt.Println("\tcreategenesisblock -address DATA --添加创世区块")
	//fmt.Println("\taddblock -data DATA --添加区块")
	fmt.Println("\tsend -from SourceAddress - to TargetAddress -amount Amount --转账交易")
	fmt.Println("\tprintchain --打印区块")
	fmt.Println("\tgetbalance -address DATA --查询余额")

}

func (cli CLI) CreateBlockChainWithGenesisBlock(address string) {
	bc := GetBlockChainObject()
	if bc == nil {
		//如果bc为空，说明并没有创世区块,此处不需要关闭DB，因为没有被Open
		//defer bc.DB.Close()
		CreateGenesisBlockToDB(address)

	} else {
		os.Exit(1)
	}
}

func (cli CLI) AddBlockToBlockChain(txs []*Transaction) {
	bc := GetBlockChainObject()

	if bc == nil {
		//如果bc为空，说明并没有创世区块
		fmt.Println("创世区块不存在")
		printUsage()
		os.Exit(1)
	} else {
		defer bc.DB.Close()
		bc.AddBlockToBlockChain(txs)

	}

}

func (cli CLI) PrintChains() {
	bc := GetBlockChainObject()

	if bc == nil {
		//如果bc为空，说明并没有创世区块
		fmt.Println("创世区块不存在")
		printUsage()
		os.Exit(1)
	} else {
		defer bc.DB.Close()
		bc.PrintChains()

	}
}

func (cli *CLI) Send(from, to, amount []string) {
	bc := GetBlockChainObject()
	if bc == nil {
		fmt.Println("没有BlockChain，无法转账")
		os.Exit(1)
	}
	defer bc.DB.Close()
	bc.MineNewBlock(from, to, amount)
}

func (cli CLI) GetBalance(address string) {
	bc := GetBlockChainObject()
	if bc == nil {
		fmt.Println("没有BlockChain，无法查询")
		os.Exit(1)
	}
	defer bc.DB.Close()
	total := bc.GetBalance(address)
	fmt.Printf("%s,余额是:%d\n", address, total)
}
