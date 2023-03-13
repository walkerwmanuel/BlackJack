package data

//Doing stuff in DB is sql querry

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"walkerwmanuel/blackjack/types"

	_ "github.com/mattn/go-sqlite3"
)

// const dbFilename = "data.db"

// ConnectDatabase - Begins DB operations and creates data.db inside data directory if not there
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

	statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS " + tableName + " (username TEXT, password TEXT, money INT, PRIMARY KEY(username))")
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
	stmt, err := tx.Prepare("INSERT or REPLACE INTO players (username, password, money) VALUES (?, ?, ?)")

	if err != nil {
		fmt.Println("Error!")
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newPlayer.Username, newPlayer.Password, newPlayer.Money)

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

	jsonPlayer, _ := json.Marshal(newGame.Players)

	_, err = stmt.Exec(newGame.Id, jsonPlayer)

	if err != nil {
		fmt.Println("Error2!")
		return false, err
	}

	tx.Commit()

	return true, nil
}

// GetPlayerByUsername - Gets a player from db from username and stores it in a pointer to slice of type player
func GetPlayerByUsername(username string) (*types.Player, error) {

	rows, err := DB.Query("SELECT * FROM players WHERE username like '%" + username + "%'")

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		ourPlayer := types.Player{}
		err = rows.Scan(&ourPlayer.Username, &ourPlayer.Password, &ourPlayer.Money)
		if err != nil {
			log.Fatal(err)
		}
		return &ourPlayer, err
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return nil, err
}

// UpdatePlayer - Updates player at their username
func UpdatePlayer(player *types.Player) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE players SET password = ?, money = ? WHERE Username = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(player.Password, player.Money, player.Username)

	if err != nil {
		return false, err
	}

	tx.Commit()
	return true, nil
}
