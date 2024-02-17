package cmd

import (
	"cyrusn/rummikub/db"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Mark struct {
	player_id int
	score     int
}

var addCmd = flag.NewFlagSet("add", flag.ExitOnError)

func Add() {
	var filePath = addCmd.String("f", DEFAULT_FILE_PATH, FILE_PATH_HELP)
	var helpText = addCmd.Bool("h", false, "Show help text")
	addCmd.Parse(os.Args[2:])

	if *helpText {
		fmt.Println(`Usage of add:
    Enter tiles for each player. The tiles should be comma separated.`)
		os.Exit(0)
	}

	sqlDb := Connect(*filePath, false)

	players, err := db.GetPlayers(sqlDb)
	handleError(err)

	lastRound, err := db.GetLastRound(sqlDb)
	handleError(err)

	var round = lastRound + 1
	fmt.Println("Round:", round)
	fmt.Println("----------")
	var marks []Mark

	var winnerMark int
	for _, player := range players {
		fmt.Println("Players:", player.Username)
		var tiles string

		fmt.Print("Tiles:")
		fmt.Scanln(&tiles)

		var scoreStrings = strings.Split(tiles, ",")
		score := sum(scoreStrings...)
		winnerMark += score
		marks = append(marks, Mark{player.Id, -score})
	}

	for _, mark := range marks {
		score := mark.score
		if score == 0 {
			score = winnerMark
		}
		err = db.InsertScore(sqlDb, round, mark.player_id, score)
		handleError(err)
	}

	scores, err := db.GetScoresByRound(sqlDb, round)
	handleError(err)

	fmt.Println("----------")

	for _, score := range scores {
		fmt.Println(score.Username+":", score.Mark)
	}

}
