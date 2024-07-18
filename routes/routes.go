package routes

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/denisemignoli/to-do-list/controllers"
	"github.com/denisemignoli/to-do-list/repositories"
	"github.com/gin-gonic/gin"
)

const (
	username = "root"
	password = "code2022"
	host     = "localhost"
	port     = 3306
	database = "db_tasks"
)

func SetupRoutes(router *gin.Engine) {
	// Inicializa o banco de dados
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Repositórios
	taskRepo := repositories.NewTaskMySQLRepository(db)
	userRepo := repositories.NewUserMySQLRepository(db)

	// Controladores
	taskController := controllers.NewTaskController(taskRepo)
	userController := controllers.NewUserController(userRepo)

	// Rotas para tarefas associadas ao usuário
	router.GET("/users/:userid/tasks", taskController.GetTasksByUser)
	router.GET("/users/:userid/tasks/:id", taskController.GetTaskByID)
	router.POST("/users/:userid/tasks", taskController.PostTasks)
	router.PUT("/users/:userid/tasks/:id", taskController.UpdateTask)
	router.DELETE("/users/:userid/tasks/:id", taskController.DeleteTask)

	// Rotas para usuários
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
}
