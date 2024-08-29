package routes

import (
	"database/sql"

	"github.com/denisemignoli/to-do-list/controllers"
	"github.com/denisemignoli/to-do-list/repositories"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, db *sql.DB) {
	// Repositório e Controlador
	userRepo := repositories.NewUserMySQLRepository(db)
	userController := controllers.NewUserController(userRepo)

	// Rotas de usuários
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
}
