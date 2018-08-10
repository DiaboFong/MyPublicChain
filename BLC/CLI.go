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
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	//2.设置命令后的参数对象
	flagCreateBlockChainData := createBlockChainCmd.String("data", "GenesisBlock", "创建带有创世区块的区块链")
	flagAddBlockData := addBlockCmd.String("data", "helloworld", "区块的交易数据")

	//3.解析
	switch os.Args[1] {
	case "creategenesisblock":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
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
		cli.CreateBlockChainWithGenesisBlock([]*Transaction{})

	}

	if addBlockCmd.Parsed() {
		fmt.Println("添加区块", *flagAddBlockData)
		if *flagAddBlockData == "" {
			printUsage()
		}
		//添加区块
		cli.AddBlockToBlockChain([]*Transaction{})

	}
	if printChainCmd.Parsed() {
		fmt.Println("打印区块信息")
		cli.PrintChains()
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
	fmt.Println("\tcreategenesisblock -data DATA --添加创世区块")
	fmt.Println("\taddblock -data DATA --添加区块")
	fmt.Println("\tprintchain --打印区块")

}

func (cli CLI) CreateBlockChainWithGenesisBlock(txs []*Transaction) {
	bc := GetBlockChainObject()
	defer bc.DB.Close()
	if bc == nil {
		//如果bc为空，说明并没有创世区块
		CreateGenesisBlockToDB(txs)
	} else {
		os.Exit(1)
	}
}

func (cli CLI) AddBlockToBlockChain(txs []*Transaction) {
	bc := GetBlockChainObject()
	defer bc.DB.Close()
	if bc == nil {
		//如果bc为空，说明并没有创世区块
		fmt.Println("创世区块不存在")
		printUsage()
		os.Exit(1)
	} else {
		bc.AddBlockToBlockChain(txs)

	}

}

func (cli CLI) PrintChains() {
	bc := GetBlockChainObject()
	defer bc.DB.Close()
	if bc == nil {
		//如果bc为空，说明并没有创世区块
		fmt.Println("创世区块不存在")
		printUsage()
		os.Exit(1)
	} else {
		bc.PrintChains()

	}
}
