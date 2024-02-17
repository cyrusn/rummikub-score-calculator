package cmd

import (
	"cyrusn/rummikub/db"
	"flag"
	"fmt"
	"os"
)

var setupCmd = flag.NewFlagSet("setup", flag.ExitOnError)

func Setup() {
	var noOfPlayers = setupCmd.Int("n", 4, "Number of Players")
	var filePath = setupCmd.String("f", DEFAULT_FILE_PATH, FILE_PATH_HELP)
	setupCmd.Parse(os.Args[2:])

	sqlDb := Connect(*filePath, true)

	os.Remove(*filePath)

	err := db.CreateTable(sqlDb)
	handleError(err)

	for i := 0; i < *noOfPlayers; i++ {
		fmt.Println("Enter player", i+1, "name:")

		var playerName string
		fmt.Scanln(&playerName)

		err := db.InsertPlayer(sqlDb, playerName)
		handleError(err)
	}
}
