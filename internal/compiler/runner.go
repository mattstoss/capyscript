package compiler

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Opts struct {
	Input   io.Reader
	Command string

	// run options
	RunMode string
}

func ParseOpts(args []string) (Opts, error) {
	if len(args) < 2 {
		return Opts{}, errors.New("usage: capyscript run myfile.capy")
	}
	command := args[1]
	if command == "run" {
		return parseRunOpts(args[2:])
	}
	return Opts{}, errors.New("unrecognized command")
}

func Run(opts Opts) error {
	bytes, err := io.ReadAll(opts.Input)
	if err != nil {
		return errors.New("Failed to read")
	}
	input := string(bytes)
	tokens, err := Scan(input)
	if err != nil {
		return err
	}
	if opts.RunMode == "debug_scanner" {
		fmt.Println(tokens)
		return nil
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

func parseOpts(args []string) (map[string]string, error) {
	result := make(map[string]string)
	for _, opt := range args {
		parts := strings.Split(opt, "=")
		if len(parts) != 2 {
			return result, fmt.Errorf("invalid opt format: %s", opt)
		}
		key := parts[0]
		val := parts[1]
		key, exists := strings.CutPrefix(key, "--")
		if !exists {
			return result, fmt.Errorf("invalid opt format: %s", opt)
		}
		_, exists = result[key]
		if exists {
			return result, fmt.Errorf("opt defined more than once: --%s", key)
		}
		result[key] = val
	}
	return result, nil
}

func parseRunOpts(args []string) (Opts, error) {
	if len(args) < 1 {
		return Opts{}, errors.New("run command must have at least one argument")
	}
	filename := args[0]
	file, err := os.Open((filename))
	if err != nil {
		return Opts{}, err
	}
	opts, err := parseOpts(args[1:])
	if err != nil {
		return Opts{}, err
	}
	runMode := "normal"
	for key, val := range opts {
		switch key {
		case "mode":
			if val == "normal" {
				runMode = val
			} else if val == "debug_scanner" {
				runMode = "debug_scanner"
			} else {
				return Opts{}, errors.New("Unrecognized mode opt")
			}
		default:
			return Opts{}, errors.New("unrecognized run opt")
		}
	}
	runOpts := Opts{
		Input:   bufio.NewReader(file),
		Command: "run",
		RunMode: runMode,
	}
	return runOpts, nil
}
