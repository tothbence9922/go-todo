package todo

import (
	"errors"

	status "github.com/tothbence9922/gotodo/internal/status/model"
	todo "github.com/tothbence9922/gotodo/internal/todo/model"
)

type Todo interface {
	Create(id uint64, title string, description string, status status.Status) (*todo.Todo, error)
}

func Create(id uint64, title string, description string, status status.Status) (*todo.Todo, error) {
	if len(title) == 0 || len(description) == 0 {
		return nil, errors.New("Todo Create parameters could not be validated.")
	}
	return &todo.Todo{Id: id, Title: title, Description: description, Status: status}, nil
}
