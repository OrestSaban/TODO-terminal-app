package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var todos Todos

	// Load tasks from file
	err := todos.loadFromFile("todos.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("No command provided. Use 'todo help' for a list of commands.")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add \"task description\"")
			return
		}
		task := strings.Join(os.Args[2:], " ")
		todos.add(task)
		fmt.Println("Task added.")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete index")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid index. Please provide a valid number.")
			return
		}
		todos.delete(index)
		fmt.Println("Task deleted.")

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todo update index \"new task description\"")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid index. Please provide a valid number.")
			return
		}
		newTask := strings.Join(os.Args[3:], " ")
		todos.update(index, newTask)
		fmt.Println("Task updated.")

	case "list":
		todos.print()

	case "clear":
		todos.clear()
		os.Remove("todos.json")
		fmt.Println("All tasks cleared and file deleted.")

	case "toggle":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo toggle index")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid index. Please provide a valid number.")
			return
		}
		todos.toggle(index)

	case "help":
		fmt.Println("Available commands:")
		fmt.Println("  todo add \"task description\"  - Add a new task")
		fmt.Println("  todo delete index             - Delete a task by index")
		fmt.Println("  todo update index \"new task\" - Update a task by index")
		fmt.Println("  todo list                     - List all tasks")
		fmt.Println("  todo clear                    - Clear all tasks and delete the file")
		fmt.Println("  todo toggle index             - Toggle the completion status of a task")
		fmt.Println("  todo help                     - Show this help message")

	default:
		fmt.Println("Unknown command. Use 'todo help' for a list of commands.")
	}

	// Save tasks to file before exiting
	err = todos.saveToFile("todos.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}