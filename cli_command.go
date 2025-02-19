package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Handles adding a new task
func handleAddTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter task description: ")
	scanner.Scan()
	description := scanner.Text()

	fmt.Print("Enter time remaining (hours): ")
	scanner.Scan()
	timeRemaining, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	newTask := Task{
		Title:         title,
		Description:   description,
		Completed:     false,
		TimeRemaining: timeRemaining,
	}

	if err := addTask(filename, newTask); err != nil {
		fmt.Printf("Error adding task: %v\n", err)
	}
}

// Handles listing tasks
func handleListTasks() {
	tasks, err := readTask(filename)
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("\n===== Task List =====")
	for _, t := range tasks {
		status := "[ ]"
		if t.Completed {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s (%.1f hours left)\n", t.ID, status, t.Title, t.TimeRemaining)
	}
}

// Handles updating a task
func handleUpdateTask(scanner *bufio.Scanner) {
	handleListTasks()

	fmt.Print("Enter task ID to update: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	fmt.Print("Enter new title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter new description: ")
	scanner.Scan()
	description := scanner.Text()

	fmt.Print("Is the task completed? (yes/no): ")
	scanner.Scan()
	completed := strings.ToLower(scanner.Text()) == "yes"

	fmt.Print("Enter new time remaining (hours): ")
	scanner.Scan()
	timeRemaining, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	updatedTask := Task{
		Title:         title,
		Description:   description,
		Completed:     completed,
		TimeRemaining: timeRemaining,
	}

	if err := updateTask(filename, id, updatedTask); err != nil {
		fmt.Printf("Error updating task: %v\n", err)
	}
}

// Handles deleting a task
func handleDeleteTask(scanner *bufio.Scanner) {
	handleListTasks()

	fmt.Print("Enter task ID to delete: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	if err := deleteTask(filename, id, ""); err != nil {
		fmt.Printf("Error deleting task: %v\n", err)
	}
}
