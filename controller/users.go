package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTasks(context *gin.Context) {
	rows, err := Db.Query(context, `select * from tasks`)
	if err != nil {
		fmt.Println(err)
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
}

func CreateTask(context *gin.Context) {

}
