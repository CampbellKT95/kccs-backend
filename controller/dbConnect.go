package controller

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var Db *pgx.Conn
var err error

func DbConnect() {
	//connecting to postgresql
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("could not load .env variable")
		os.Exit(1)
	}

	Db, err = pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// defer Db.Close(context.Background())

	fmt.Println("Connected to sql db")
}
