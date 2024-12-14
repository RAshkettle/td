package task

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MemoryPersistor struct{}

func NewMemoryPersistor() *MemoryPersistor {
	return &MemoryPersistor{}
}

func (m *MemoryPersistor) saveTasks() {
	// Do nothing
}

func (m *MemoryPersistor) getTasks() {
	// Do nothing
}

func TestAdd(t *testing.T) {
	mp := NewMemoryPersistor()
	Add("Test Task", mp)
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
	mp := NewMemoryPersistor()

	Add("Test Task", mp)
	Add("Test Task 2", mp)
	Add("Test Task 3", mp)

	task := getTaskByFriendlyId(2)
	assert.Equal(t, "Test Task 2", task.Description)
	Tasks = make(map[uuid.UUID]Task)
	currentFriendlyIdMax = 0
}

func TestCloseTask(t *testing.T) {
	mp := NewMemoryPersistor()

	Add("Test Task", mp)
	Add("Test Task 2", mp)
	Add("Test Task 3", mp)
	Close(1, mp)
	task := getTaskByFriendlyId(1)
	assert.Equal(t, Completed, task.Status)
	task = getTaskByFriendlyId(2)
	assert.Equal(t, Open, task.Status)
	Tasks = make(map[uuid.UUID]Task)
	currentFriendlyIdMax = 0
}

func TestReopenTask(t *testing.T) {
	mp := NewMemoryPersistor()

	Add("Test Task", mp)
	Add("Test Task 2", mp)
	Add("Test Task 3", mp)
	Close(1, mp)
	Close(2, mp)
	ReOpen(1, mp)
	task := getTaskByFriendlyId(1)
	assert.Equal(t, Open, task.Status)
	task = getTaskByFriendlyId(2)
	assert.Equal(t, Completed, task.Status)
	Tasks = make(map[uuid.UUID]Task)
	currentFriendlyIdMax = 0
}

func TestDelete(t *testing.T) {
	mp := NewMemoryPersistor()

	Add("Test Task", mp)
	Add("Test Task 2", mp)
	Add("Test Task 3", mp)
	Delete(1, mp)
	assert.Equal(t, 2, len(Tasks))
	for _, task := range Tasks {
		assert.NotEqual(t, 1, task.FriendlyId)
	}
	Tasks = make(map[uuid.UUID]Task)
	currentFriendlyIdMax = 0
}
