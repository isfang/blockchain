package main

import (
	"flag"
	"blockchain/block"
	"fmt"
	"strconv"
)
var action = flag.String("s", "", "add-add block, show-show chain")
var data = flag.String("d", "", "data for block")

func main() {
	flag.Parse()
	//blockChain := block.NewBlockChain()
	//
	//blockChain.AddBlock("block-1")
	//blockChain.AddBlock("block-2")
	//blockChain.AddBlock("block-3")
	//
	//for _, b := range  blockChain.Blocks {
	//	fmt.Printf("Prev hash: %x\n", b.PreHash)
	//	fmt.Printf("Data: %s\n", b.Data)
	//	fmt.Printf("Hash: %x\n", b.Hash)
	//	pow := block.NewWorkProof(b)
	//	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	//	fmt.Println()
	//}

	//blockChain := block.NewBlockChain()
	//defer blockChain.BlotDB.Close()
	//
	//c := cli.CLI{blockChain}
	//c.Run()

	blockChain := block.NewBlockChain()
	defer blockChain.BlotDB.Close()


	switch *action {
	case "add":

		if *data == "" {
			fmt.Println("error data")
			return
		} else {
			fmt.Println("add block with data", *data)
			blockChain.AddBlock(*data)
		}
	case "show":

		fmt.Println("show chain")
		bci := blockChain.Iterator()

		for {
			b := bci.Next()

			fmt.Printf("Prev hash: %x\n", b.PreHash)
			fmt.Printf("Hash: %x\n", b.Hash)
			p := block.NewWorkProof(b)
			fmt.Printf("PoW: %s\n", strconv.FormatBool(p.Validate()))
			fmt.Println()

			if len(b.PreHash) == 0 {
				break
			}
		}
	}

}
