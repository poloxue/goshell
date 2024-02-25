package main

import (
	"errors"
	"os"
)

type BuiltinCmder interface {
	Execute(args ...string) error
}

type ChangeDirCommand struct{}

func (c *ChangeDirCommand) Execute(args ...string) error {
	if len(args) < 1 {
		return errors.New("Expected path argument")
	}
	return os.Chdir(args[0])
}
