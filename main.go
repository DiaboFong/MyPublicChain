package main

import (
	"MyPublicChain/BLC"
	"fmt"
)

func main() {
	//验证Block的序列化与反序列化

	block := BLC.NewBlock("brucefeng block", make([]byte, 3, 3), 1)

	fmt.Println(block)
	//测试序列化
	blockBytes := block.Serialize()
	fmt.Println(blockBytes)
	//测试反序列化

	block2 := BLC.DeSerializeBlock(blockBytes)
	fmt.Println(block2)

}
