package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename = "Task.json"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display menu
		fmt.Println("\n===== To-Do List Manager =====")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Update Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			handleAddTask(scanner)
		case "2":
			handleListTasks()
		case "3":
			handleUpdateTask(scanner)
		case "4":
			handleDeleteTask(scanner)
		case "5":
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
