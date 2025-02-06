package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const todoFile = "todos.json"

func LoadTodos() ([]Todo, error) {
	file, err := os.ReadFile(todoFile)
	if err != nil {
		return []Todo{}, err
	}

	var todos []Todo
	json.Unmarshal(file, &todos)
	return todos, nil
}

func SaveTodos(todos []Todo) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	return os.WriteFile(todoFile, data, 0755)
}

func addTodo(title string) {
	todos, _ := LoadTodos()
	id := len(todos) + 1
	todos = append(todos, Todo{Id: id, Title: title, Completed: false})
	SaveTodos(todos)
	fmt.Println("Todo added successfully")
}
func ListTodos() {
	todos, _ := LoadTodos()

	if len(todos) == 0 {
		fmt.Println("No todos found")
	} else {
		fmt.Println("List of todos")
		for _, todo := range todos {
			status := "Not Completed"
			if todo.Completed {
				status = "Completed"
			}
			fmt.Printf("[%s] %d - %s\n", status, todo.Id, todo.Title)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an action")
		fmt.Println("  go run main.go list           → Afficher les tâches")
		fmt.Println("  go run main.go add 'Tâche'    → Ajouter une tâche")
		return
	}

	action := os.Args[1]
	switch action {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a title")
			return
		}
		title := os.Args[2]
		addTodo(title)
	case "list":
		ListTodos()
	default:
		fmt.Println("Invalid action. Please use add or list")
	}
}
