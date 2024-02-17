package db

import "database/sql"

type Player struct {
	Id       int
	Username string
}

func InsertPlayer(db *sql.DB, playerName string) error {
	_, err := db.Exec("INSERT INTO Players(username) VALUES (?)", playerName)
	if err != nil {
		return err
	}
	return nil
}

func GetPlayers(db *sql.DB) ([]Player, error) {
	rows, err := db.Query("SELECT id, username FROM Players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var players []Player

	for rows.Next() {
		var player Player
		err = rows.Scan(&player.Id, &player.Username)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}
