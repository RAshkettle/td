package task

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	PersistData = false
	Add("Test Task")
	assert.Equal(t, 1, len(Tasks))
	foundTask := false
	for _, task := range Tasks {
		if task.FriendlyId == 1 {
			assert.Equal(t, "Test Task", task.Description)
			foundTask = true
		}
	}
	assert.True(t, foundTask)
	Tasks = make(map[uuid.UUID]Task)
	currentFriendlyIdMax = 0
}

func TestGetTaskByFriendlyId(t *testing.T) {
	PersistData = false
	Add("Test Task")
	Add("Test Task 2")
	Add("Test Task 3")

	task := getTaskByFriendlyId(2)
	assert.Equal(t, "Test Task 2", task.Description)
	Tasks = make(map[uuid.UUID]Task)
	currentFriendlyIdMax = 0
}

func TestCloseTask(t *testing.T) {
	PersistData = false
	Add("Test Task")
	Add("Test Task 2")
	Add("Test Task 3")
	Close(1)
	task := getTaskByFriendlyId(1)
	assert.Equal(t, Completed, task.Status)
	task = getTaskByFriendlyId(2)
	assert.Equal(t, Open, task.Status)
	Tasks = make(map[uuid.UUID]Task)
	currentFriendlyIdMax = 0
}

func TestReopenTask(t *testing.T) {
	PersistData = false
	Add("Test Task")
	Add("Test Task 2")
	Add("Test Task 3")
	Close(1)
	Close(2)
	ReOpen(1)
	task := getTaskByFriendlyId(1)
	assert.Equal(t, Open, task.Status)
	task = getTaskByFriendlyId(2)
	assert.Equal(t, Completed, task.Status)
	Tasks = make(map[uuid.UUID]Task)
	currentFriendlyIdMax = 0
}

func TestDelete(t *testing.T) {
	PersistData = false
	Add("Test Task")
	Add("Test Task 2")
	Add("Test Task 3")
	Delete(1)
	assert.Equal(t, 2, len(Tasks))
	for _, task := range Tasks {
		assert.NotEqual(t, 1, task.FriendlyId)
	}
	Tasks = make(map[uuid.UUID]Task)
	currentFriendlyIdMax = 0
}
