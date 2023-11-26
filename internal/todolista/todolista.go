package todolist

import (
	"errors"

	todo "github.com/tothbence9922/gotodo/internal/todo/model"
	todolist "github.com/tothbence9922/gotodo/internal/todolist/model"
)

type TodoList interface {
	Create(title string) (*todolist.TodoList, error)
}

func Create(title string) (*todolist.TodoList, error) {
	if len(title) == 0 {
		return nil, errors.New("TodoList Create parameters could not be validated.")
	}
	emptySlice := make([]todo.Todo, 0)
	return &todolist.TodoList{Title: title, Todos: emptySlice}, nil
}
