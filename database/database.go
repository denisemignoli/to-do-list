package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	username = "root"
	password = "code2022"
	host     = "localhost"
	port     = 3306
	database = "db_tasks"
)

func InitDB() (*sql.DB, error) {
	// Define o DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)

	// Abre uma conexão com o banco de dados
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging the database:", err)
		return nil, err
	}

	// Cria as tabelas, se necessário
	CreateTables(db)

	return db, nil
}

func CreateTables(db *sql.DB) {
	// Criação da tabela users
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("Erro ao criar tabela users:", err)
	}

	// Criação da tabela tasks
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			completed BOOLEAN NOT NULL DEFAULT 0,
			user_id BIGINT,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		log.Fatal("Erro ao criar tabela tasks:", err)
	}
}
