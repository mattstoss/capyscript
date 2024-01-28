package main

import (
	"fmt"
	"os"

	"github.com/mattstoss/capyscript/internal/compiler"
)

func main() {
	opts, err := compiler.ParseOpts(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = compiler.Run(opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
