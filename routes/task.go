package routes

import (
	"database/sql"

	"github.com/denisemignoli/to-do-list/controllers"
	"github.com/denisemignoli/to-do-list/repositories"
	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine, db *sql.DB) {
	// Repositório e Controlador
	taskRepo := repositories.NewTaskMySQLRepository(db)
	taskController := controllers.NewTaskController(taskRepo)

	// Rotas de tarefas
	router.GET("/users/:userid/tasks", taskController.GetTasksByUser)
	router.GET("/users/:userid/tasks/:id", taskController.GetTaskByID)
	router.POST("/users/:userid/tasks", taskController.PostTasks)
	router.PUT("/users/:userid/tasks/:id", taskController.UpdateTask)
	router.DELETE("/users/:userid/tasks/:id", taskController.DeleteTask)
}
