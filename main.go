package main

import (
	// standard library
	"fmt"

	//personal packages
	"github.com/CampbellKT95/kccs-backend/controller"

	//3rd party libraries
	"github.com/gin-gonic/gin"
)

func main() {
	//connect to postgresql
	controller.DbConnect()

	//establishing router
	router := gin.Default()

	router.GET("/tasks", controller.GetTasks)
	router.POST("/tasks", controller.CreateTask)
	router.PUT("/tasks/:id", controller.UpdateTask)
	// router.DELETE("/tasks/:id", controller.DeleteTask)

	fmt.Println("server running")

	router.Run()
}
