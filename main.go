package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/ssanjeewa/demo/blockchain"
)



func main() {
	chain := InitBlockchain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, b := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", b.PrevHash)
		fmt.Printf("Datam in Block: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Println("-----------------------------")
	}
}
