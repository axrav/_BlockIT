package main

import (
	"blockit/handlers"
	"blockit/router"
	"log"
)

func main() {
	router := router.NewBlockRouter()
	go handlers.Block.Background()
	log.Fatal(router.Listen(":3000"))
}
