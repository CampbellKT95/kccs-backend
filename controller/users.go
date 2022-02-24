package controller

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	name        string
	description string
	dueDate     time.Time
	status      bool
}

// fetching all tasks currently open
func GetTasks(conn *gin.Context) {
	rows, err := Db.Query(conn, `select * from tasks`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rows.Close()

	var name string
	var description string
	var dueDate time.Time
	var status bool

	//ranging over rows, one at a time
	for rows.Next() {
		err := rows.Scan(&name, &description, &dueDate, &status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "scan failed")
		}

		fmt.Println(name, description, dueDate, status)
	}

	if rows.Err() != nil {
		log.Fatal("error scanning rows", err)
	}
}

func CreateTask(conn *gin.Context) {

	// stmt := `INSERT INTO tasks (name, description, dueDate, status) VALUES (?, ?, ?, ?)`
	// _, err = Db.Exec(conn, stmt, "cook", "make dinner", "2022-03-10", "false")

	// stmt, err := Db.Exec(conn, `INSERT INTO tasks (name, description, due_date, status) VALUES ('rest', 'like, take a nap?', '2022-02-23', false)`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	stmt, err := Db.Exec(conn, `INSERT INTO tasks (name, description, due_date, status) VALUES ($1, $2, $3, $4)`, `play`, `video games, maybe sekiro?`, `2022-02-24`, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("entry added", stmt)
}
