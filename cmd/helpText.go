package cmd

import "fmt"

func Help() {
	fmt.Println("Usage: rummikub <command>")
	fmt.Println("  setup, s\tSetup the game")
	fmt.Println("  add, a\tAdd a round")
	fmt.Println("  remove, rm\tRemove a round")
	fmt.Println("  result, r\tShow the result")
}
