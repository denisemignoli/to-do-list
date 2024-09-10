package repositories

import (
	"database/sql"
	"log"

	"github.com/denisemignoli/to-do-list/models"
	_ "github.com/go-sql-driver/mysql"
)

type TaskMySQLRepository struct {
	db *sql.DB
}

func NewTaskMySQLRepository(db *sql.DB) *TaskMySQLRepository {
	return &TaskMySQLRepository{
		db: db,
	}
}

func (t *TaskMySQLRepository) GetTasksByUserID(userID int64) ([]models.Task, error) {
	var tasks []models.Task

	rows, err := t.db.Query("SELECT * FROM `tasks` WHERE `user_id` = ?", userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Completed, &task.UserID); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
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
		err := rows.Scan(&task.ID, &task.Name, &task.Completed, &task.UserID)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func (t *TaskMySQLRepository) SaveTask(newTask models.Task) (int64, error) {
	result, err := t.db.Exec("INSERT INTO `tasks` (`name`, `completed`, `user_id`) VALUES (?, ?, ?)", newTask.Name, newTask.Completed, newTask.UserID)
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
	_, err := t.db.Exec("UPDATE `tasks` SET `name` = ?, `completed` = ?, `user_id` = ? WHERE `id` = ?", updatedTask.Name, updatedTask.Completed, updatedTask.UserID, updatedTask.ID)
	if err != nil {
		return nil, err
	}
	return &updatedTask, nil
}

func (t *TaskMySQLRepository) GetTaskByID(id int64) (*models.Task, error) {
	var task models.Task

	err := t.db.QueryRow("SELECT `id`, `name`, `completed`, `user_id` FROM `tasks` WHERE `id` = ?", id).Scan(&task.ID, &task.Name, &task.Completed, &task.UserID)
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
