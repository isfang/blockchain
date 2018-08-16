package main

import (
	"flag"
	"blockchain/block"
	"os"
	"fmt"
	"strconv"
	"strings"
)
var action = flag.String("a", "", "cg: create a genesis chain; show: show chain; cost: from/to/amount")
var data = flag.String("d", "", "address for user")

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

	//blockChain := block.NewBlockChain()
	//defer blockChain.BlotDB.Close()
	//
	//
	//switch *action {
	//case "add":
	//
	//	if *data == "" {
	//		fmt.Println("error data")
	//		return
	//	} else {
	//		fmt.Println("add block with data", *data)
	//		blockChain.AddBlock(*data)
	//	}
	//case "show":
	//
	//	fmt.Println("show chain")
	//	bci := blockChain.Iterator()
	//
	//	for {
	//		b := bci.Next()
	//
	//		fmt.Printf("Prev hash: %x\n", b.PreHash)
	//		fmt.Printf("Hash: %x\n", b.Hash)
	//		p := block.NewWorkProof(b)
	//		fmt.Printf("PoW: %s\n", strconv.FormatBool(p.Validate()))
	//		fmt.Println()
	//
	//		if len(b.PreHash) == 0 {
	//			break
	//		}
	//	}
	//}
	switch *action {
	case "cg":
		fmt.Println("create a new chain with a genesis block.")
		if *data == "" {
			fmt.Println("use -d setup address")
			os.Exit(1)
		}
		bc := block.CreateBlockchain(*data)
		defer bc.BlotDB.Close()

	case "show":
		fmt.Println("show chain")

		 bc := block.NewBlockChain("")
		 defer  bc.BlotDB.Close()

		 itrator := bc.Iterator()

		for {
			b := itrator.Next()

			fmt.Printf("Prev hash: %x\n", b.PreHash)
			fmt.Printf("Hash: %x\n", b.Hash)
			pow := block.NewWorkProof(b)
			fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
			fmt.Println()
			if len(b.PreHash) == 0 {
				fmt.Println("is end.")
				break
			}

		}
	case "cost":

		fmt.Println("cost")

		if *data == "" {
			fmt.Println("-d format with from/to/amount")
			os.Exit(1)
		}

		params := strings.Split(*data, "/")

		if len(params) < 3 {
			fmt.Println("-d format with from/to/amount")
			os.Exit(1)
		}

		from := params[0]
		to := params[1]
		amount := params[2]

		fmt.Println("from is ", from, " to is ", to, " amount is ", amount)

		bc := block.NewBlockChain(from)
		defer bc.BlotDB.Close()



	}

}
