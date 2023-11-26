package todo

import (
	status "github.com/tothbence9922/gotodo/internal/status/model"
)

type Todo struct {
	Id          uint64
	Title       string
	Description string
	Status      status.Status
}
