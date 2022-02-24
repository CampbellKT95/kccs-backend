package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	var parsedData Task

	// takes in the json from post request
	jsonData, err := ioutil.ReadAll(conn.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	//below shows we are successfully retreiving the data from post request. how do we parse it correctly?
	jsonbody := string(jsonData)

	fmt.Println(jsonbody)

	decoder := json.NewDecoder(strings.NewReader(string(jsonData)))

	_ = decoder.Decode(&parsedData)

	fmt.Printf("received data: %s:%s/n", parsedData.name, parsedData.description)

	//BELOW STATEMENT WORKS, BUT HOW CAN I GET THE BODY DATA TO INSERT INTO THE PREPARED STATEMENT?

	// stmt, err := Db.Exec(conn, `INSERT INTO tasks (name, description, due_date, status) VALUES ($1, $2, $3, $4)`, `play`, `video games, maybe sekiro?`, `2022-02-24`, false)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("entry added", stmt)
}
