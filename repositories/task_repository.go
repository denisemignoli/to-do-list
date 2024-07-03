package repositories

import "github.com/denisemignoli/to-do-list/models"

type TaskRepository interface {
	GetTasks() []models.Task
	SaveTask(newTask models.Task) (int64, error)
	UpdateTask(updatedTask models.Task) (*models.Task, error)
	GetTaskByID(id int64) (*models.Task, error)
	DeleteTask(id int64) error
}
