package todolist

import (
	"errors"
)

type TodoList struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	TodoItems   []TodoItem `json:"todo_items"`
}

// TodoListRepository defines the interface for todo list data storage and retrieval operations
type TodoListRepository interface {
	CreateTodoList(todoList *TodoList) (*TodoList, error)
	GetTodoListByID(id int64) (*TodoList, error)
	GetAllTodoLists() ([]*TodoList, error)
	UpdateTodoList(todoList *TodoList) error
	DeleteTodoList(id int64) error

	CreateTodoItem(todoListID int64, todoItem *TodoItem) (*TodoItem, error)
	GetTodoItemByID(todoListID, todoItemID int64) (*TodoItem, error)
	UpdateTodoItem(todoListID int64, todoItem *TodoItem) error
	DeleteTodoItem(todoListID, todoItemID int64) error
}

// Validation functions for TodoList and TodoItem
func (tl *TodoList) Validate() error {
	if tl.Name == "" {
		return errors.New("todo list name is required")
	}
	return nil
}

func (ti *TodoItem) Validate() error {
	if ti.Title == "" {
		return errors.New("todo item title is required")
	}
	if ti.Priority < 1 {
		return errors.New("todo item priority should be greater than or equal to 1")
	}
	return nil
}
