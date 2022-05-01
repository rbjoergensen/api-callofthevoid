package api-callofthevoid

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", getHealth)
	router.GET("/albums", getAlbums)

	err := router.Run("localhost:80")
	if err != nil {
		log.Fatal(err)
	}
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, status)
}

type health struct {
	Status string `json:"status"`
}

var status = health{Status: "healthy"}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
