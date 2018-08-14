package main

import (
	"publicchain1803/day04_02_Transaction_多笔交易/BLC"
)

func main() {

	//bytes:=make([]byte,3,3)
	//fmt.Println(bytes)
	//1.测试区块
	//block:=BLC.NewBlock("I am  a block",make([]byte,32,32),0)
	//fmt.Println(block)

	//2.测试创世区块
	//genesisBlock:=BLC.CreateGenesisBlock("Genesis Block")
	//fmt.Println(genesisBlock)

	//3.创建一个区块链
	//blockChain:=BLC.CreateBlockChainWithGenesisBlock("Genesis Block")
	//fmt.Println(blockChain)
	//fmt.Println(blockChain.Blocks)
	//fmt.Println(blockChain.Blocks[0])

	//4.测试添加区块
	//blockChain:=BLC.CreateBlockChainWithGenesisBlock("Genesis Block")
	//blockChain.AddBlockToBlockChain("Send 100RMB to wangergou",blockChain.Blocks[len(blockChain.Blocks)-1].Hash,blockChain.Blocks[len(blockChain.Blocks)-1].Height+1)
	//blockChain.AddBlockToBlockChain("Send 100RMB to rose",blockChain.Blocks[len(blockChain.Blocks)-1].Hash,blockChain.Blocks[len(blockChain.Blocks)-1].Height+1)
	//blockChain.AddBlockToBlockChain("Send 1000RMB to lixiaohua",blockChain.Blocks[len(blockChain.Blocks)-1].Hash,blockChain.Blocks[len(blockChain.Blocks)-1].Height+1)
	//fmt.Println(blockChain)
	//
	//
	//pow:=BLC.NewProofOfWork(blockChain.Blocks[0])
	//fmt.Println(pow.IsValid())

	//5.验证block的序列化和反序列化
	//block:= BLC.NewBlock("helloworld",make([]byte,32,32),0)
	//fmt.Println(block)

	//blockBytes:=block.Serialize()
	//fmt.Println(blockBytes)
	//
	//block2:=BLC.DeserializeBlock(blockBytes)
	//fmt.Println(block2)

	//6.测试block存入到数据库中
	//db,err:=bolt.Open("my.db",0600,nil)
	//if err != nil{
	//	log.Panic(err)
	//}
	//defer db.Close()
	//
	////将block存入到数据库中
	//err =db.Update(func(tx *bolt.Tx) error {
	//	b,err:=tx.CreateBucketIfNotExists([]byte("blocks"))
	//	if err !=nil{
	//		log.Panic(err)
	//	}
	//	if b != nil{
	//		err :=b.Put([]byte("l"),block.Serialize())
	//		if err != nil {
	//			fmt.Println("存储失败。。")
	//		}
	//	}
	//	return nil
	//})
	//if err != nil{
	//	fmt.Println("更新数据库失败。。")
	//}
	//
	//
	//err  =db.View(func(tx *bolt.Tx) error {
	//	b:=tx.Bucket([]byte("blocks"))
	//	if b != nil{
	//		blockBytes:=b.Get([]byte("l"))
	//		block:=BLC.DeserializeBlock(blockBytes)
	//		fmt.Println(block)
	//	}
	//
	//
	//	return nil
	//})
	//if err != nil{
	//	fmt.Println("查询失败。。")
	//}

	//7.测试获取blockchain对象

	//blockchain:=BLC.CreateBlockChainWithGenesisBlock("Genesisblock..data...")
	//fmt.Println(blockchain)
	//defer blockchian.DB.Close()
	//8.测试添加新的区块
	//blockchain:=BLC.CreateBlockChainWithGenesisBlock("创世区块的数据")
	//blockchain.AddBlockToBlockChain("send 100RMB to wangergou")
	//blockchain.AddBlockToBlockChain("send 1000RMB to tianzhongtian")
	//blockchain.AddBlockToBlockChain("send 10000RMB to lixiaohua")
	//defer blockchain.DB.Close()

	//9.测试打印所有的区块
	//blockchain.PrintChains()

	//CLI
	//blockchain:=BLC.CreateBlockChainWithGenesisBlock("Genesisblock..data...")
	cli:=BLC.CLI{}
	cli.Run()

}
/*
数据库中：
 */