package cli

import (
	"blockchain/block"
	"fmt"
	"os"
	"flag"
	"log"
)

type CLI struct {
	BC *block.BlockChain
}

const (
	USAGE = `
Usage:
  addblock -data BLOCK_DATA    add a block to the blockchain
  printchain                   print all the blocks of the blockchain
`
	CMD_ADD_BLOCK = "addBlock"
	CMD_PRINT_CHAIN = "printChainCmd"
)

func (c *CLI)pUsage() {
	fmt.Println(USAGE)
}


func (c *CLI)validateArgs()  {

	if len(os.Args) < 2 {
		c.pUsage()
		os.Exit(1)
	}
}

func (c *CLI)Run()  {

	c.validateArgs()

	addBlockCmd := flag.NewFlagSet(CMD_ADD_BLOCK, flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(CMD_PRINT_CHAIN, flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case CMD_ADD_BLOCK:
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case CMD_PRINT_CHAIN:
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		c.pUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		c.BC.AddBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		c.pUsage()
	}
}