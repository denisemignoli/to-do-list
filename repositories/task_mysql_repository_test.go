package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/denisemignoli/to-do-list/models"
	"github.com/denisemignoli/to-do-list/repositories"
	"github.com/stretchr/testify/assert"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *repositories.TaskMySQLRepository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a database connection", err)
	}

	repo := repositories.NewTaskMySQLRepository(db)
	return db, mock, repo
}

func TestGetTasks(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "completed"}).
		AddRow(1, "Task 1", false).
		AddRow(2, "Task 2", true)

	mock.ExpectQuery("SELECT \\* FROM `tasks`").WillReturnRows(rows)

	tasks := repo.GetTasks()

	assert.Len(t, tasks, 2)
	assert.Equal(t, int64(1), tasks[0].ID)
	assert.Equal(t, "Task 1", tasks[0].Name)
	assert.False(t, tasks[0].Completed)
	assert.Equal(t, int64(2), tasks[1].ID)
	assert.Equal(t, "Task 2", tasks[1].Name)
	assert.True(t, tasks[1].Completed)
}

func TestSaveTask(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	newTask := models.Task{Name: "New Task", Completed: false}

	mock.ExpectExec("INSERT INTO `tasks`").
		WithArgs(newTask.Name, newTask.Completed).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := repo.SaveTask(newTask)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
}

func TestUpdateTask(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	updatedTask := models.Task{ID: 1, Name: "Updated Task", Completed: true}

	mock.ExpectExec("UPDATE `tasks` SET `name` = \\?, `completed` = \\? WHERE `id` = \\?").
		WithArgs(updatedTask.Name, updatedTask.Completed, updatedTask.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	task, err := repo.UpdateTask(updatedTask)

	assert.NoError(t, err)
	assert.Equal(t, updatedTask.ID, task.ID)
	assert.Equal(t, updatedTask.Name, task.Name)
	assert.Equal(t, updatedTask.Completed, task.Completed)
}

func TestGetTaskByID(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "name", "completed"}).
		AddRow(1, "Task 1", false)

	mock.ExpectQuery("SELECT `id`, `name`, `completed` FROM `tasks` WHERE `id` = ?").
		WithArgs(1).
		WillReturnRows(row)

	task, err := repo.GetTaskByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, int64(1), task.ID)
	assert.Equal(t, "Task 1", task.Name)
	assert.False(t, task.Completed)
}

func TestDeleteTask(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec("DELETE FROM `tasks` WHERE `id` = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteTask(1)

	assert.NoError(t, err)
}
