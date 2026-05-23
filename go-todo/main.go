package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

const filename = "todos.json"

func loadTodos() []Todo {
	data, err := os.ReadFile(filename)

	if err != nil {
		return []Todo{}
	}

	var todos []Todo

	json.Unmarshal(data, &todos)

	return todos
}

func saveTodos(todos []Todo) {
	data, _ := json.MarshalIndent(todos, "", "  ")
	os.WriteFile(filename, data, 0644)
}

func addTodo(task string) {
	todos := loadTodos()
	maxID := 0
	for _, t := range todos {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	newTodo := Todo{
		ID:   maxID + 1,
		Task: task,
		Done: false,
	}

	todos = append(todos, newTodo)
	saveTodos(todos)
	fmt.Println("Todo added successfully")
}

func listTodos() []Todo {
	todos := loadTodos()
	if len(todos) == 0 {
		fmt.Println("No todos found")
		return []Todo{}
	}

	fmt.Println("Todo List:")
	for _, todo := range todos {
		status := " "
		if todo.Done {
			status = "X"
		}
		fmt.Printf("%d. [%s] %s\n", todo.ID, status, todo.Task)
	}
	return todos
}

func markDone(id int) {
	todos := loadTodos()
	for i, todo := range todos {
		if todo.ID == id {
			todo.Done = true
			todos[i] = todo
			saveTodos(todos)
			fmt.Println("Todo marked as done")
			return
		}
	}
	fmt.Println("Todo not found")
}

func deleteTodo(id int) {
	todos := loadTodos()
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			saveTodos(todos)
			fmt.Println("Todo deleted successfully")
			return
		}
	}
	fmt.Println("Todo not found")
}

func main() {
	if len(os.Args) >= 2 {

		command := os.Args[1]

		switch command {
		case "add":
			if len(os.Args) < 3 {
				fmt.Println("Please provide a task to add")
				return
			}
			task := os.Args[2]
			addTodo(task)
			return

		case "list":
			listTodos()
			return

		case "complete":
			if len(os.Args) < 3 {
				fmt.Println("Please provide a task ID to mark as done")
				return
			}

			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Please provide a valid task ID")
				return
			}
			markDone(id)
			return
		case "delete":
			if len(os.Args) < 3 {
				fmt.Println("Please provide a task ID to delete")
				return
			}

			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Please provide a valid task ID")
				return
			}
			deleteTodo(id)
			return
		case "help":
			fmt.Println(`Usage
							add <task> - Add a new task
							list - List all tasks
							complete <id> - Mark a task as done
							delete <id> - Delete a task
							help - Show this help message`)
			return
		default:
			fmt.Println("Invalid command")
			return
		}
	}

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			todos := listTodos()
			json.NewEncoder(w).Encode(todos)
		case "POST":
			var newTodo Todo
			json.NewDecoder(r.Body).Decode(&newTodo)
			addTodo(newTodo.Task)
			json.NewEncoder(w).Encode(newTodo)
		case "PUT":
			var updatedTodo Todo
			json.NewDecoder(r.Body).Decode(&updatedTodo)
			markDone(updatedTodo.ID)
			json.NewEncoder(w).Encode(updatedTodo)
		case "DELETE":
			var deletedTodo Todo
			json.NewDecoder(r.Body).Decode(&deletedTodo)
			deleteTodo(deletedTodo.ID)
			json.NewEncoder(w).Encode(deletedTodo)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", nil)
}
