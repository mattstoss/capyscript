package main

import (
	"log"
	"os"

	"github.com/mattstoss/capyscript/internal/compiler"
)

func main() {
	opts, err := compiler.NewOpts(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	reader, err := compiler.NewReader(opts)
	if err != nil {
		log.Fatal(err)
	}
	err = compiler.Run(reader, opts)
	if err != nil {
		log.Fatal(err)
	}
}
