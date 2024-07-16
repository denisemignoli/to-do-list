package main

import (
	"github.com/denisemignoli/to-do-list/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run("localhost:8080")
}
