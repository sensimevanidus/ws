package main

import (
	"log"
	"os"

	"github.com/sensimevanidus/repl"
)

func main() {
	if err := repl.Run(os.Stdin, os.Stdout, newWebSocketProcessor()); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
