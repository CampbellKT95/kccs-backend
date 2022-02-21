package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(context *gin.Context) {
	name := context.DefaultQuery("name")
	description := context.DefaultQuery("description")
	dueDate := context.DefaultQuery("dueDate")
	status := context.DefaultQuery("status")

	context.IndentedJSON(http.StatusOK, name, description, dueDate, status)
}

func CreateTask(context *gin.Context) {

}
