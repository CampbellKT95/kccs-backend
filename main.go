package main

import (
	// standard library

	//personal packages
	"net/http"

	"github.com/CampbellKT95/kccs-backend/bots"
	"github.com/CampbellKT95/kccs-backend/controller"

	//3rd party libraries
	"github.com/gin-gonic/gin"
)

func main() {
	//connect to postgresql
	controller.DbConnect()

	//establishing router
	router := gin.Default()

	// to-do tasks
	router.GET("/tasks", controller.GetTasks)
	router.POST("/tasks", controller.CreateTask)
	router.PUT("/tasks/:id", controller.UpdateTask)
	// router.DELETE("/tasks/:id", controller.DeleteTask)

	// bots
	http.HandleFunc("/search/news", bots.NewsScrap)
	http.HandleFunc("/search/stocks", bots.StockScrap)
	bots.RetrieveTweets()

	router.Run()
}
