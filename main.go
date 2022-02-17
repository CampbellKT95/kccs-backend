package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	// router := gin.Default()
	// router.GET("/ping", func(context *gin.Context) {
	// 	context.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// router.Run()

	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("could not load .env variable")
		os.Exit(1)
	} else {
		fmt.Printf(".env loaded")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select * from tasks")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var name string
	var description string
	var dueDate time.Time
	var status bool

	for rows.Next() {

		err = rows.Scan(&name, &description, &dueDate, &status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "scan failed")
		}

		fmt.Println(name)
	}
	fmt.Println(name)

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "query failed", rows.Err())
	}

	fmt.Println("Connected to sql db")
}
