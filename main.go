package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", getRoot)
	router.GET("/health", getHealth)
	router.GET("/is-alive", getAliveness)
	router.GET("/api/v1/albums", getAlbums)

	err := router.Run(":80")
	if err != nil {
		log.Fatal(err)
	}
}

func getRoot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rootObject)
}

type root struct {
	Links rootLinks `json:"_Links"`
}

type rootLinks struct {
	Health string `json:"Health"`
	Alive  string `json:"Aliveness"`
	Albums string `json:"Albums"`
}

var rootObject = root{Links: rootLinks{
	Health: "https://api.callofthevoid.dk/health",
	Alive:  "https://api.callofthevoid.dk/is-alive",
	Albums: "https://api.callofthevoid.dk/api/v1/albums",
}}

func getAliveness(c *gin.Context) {
	c.String(http.StatusOK, "alive")
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
