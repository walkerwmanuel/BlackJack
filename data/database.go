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

// CreateTableGames - Creates a Games table to match Game struct and uses id as primary key
func CreateTableGames() error {
	tableName := "Games"

	statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS " + tableName + " (id INT, players BLOB, PRIMARY KEY(id))")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

// CreateTablePlayers - Creates a Players table to match Player struct and uses username as primary key
func CreateTablePlayers() error {
	tableName := "Players"

	statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS " + tableName + " (username INT, password TEXT, PRIMARY KEY(username))")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

// InsertPlayerToDB - Inserts player to DB
func InsertPlayerToDB(newPlayer *types.Player) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("INSERT or REPLACE INTO players (username, password) VALUES (?, ?)")

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

// InsertGameToDB - Inserts game to DB
func InsertGameToDB(newGame *types.Game) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("INSERT or REPLACE INTO games (id, players) VALUES (?, ?)")

	if err != nil {
		fmt.Println("Error!")
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newGame.Id, newGame.Players)

	if err != nil {
		fmt.Println("Error2!")
		return false, err
	}

	tx.Commit()

	return true, nil
}
