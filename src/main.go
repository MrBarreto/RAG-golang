package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/jonathanhecl/gollama"
	_ "github.com/lib/pq"
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

	ctx := context.Background()

	g := gollama.New("llama3.2") // Create a new Gollama with the default model

	g.Verbose = true

	if err := g.PullIfMissing(ctx); err != nil { // Pull the model if it is not available
		fmt.Println("Error:", err)
		return
	}

	prompt := "what is the capital of Argentina? and who is the founder of peronism?"

	type Capital struct {
		Capital string `json:"capital" required:"true" description:"the capital of a country"`
		Founder string `json:"founder" required:"true" description:"founder of peronism"`
	}

	option := gollama.StructToStructuredFormat(Capital{}) // Convert the struct to a structured format

	fmt.Printf("Option: %+v\n", option)

	output, err := g.Chat(ctx, prompt, option) // Generate a response
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\n%s\n", output.Content) // Print the response
}
