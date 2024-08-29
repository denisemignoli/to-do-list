package routes

import (
	"database/sql"
	"fmt"
	"log"

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

	// Configurar módulos de rotas
	SetupTaskRoutes(router, db)
	SetupUserRoutes(router, db)
}
