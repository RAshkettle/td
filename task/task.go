package task

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type TaskPersistor interface {
	saveTasks()
	getTasks()
}

type FileSystemPersistor struct{}

func NewFileSystemPersistor() *FileSystemPersistor {
	return &FileSystemPersistor{}
}

func (f *FileSystemPersistor) saveTasks() {
	jsonData, err := json.MarshalIndent(Tasks, "", "    ")
	if err != nil {
		panic(err)
	}
	// Save the tasks to a file
	err = os.WriteFile(f.getFileLocation(), jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func (f *FileSystemPersistor) getTasks() {
	if !f.checkFileExists(f.getFileLocation()) {
		// The file doesn't exist, return an empty map
		Tasks = make(map[uuid.UUID]Task)
		return
	}

	// If we got here, the file exists, so read the data from the file
	data, err := os.ReadFile(f.getFileLocation())
	if err != nil {
		panic(err)
	}

	// Unmarshal the data into the Tasks map
	Tasks = make(map[uuid.UUID]Task)
	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		panic(err)
	}

	// Find the highest friendly id
	for _, task := range Tasks {
		if task.FriendlyId > currentFriendlyIdMax {
			currentFriendlyIdMax = task.FriendlyId
		}
	}
}

func (f *FileSystemPersistor) checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func (f *FileSystemPersistor) getFileLocation() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Create a path to a file or directory in the home directory
	configPath := filepath.Join(homeDir, ".tasks.json")
	return configPath
}

type Status int

func (s Status) String() string {
	return [...]string{"Open", "Complete"}[s]
}

const (
	Open Status = iota
	Completed
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	FriendlyId  int       `json:"friendly_id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

var (
	Tasks                map[uuid.UUID]Task = make(map[uuid.UUID]Task)
	taskFileExists       bool               = false
	currentFriendlyIdMax int                = 0
)

func Add(decription string, tp TaskPersistor) {
	tp.getTasks()
	id, _ := uuid.NewRandom()
	currentFriendlyIdMax++
	friendlyId := currentFriendlyIdMax

	t := Task{
		ID:          id,
		FriendlyId:  friendlyId,
		Description: decription,
		Status:      Open,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	Tasks[id] = t
	tp.saveTasks()
}

func getTaskByFriendlyId(id int) Task {
	for _, task := range Tasks {
		if task.FriendlyId == id {
			return task
		}
	}
	return Task{}
}

// Delete a task using its friendly id
func Delete(id int, tp TaskPersistor) {
	tp.getTasks()
	t := getTaskByFriendlyId(id)
	delete(Tasks, t.ID)
	tp.saveTasks()
}

func Close(id int, tp TaskPersistor) {
	tp.getTasks()
	t := getTaskByFriendlyId(id)
	t.Status = Completed
	t.CompletedAt = time.Now()
	Tasks[t.ID] = t
	tp.saveTasks()
}

func ReOpen(id int, tp TaskPersistor) {
	tp.getTasks()
	t := getTaskByFriendlyId(id)
	t.Status = Open
	t.CompletedAt = time.Time{}
	Tasks[t.ID] = t
	tp.saveTasks()
}

func List(tp TaskPersistor) {
	tp.getTasks()
}
