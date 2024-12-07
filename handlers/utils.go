package handlers

import "github.com/devkaare/todo/database"

var todos []*database.Todo

func getTodoByID(id int) (*database.Todo, bool) {
	for i, todoFromTodos := range todos {
		if todoFromTodos.ID == id {
			// Return todo at current index in todos
			return todos[i], true
		}
	}
	return &database.Todo{}, false
}

func deleteTodoByID(id int) bool {
	for i, todoFromTodos := range todos {
		if todoFromTodos.ID == id {
			// Delete todo at current index in todos
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}

func updateTodo(todo *database.Todo) bool {
	for i, todoFromTodos := range todos {
		if todoFromTodos.ID == todo.ID {
			// Update todo at current index in todos
			todos[i] = todo
			return true
		}
	}
	return false
}
