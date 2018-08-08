package main

import (
	"MyPublicChain/BLC"
	"fmt"
	"encoding/hex"
)

func main() {

	//添加新的区块至DB中

	bc := BLC.CreateBlockChainWithGenesisBlock("brucefeng GenesisBlock")

	bc.AddBlockToBlockChain("第二个区块")
	fmt.Println(hex.EncodeToString(bc.Tip))
	bc.AddBlockToBlockChain("第三个区块")
	fmt.Println(hex.EncodeToString(bc.Tip))
	bc.AddBlockToBlockChain("第四个区块")
	fmt.Println(hex.EncodeToString(bc.Tip))
	defer bc.DB.Close()

}
