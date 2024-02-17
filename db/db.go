package db

import (
	"database/sql"
)

type Schema struct {
	name    string
	content string
}

const playersSchema = `
CREATE TABLE IF NOT EXISTS Players( 
  id INTEGER PRIMARY KEY AUTOINCREMENT, 
  username TEXT NOT NULL
	);`

const scoresSchema = `
CREATE TABLE IF NOT EXISTS scores (
  id integer primary key autoincrement,
  round integer not null,
  score integer not null,
  player_id integer not null,
  foreign key(player_id) references players(id)
  );

CREATE INDEX player_id_idx on scores (player_id);
CREATE INDEX round_idx on scores (round);
`

var schemas = []Schema{
	{"players", playersSchema},
	{"scores", scoresSchema},
}

func CreateTable(db *sql.DB) error {
	for _, schema := range schemas {
		if _, err := db.Exec(schema.content); err != nil {
			return err
		}
	}
	return nil
}
