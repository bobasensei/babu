package main

import (
	"os"

	"github.com/bobasensei/babu/cmd"
)

func main() {
	if err := cmd.Cmd().Execute(); err != nil {
		os.Exit(1)
	}
}
