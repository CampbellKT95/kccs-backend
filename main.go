package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

// using table: tasks
func dbConnect() {

	conn, err := pgx.Connect(context.Background(), os.Getenv(DB_URL))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var description string
	// what data type is a date?
	var dueDate Time
	var status boolean

	err = conn.QueryRow(context.Background(), "select * from tasks", 42).Scan(&name, &description, &dueDate, &status)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(name, description, dueDate, status)
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
