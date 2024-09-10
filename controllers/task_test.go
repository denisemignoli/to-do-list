package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisemignoli/to-do-list/controllers"
	"github.com/denisemignoli/to-do-list/mocks"
	"github.com/denisemignoli/to-do-list/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTasksByUser(t *testing.T) {
	// arrange
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockRepo := new(mocks.TaskRepository)
	userID := int64(1)
	mockRepo.On("GetTasksByUserID", userID).Return([]models.Task{
		{ID: 1, Name: "Test Task 1", Completed: false, UserID: userID},
		{ID: 2, Name: "Test Task 2", Completed: true, UserID: userID},
	})

	taskController := controllers.NewTaskController(mockRepo)
	router.GET("/users/:userid/tasks", taskController.GetTasksByUser)

	// act
	req, _ := http.NewRequest(http.MethodGet, "/users/1/tasks", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `[{"id":1,"name":"Test Task 1","completed":false,"user_id":1},{"id":2,"name":"Test Task 2","completed":true,"user_id":1}]`, rr.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestPostTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockRepo := new(mocks.TaskRepository)
	userID := int64(1)
	newTask := models.Task{Name: "New Task", Completed: false, UserID: userID}
	mockRepo.On("SaveTask", newTask).Return(int64(1), nil)

	taskController := controllers.NewTaskController(mockRepo)
	router.POST("/users/:userid/tasks", taskController.PostTasks)

	taskJSON, _ := json.Marshal(newTask)
	req, _ := http.NewRequest(http.MethodPost, "/users/1/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	newTask.ID = 1
	expectedResponse, _ := json.Marshal(newTask)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.JSONEq(t, string(expectedResponse), rr.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockRepo := new(mocks.TaskRepository)
	updatedTask := models.Task{ID: 1, Name: "Updated Task", Completed: true}
	mockRepo.On("UpdateTask", updatedTask).Return(&updatedTask, nil)

	taskController := controllers.NewTaskController(mockRepo)
	router.PUT("/tasks/:id", taskController.UpdateTask)

	taskJSON, _ := json.Marshal(updatedTask)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	expectedResponse, _ := json.Marshal(updatedTask)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, string(expectedResponse), rr.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockRepo := new(mocks.TaskRepository)
	task := models.Task{ID: 1, Name: "Test Task", Completed: false}
	mockRepo.On("GetTaskByID", int64(1)).Return(&task, nil)

	taskController := controllers.NewTaskController(mockRepo)
	router.GET("/tasks/:id", taskController.GetTaskByID)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	expectedResponse, _ := json.Marshal(task)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, string(expectedResponse), rr.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockRepo := new(mocks.TaskRepository)
	mockRepo.On("DeleteTask", int64(1)).Return(nil)

	taskController := controllers.NewTaskController(mockRepo)
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"message":"Task deleted"}`, rr.Body.String())
	mockRepo.AssertExpectations(t)
}
