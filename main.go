package main

import (
	"MyPublicChain/BLC"
	"fmt"
)

func main() {

	block := BLC.NewBlock("This is a Genesis Block", make([]byte, 3, 3), 0)
	fmt.Println(block)
}
