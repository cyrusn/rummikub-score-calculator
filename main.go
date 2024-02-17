package main

import (
	"cyrusn/rummikub/cmd"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Show help text if subcommand is not given.
	if len(os.Args) < 2 {
		cmd.Help()
		os.Exit(0)
	}

	mode := os.Args[1]
	switch mode {

	case "setup", "s":
		cmd.Setup()

	case "add", "a":
		cmd.Add()

	case "result", "r":
		cmd.GetResult()

	case "remove", "rm":
		cmd.Remove()

	default:
		cmd.Help()
	}
}
