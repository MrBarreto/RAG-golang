package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)
const (
	host = "timescaledb"
	port = 5432
	user = "postgres"
	password = "password"
	dbname = "test_db"
)

func main() {
	fmt.Println("Teste")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
    		panic(err)
  	}
  	defer db.Close()

  	err = db.Ping()
  	if err != nil {
    		panic(err)
  	}

  	fmt.Println("Successfully connected!")
}
