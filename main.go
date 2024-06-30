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
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", postTasks)

	router.Run("localhost:8080")
}

// getTasks responds with the list of all tasks as JSON.
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

// postTasks adds a task from JSON received in the request body.
func postTasks(c *gin.Context) {
	var newTask task

	// Call BindJSON to bind the received JSON to newTask.
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	// Add the new tasks to the slice.
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

// getTaskByID locates the task whose ID value matches the id parameter sent by the client, then returns that task as a response.
func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of tasks, looking for a task whose ID value matches the parameter.
	for _, t := range tasks {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
