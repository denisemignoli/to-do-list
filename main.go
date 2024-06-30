package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

var tasks = []task{
	{ID: "1", Name: "Fazer compras", Status: false},
	{ID: "2", Name: "Agendar médico", Status: true},
	{ID: "3", Name: "Cortar cabelo", Status: false},
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.Run("localhost:8080")
}

// getTasks responds with the list of all tasks as JSON.
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}
