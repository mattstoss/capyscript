package compiler

import (
	"errors"
	"fmt"
	"io"
)

type Opts struct {
}

func NewOpts(args []string) (Opts, error) {
	return Opts{}, nil
}

func NewReader(opts Opts) (io.Reader, error) {
	return nil, errors.New("Not implemented")
}

func Run(r io.Reader, opts Opts) error {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return errors.New("Failed to read")
	}
	input := string(bytes)
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
