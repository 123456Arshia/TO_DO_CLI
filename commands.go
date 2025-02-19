package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Task struct (Exported fields for JSON handling)
type Task struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description,omitempty"`
	Completed     bool    `json:"completed"`
	TimeRemaining float64 `json:"time_remaining"`
}

// readTask reads tasks from a file and returns them
func readTask(fileName string) ([]Task, error) {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("The file does not exist. Creating a new task list.")
			return []Task{}, nil
		}
		return nil, fmt.Errorf("could not open the file: %v", err)
	}
	defer file.Close()

	// Decode JSON
	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		if err == io.EOF {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("could not decode the file: %v", err)
	}

	return tasks, nil
}

// addTask adds a new task to the task list
func addTask(filename string, newTask Task) error {
	// Read existing tasks
	tasks, err := readTask(filename)
	if err != nil {
		return fmt.Errorf("could not read the file: %v", err)
	}

	// Assign a new ID and append the task
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	// Convert tasks back to JSON
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("could not encode task: %v", err)
	}

	// Write updated tasks back to file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write the file: %v", err)
	}

	fmt.Println("Task added successfully!")
	return nil
}

// deleteTask removes a task by ID or title
func deleteTask(filename string, id int, name string) error {
	tasks, err := readTask(filename)
	if err != nil {
		return fmt.Errorf("could not read the file: %v", err)
	}

	var updatedTasks []Task
	var found bool

	// Find and remove the task
	for _, t := range tasks {
		if (id != 0 && t.ID == id) || (name != "" && t.Title == name) {
			found = true
			continue // Skip this task (delete it)
		}
		updatedTasks = append(updatedTasks, t)
	}

	// Check if task was found **after** looping
	if !found {
		return fmt.Errorf("task not found")
	}

	// Convert back to JSON and save
	data, err := json.MarshalIndent(updatedTasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode tasks: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	fmt.Println("Task deleted successfully!")
	return nil
}

// updateTask modifies an existing task based on ID
func updateTask(filename string, id int, updatedTask Task) error {
	tasks, err := readTask(filename)
	if err != nil {
		return fmt.Errorf("failed to read tasks: %v", err)
	}

	var found bool

	// Find and update the task
	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = updatedTask
			tasks[i].ID = id // Keep the original ID
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	// Convert back to JSON and save
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode tasks: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	fmt.Println("Task updated successfully!")
	return nil
}
