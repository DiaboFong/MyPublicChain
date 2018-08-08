package main

import (
	"MyPublicChain/BLC"
	"fmt"
)

func main() {

	bc := BLC.CreateBlockChainWithGenesisBlock("brucefeng GenesisBlock")
	fmt.Println(bc)

}
