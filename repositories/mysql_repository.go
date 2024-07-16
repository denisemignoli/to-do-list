package repositories

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/denisemignoli/to-do-list/models"
	_ "github.com/go-sql-driver/mysql"
)

type TaskMySQLRepository struct {
	db *sql.DB
}

const (
	username = "root"
	password = "code2022"
	host     = "localhost"
	port     = 3306
	database = "db_tasks"
)

func NewTaskMySQLRepository() *TaskMySQLRepository {
	// build the DNS
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
	// open the connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("impossible to create the connection: %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// execute the script to create the table
	err = executeSQLScript(db, "scripts/create_table.sql")
	if err != nil {
		log.Fatalf("failed to execute the script: %s", err)
	}

	return &TaskMySQLRepository{
		db: db,
	}
}

func executeSQLScript(db *sql.DB, filename string) error {
	// read the file
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	// execute the script
	_, err = db.Exec(string(b))
	return err
}

func (t *TaskMySQLRepository) GetTasks() []models.Task {
	var tasks []models.Task

	rows, err := t.db.Query("SELECT * FROM `tasks`")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Completed)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func (t *TaskMySQLRepository) SaveTask(newTask models.Task) (int64, error) {
	result, err := t.db.Exec("INSERT INTO `tasks` (`name`, `completed`) VALUES (?, ?)", newTask.Name, newTask.Completed)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t *TaskMySQLRepository) UpdateTask(updatedTask models.Task) (*models.Task, error) {
	_, err := t.db.Exec("UPDATE `tasks` SET `name` = ?, `completed` = ? WHERE `id` = ?", updatedTask.Name, updatedTask.Completed, updatedTask.ID)
	if err != nil {
		return nil, err
	}
	return &updatedTask, nil
}

func (t *TaskMySQLRepository) GetTaskByID(id int64) (*models.Task, error) {
	var task models.Task

	err := t.db.QueryRow("SELECT `id`, `name`, `completed` FROM `tasks` WHERE `id` = ?", id).Scan(&task.ID, &task.Name, &task.Completed)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *TaskMySQLRepository) DeleteTask(id int64) error {
	_, err := t.db.Exec("DELETE FROM `tasks` WHERE `id` = ?", id)
	if err != nil {
		return err
	}

	return nil
}
