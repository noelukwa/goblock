package main

import (
	"encoding/hex"
	"flag"
	"os"
	"strconv"
)

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		Handle(err)
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		Handle(err)
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func Handle(err error) {
	if err != nil {
		panic(err)
	}
}

func (cli *CLI) printUsage() {
	println("Usage:")
	println(" addblock -data BLOCK_DATA - add a block to the blockchain")
	println(" printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		println("Prev. hash:", hex.EncodeToString(block.PrevBlockHash))
		println("Data:", string(block.Data))
		println("Hash:", hex.EncodeToString(block.Hash))
		pow := NewProofOfWork(block)
		println("PoW:", strconv.FormatBool(pow.Validate()))
		println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
