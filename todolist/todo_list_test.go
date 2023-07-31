package todolist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lpgod/task/todolist"
)

func TestTodoListValidation(t *testing.T) {
	// Valid todo list
	validTodoList := &todolist.TodoList{
		Name:        "Groceries",
		Description: "Buy groceries for the week",
		TodoItems: []todolist.TodoItem{
			{Title: "Milk", Priority: 2},
			{Title: "Eggs", Priority: 1},
		},
	}
	assert.NoError(t, validTodoList.Validate())

	// Invalid todo list (missing name)
	invalidTodoList := &todolist.TodoList{
		Description: "Invalid todo list without name",
		TodoItems: []todolist.TodoItem{
			{Title: "Task 1", Priority: 1},
		},
	}
	assert.Error(t, invalidTodoList.Validate())
}

func TestTodoItemValidation(t *testing.T) {
	// Valid todo item
	validTodoItem := &todolist.TodoItem{
		Title:    "Task 1",
		Priority: 2,
	}
	assert.NoError(t, validTodoItem.Validate())

	// Invalid todo item (missing title)
	invalidTodoItem := &todolist.TodoItem{
		Priority: 1,
	}
	assert.Error(t, invalidTodoItem.Validate())

	// Invalid todo item (invalid priority)
	invalidTodoItem.Priority = 0
	assert.Error(t, invalidTodoItem.Validate())
}
