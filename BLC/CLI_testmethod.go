package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) TestMethod(){
	bc :=GetBlockChainObject()
	if bc == nil{
		fmt.Println("没有数据库，无法获取utxo。。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	unSpentUTXOsMap:=bc.FindUnspentUTXOMap()
	fmt.Println("长度：",len(unSpentUTXOsMap))
	for txIDStr,txoutputs:=range unSpentUTXOsMap{
		fmt.Println("交易ID：",txIDStr)
		for _,utxo:=range txoutputs.UTXOs{
			//utxo:txid,index,output
			//output:value,公钥哈希-->地址
			fmt.Println("\t金额：",utxo.Output.Value)
			fmt.Printf("\t地址%s：\n",GetAddressByPubKeyHash(utxo.Output.PubKeyHash))
			fmt.Println("---------------------------------")
		}
	}

	utxoSet:=&UTXOSet{bc}
	utxoSet.ResetUTXOSet()
}
