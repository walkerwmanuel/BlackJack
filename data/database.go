package data

//Doing stuff in DB is sql querry

import (
	"database/sql"
	"fmt"
	"walkerwmanuel/blackjack/types"

	_ "github.com/mattn/go-sqlite3"
)

// const dbFilename = "data.db"

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

var DB *sql.DB

// InsertPlayerToDB - Inserts player to DB
func InsertPlayerToDB(newPlayer *types.Player) (bool, error) {
	// ERROR HERE
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("INSERT INTO players (username, password) VALUES (?, ?)")

	if err != nil {
		fmt.Println("Error!")
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newPlayer.Username, newPlayer.Password)

	if err != nil {
		fmt.Println("Error2!")
		return false, err
	}

	tx.Commit()

	return true, nil
}
