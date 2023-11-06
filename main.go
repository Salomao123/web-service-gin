package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json: "id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}

type exception struct {
	StatusCode int    `json: "statusCode"`
	Message    string `json: "message"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaugnhan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbum)
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {

	id := c.Param("id")

	for _, album := range albums {
		if id == album.ID {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, exception{Message: "album não encontrado", StatusCode: http.StatusNotFound})

}

func postAlbum(c *gin.Context) {

	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	for _, currentAlbum := range albums {
		if currentAlbum.ID == newAlbum.ID {
			var errorResponse exception
			errorResponse.Message = "Id já cadastrado na base"
			errorResponse.StatusCode = http.StatusBadRequest
			c.IndentedJSON(http.StatusBadRequest, errorResponse)
			return
		}
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}
