package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id          string `json:"id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Status      bool   `json: "status"`
}

// ------------------------------------------------------------------------
// fetching all tasks currently open
func GetTasks(conn *gin.Context) {
	rows, err := Db.Query(conn, `select * from tasks`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rows.Close()

	// var Id string
	// var Name string
	// var Description string
	// var Status bool

	var tasks []Task

	//ranging over rows, one at a time
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Status)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
		fmt.Println(tasks)
	}

	//encodes the tasks into json
	response, err := json.Marshal(tasks)

	conn.Data(http.StatusOK, "application/json", response)

	if rows.Err() != nil {
		log.Fatal(err)
	}
}

// ------------------------------------------------------------------------
func CreateTask(conn *gin.Context) {
	// takes in the json from post request
	jsonData, err := ioutil.ReadAll(conn.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	var parsedData Task

	//will need to loop through the jsonData & parse seperately since time needs a specific method for parsing
	err = json.Unmarshal(jsonData, &parsedData)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsedData.Name)

	//can supposedly reuse body data with the below code?
	// conn.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))

	//returns data to clientside
	conn.Data(http.StatusOK, "application/json", jsonData)

	// --------------------

	stmt, err := Db.Exec(conn, `INSERT INTO tasks (id, name, description, status) VALUES ($1, $2, $3, $4)`, parsedData.Id, parsedData.Name, parsedData.Description, parsedData.Status)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("entry added", stmt)
}

// ------------------------------------------------------------------------
// func UpdateTask() {}

// ------------------------------------------------------------------------
// func DeleteTask() {}
