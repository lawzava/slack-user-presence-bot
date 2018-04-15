package main

import (
	"database/sql"
	"fmt"
	"time"
)

// Handling writing users data to a table
func writeToDB(db *sql.DB, users []userRawData) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error while initiating transactions: %v", err)
	}
	stmt, err := tx.Prepare("insert into users(id, timestamp) values(?, ?)")
	if err != nil {
		return fmt.Errorf("error while preparing transactions: %v", err)
	}
	defer stmt.Close()

	timestamp := time.Now().Unix()
	for _, user := range users {
		if user.Presence == "active" {
			_, err = stmt.Exec(user.ID, timestamp)
			if err != nil {
				return fmt.Errorf("error while executing transaction: %v", err)
			}
		}
	}
	return tx.Commit()
}

type userHistoricalData struct {
	ID        string
	Timestamp int64
}

// Handling historical users data retrieval from DB
func readUsersDataFromDB(db *sql.DB, startDate, endDate int64) ([]userHistoricalData, error) {
	rows, err := db.Query("SELECT id, timestamp from users where timestamp >= $1 AND timestamp <= $2",
		startDate,
		endDate,
	)
	if err != nil {
		return nil, fmt.Errorf("error while querying database: %v", err)
	}
	defer rows.Close()

	var users []userHistoricalData
	for rows.Next() {
		var user userHistoricalData
		err = rows.Scan(&user.ID, &user.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("error while scanning row: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}
