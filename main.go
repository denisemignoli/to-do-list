package main

import (
	"log"

	"github.com/denisemignoli/to-do-list/database"
	"github.com/denisemignoli/to-do-list/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o banco de dados
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Inicializa o router do Gin
	router := gin.Default()

	// Configura as rotas, passando o banco de dados para elas
	routes.SetupTaskRoutes(router, db)
	routes.SetupUserRoutes(router, db)

	// Inicia o servidor
	router.Run("localhost:8080")
}
