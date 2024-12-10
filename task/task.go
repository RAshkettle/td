package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

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

func Add(decription string) {
	getSavedTasks()
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
	saveTasks()
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
func Delete(id int) {
	getSavedTasks()
	t := getTaskByFriendlyId(id)
	delete(Tasks, t.ID)
	saveTasks()
}

func Close(id int) {
	getSavedTasks()
	t := getTaskByFriendlyId(id)
	t.Status = Completed
	t.CompletedAt = time.Now()
	Tasks[t.ID] = t
	saveTasks()
}

func ReOpen(id int) {
	getSavedTasks()
	t := getTaskByFriendlyId(id)
	t.Status = Open
	t.CompletedAt = time.Time{}
	Tasks[t.ID] = t
	saveTasks()
}

func List() {
	getSavedTasks()
}

func getSavedTasks() {
	if !checkFileExists(getFileLocation()) {
		// The file doesn't exist, return an empty map
		Tasks = make(map[uuid.UUID]Task)
		return
	}

	// If we got here, the file exists, so read the data from the file
	data, err := os.ReadFile(getFileLocation())
	if err != nil {
		panic(err)
	}

	// Unmarshal the data into the Tasks map
	Tasks = make(map[uuid.UUID]Task)
	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		panic(err)
	}
}

func saveTasks() {
	jsonData, err := json.MarshalIndent(Tasks, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(getFileLocation())
	fmt.Println(string(jsonData))

	// Save the tasks to a file
	err = os.WriteFile(getFileLocation(), jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func getFileLocation() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Create a path to a file or directory in the home directory
	configPath := filepath.Join(homeDir, ".tasks.json")
	return configPath
}
