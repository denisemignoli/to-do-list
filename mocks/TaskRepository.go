// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	models "github.com/denisemignoli/to-do-list/models"
	mock "github.com/stretchr/testify/mock"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// DeleteTask provides a mock function with given fields: id
func (_m *TaskRepository) DeleteTask(id int64) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTaskByID provides a mock function with given fields: id
func (_m *TaskRepository) GetTaskByID(id int64) (*models.Task, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 *models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*models.Task, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *models.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasks provides a mock function with given fields:
func (_m *TaskRepository) GetTasks() []models.Task {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTasks")
	}

	var r0 []models.Task
	if rf, ok := ret.Get(0).(func() []models.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Task)
		}
	}

	return r0
}

// SaveTask provides a mock function with given fields: newTask
func (_m *TaskRepository) SaveTask(newTask models.Task) (int64, error) {
	ret := _m.Called(newTask)

	if len(ret) == 0 {
		panic("no return value specified for SaveTask")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Task) (int64, error)); ok {
		return rf(newTask)
	}
	if rf, ok := ret.Get(0).(func(models.Task) int64); ok {
		r0 = rf(newTask)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(models.Task) error); ok {
		r1 = rf(newTask)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: updatedTask
func (_m *TaskRepository) UpdateTask(updatedTask models.Task) (*models.Task, error) {
	ret := _m.Called(updatedTask)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 *models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Task) (*models.Task, error)); ok {
		return rf(updatedTask)
	}
	if rf, ok := ret.Get(0).(func(models.Task) *models.Task); ok {
		r0 = rf(updatedTask)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Task) error); ok {
		r1 = rf(updatedTask)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
