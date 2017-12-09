package main

import (
	"database/sql"
	"net/http"
	"smaug/phoenix-league/go-server/handler"

	_ "github.com/lib/pq"
)

const connStr = "postgres://postgres:phoenix@localhost/postgres?sslmode=disable"

func connectToDB(connStr string) (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", connStr)
	return db, err
}

func main() {
	db, err := connectToDB(connStr)
	if err != nil {
		panic(err)
		return
	}

	http.Handle("/signin", handler.NewSignin(db))
	http.Handle("/api/socket", handler.NewWebSocket(db))
	http.ListenAndServe(":8081", nil)
}
