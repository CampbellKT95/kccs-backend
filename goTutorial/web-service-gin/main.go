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

var albums = []album{
	{ID: "1", TITLE: "Blue Train", ARTIST: "John Coltrane", PRICE: 56.99},
	{ID: "2", TITLE: "JERU", ARTIST: "Gerry Mulligan", PRICE: 17.99},
	{ID: "3", TITLE: "Sarah Vaughan and Clifford Brown", ARTIST: "Sarah Vaugh", PRICE: 39.99},
}

func getAlbums(context *gin.Context) {
	//serializes album struct into JSON & adds it to response
	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album

	//call BindJSON to bind the received JSON to new album
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	//add the new album to the slice
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(context *gin.Context) {
	//equates to var id = context.Param(id)
	id := context.Param("id")

	for _, a := range albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost: 8080")
}
