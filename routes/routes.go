package routes

import (
	"github.com/denisemignoli/to-do-list/controllers"
	"github.com/denisemignoli/to-do-list/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	taskController := controllers.NewTaskController(repositories.OpenConnection())

	router.GET("/tasks", taskController.GetTasks)
	router.GET("/tasks/:id", taskController.GetTaskByID)
	router.POST("/tasks", taskController.PostTasks)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)
}
