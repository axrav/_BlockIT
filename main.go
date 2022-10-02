package main

import (
	"blockit/router"
	"blockit/utils"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	router := router.NewBlockRouter()
	blockchain := *utils.CurrentBlock()
	go func() {
		for _, block := range blockchain.Blocks {
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("PrevHash: %s\nData:%s\nHash: %s\nTimestamp: %s", block.PrevHash, string(bytes), block.Hash, block.TimeStamp)
		}
	}()
	log.Fatal(router.Listen(":3000"))
}
