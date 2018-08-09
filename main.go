package main

import (
	"flag"
	"os"
	"log"
	"fmt"
)

func main() {

	/*
	Usage:
		addblock -data DATA
		printchain

	./bc printchain  //执行打印的功能
	./bc addblock -data "send 100BTC To Brucefeng"
	 */

	isValidArgs()

	//1. 创建FlagSet命令对象
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	//2.设置命令后的参数对象
	flagAddBlockData := addBlockCmd.String("data", "helloworld", "区块的交易数据")

	//3.解析
	switch os.Args[1] {
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
	if addBlockCmd.Parsed() {
		fmt.Println("添加区块", *flagAddBlockData)
	}
	if printChainCmd.Parsed() {
		fmt.Println("打印区块")
	}

}

//判断终端输入的参数长度
func isValidArgs() {
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
	fmt.Println("\taddblock -data DATA --添加区块")
	fmt.Println("\tprintchain --打印区块")

}
