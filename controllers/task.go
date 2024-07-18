package controllers

import (
	"net/http"
	"strconv"

	"github.com/denisemignoli/to-do-list/models"
	"github.com/denisemignoli/to-do-list/repositories"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskRepository repositories.TaskRepository
}

func NewTaskController(repo repositories.TaskRepository) *TaskController {
	return &TaskController{
		TaskRepository: repo,
	}
}

var tasks []models.Task
var nextID int = 1

func (tc *TaskController) GetTasksByUser(c *gin.Context) {
	userIDStr := c.Param("userid")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tasks := tc.TaskRepository.GetTasksByUserID(userID)

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) PostTasks(c *gin.Context) {
	userIDStr := c.Param("userid")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.UserID = userID

	id, err := tc.TaskRepository.SaveTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save task"})
		return
	}

	newTask.ID = id

	c.JSON(http.StatusCreated, newTask)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask.ID = id
	task, err := tc.TaskRepository.UpdateTask(updatedTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := tc.TaskRepository.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task"})
		return
	}

	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = tc.TaskRepository.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
