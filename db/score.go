package db

import "database/sql"

type Score struct {
	Username string
	Round    int
	Mark     int
}

func InsertScore(db *sql.DB, round int, player_id int, score int) error {
	_, err := db.Exec("INSERT INTO scores (round, player_id, score) VALUES (?, ?, ?)", round, player_id, score)
	if err != nil {
		return err
	}
	return nil
}

func GetScoresByRound(db *sql.DB, round int) ([]Score, error) {
	rows, err := db.Query(`
    SELECT
      username, SUM(score) 
    FROM scores
    JOIN players ON 
      player_id = players.Id 
    WHERE round = ? 
    GROUP BY
      player_id, round;`,
		round,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var scores []Score
	for rows.Next() {
		var score Score
		score.Round = round
		err = rows.Scan(&score.Username, &score.Mark)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func RemoveRoundScores(db *sql.DB, round int) error {
	_, err := db.Exec("DELETE FROM scores WHERE round = ?", round)
	return err
}
