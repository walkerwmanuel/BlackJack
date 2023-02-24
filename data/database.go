package data

import (
	"database/sql"
	"fmt"
	"log"
	"walkerwmanuel/blackjack/types"
)

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

var DB *sql.DB

// InsertPlayerToDB - Creates table players and inserts player to DB
func InsertPlayerToDB(newPlayer *types.Player) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	_, err = DB.Exec("CREATE TABLE `players` (`username` INTEGER PRIMARY KEY AUTOINCREMENT, `password` VARCHAR(64) NULL)")
	if err != nil {
		log.Fatal(err)
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
