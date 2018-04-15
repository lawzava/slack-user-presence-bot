package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Handling errors remotely
// Should be some slack notifiaction to tell when something fatal came up
func handleError(err error) {
	log.Fatal(err)
}

func main() {
	slackToken := os.Getenv("SLACK_TOKEN")

	// Connecting to SQLite DB, if not exists - creating
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initiating users table if not exists
	dbInitiate := `
	CREATE TABLE IF NOT EXISTS users (id text not null, timestamp bigint not null);`
	_, err = db.Exec(dbInitiate)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			users, err := checkUsersPresence(slackToken)
			if err != nil {
				handleError(err)
			}

			err = writeToDB(db, users)
			if err != nil {
				handleError(err)
			}
			time.Sleep(10 * time.Minute)
		}
	}()

	handler := httpHandler{
		DB:    db,
		Token: slackToken,
	}
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/help", handler.help)
	http.HandleFunc("/total", handler.total)
	log.Fatal(http.ListenAndServe(addr, nil))
}
