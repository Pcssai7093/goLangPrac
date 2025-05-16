// greetings/greetings
package greetings

import (
	"errors"
	"fmt"
)

func Greet(name string) (string, error) {

	if name == "" {
		return "", errors.New("name not found")
	}

	msg := fmt.Sprintf("Hello! %s", name)
	return msg, nil
}
