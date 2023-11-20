package hello

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if len(name) == 0 {
		return name, errors.New("Empty name!")
	}
	return fmt.Sprintf("Hello, %s", name), nil
}
