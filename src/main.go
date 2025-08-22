package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "timescaledb"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "test_db"
)

func main() {

	host := os.Getenv("host")
	portStr := os.Getenv("port")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	port, err := strconv.Atoi(portStr)

	if err != nil {
		fmt.Println("Erro ao converter a porta para um número:", err)
		port = 5432
	}

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
