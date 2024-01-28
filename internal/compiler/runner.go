package compiler

import (
	"errors"
	"fmt"
)

type Opts struct {
}

type Reader interface {
	HasNext() bool
	Next() (string, error)
}

func NewOpts(args []string) (Opts, error) {
	return Opts{}, nil
}

func NewReader(opts Opts) (Reader, error) {
	return nil, errors.New("Not implemented")
}

func Run(reader Reader, opts Opts) error {
	for reader.HasNext() {
		err := doRun(reader, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func doRun(reader Reader, opts Opts) error {
	input, err := reader.Next()
	if err != nil {
		return errors.New("Failed to read")
	}
	tokens, err := Scan(input)
	if err != nil {
		return err
	}
	node, err := Parse(tokens)
	if err != nil {
		return err
	}
	err = Interpret(node)
	if err != nil {
		return err
	}
	fmt.Println("Result:", tokens)
	return nil
}
