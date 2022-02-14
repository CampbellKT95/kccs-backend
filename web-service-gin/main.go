package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	TITLE  string  `json:"tite"`
	ARTIST string  `json:"artist"`
	PRICE  float64 `json: "price"`
}

func getAlbums(context *gin.Context) {
	//serializes album struct into JSON & adds it to response
	context.IndentedJSON(http.StatusOK, albums)
}

var albums = []album{
	{ID: "1", TITLE: "Blue Train", ARTIST: "John Coltrane", PRICE: 56.99},
	{ID: "2", TITLE: "JERU", ARTIST: "Gerry Mulligan", PRICE: 17.99},
	{ID: "3", TITLE: "Sarah Vaughan and Clifford Brown", ARTIST: "Sarah Vaugh", PRICE: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost: 8080")
}
