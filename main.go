package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

// using table: tasks
func dbConnect() {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select name from tasks")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var name string
	var description string
	var dueDate time.Time
	var status bool

	if err := rows.Scan(&name, &description, &dueDate, &status); err != nil {
		fmt.Fprintf(os.Stderr, "scan failed")
	}

	fmt.Println(rows)

	fmt.Println("Connected to sql db")
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()

	dbConnect()
}
