package todoitem

import (
	"errors"
)

// TodoItem represents a todo item within a todo list
type TodoItem struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Priority int    `json:"priority"`
}

// TodoItemRepository defines the interface for todo item data storage and retrieval operations
type TodoItemRepository interface {
	CreateTodoItem(todoListID int64, todoItem *TodoItem) (*TodoItem, error)
	GetTodoItemByID(todoListID, todoItemID int64) (*TodoItem, error)
	UpdateTodoItem(todoListID int64, todoItem *TodoItem) error
	DeleteTodoItem(todoListID, todoItemID int64) error
}

// Validation function for TodoItem
func (ti *TodoItem) Validate() error {
	if ti.Title == "" {
		return errors.New("todo item title is required")
	}
	if ti.Priority < 1 {
		return errors.New("todo item priority should be greater than or equal to 1")
	}
	return nil
}
