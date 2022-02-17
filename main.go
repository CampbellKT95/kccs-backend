package main

import (
	//personal packages

	//"db"

	// standard library
	"context"
	"fmt"
	"gin"
	"net/http"
	"os"
	"time"

	//additional libraries
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func getTasks(context *gin.Context) {
	name := context.DefaultQuery("name")
	description := context.DefaultQuery("description")
	dueDate := context.DefaultQuery("dueDate")
	status := context.DefaultQuery("status")

	context.IndentedJSON(http.StatusOK, name, description, dueDate, status)
}

func createTask(context *gin.Context) {

}

// https://github.com/gin-gonic/gin

func main() {
	//connect to db
	db.dbConnect()

	//establishing router
	router := gin.Default()

	//fetch all tasks
	router.GET("/tasks", getTasks)

	router.Run()

	// ----------------------------------------------
	// ----------------------------------------------

	//connecting to postgresql
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("could not load .env variable")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// ----------------------------------------------
	// ----------------------------------------------

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

		fmt.Println(name, description, dueDate, status)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "query failed")
	}

	fmt.Println("Connected to sql db")
}
