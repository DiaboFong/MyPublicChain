package main

import "MyPublicChain/BLC"

func main() {

	BLC.IsValidArgs()
	cli := BLC.CLI{}
	cli.Run()

}
