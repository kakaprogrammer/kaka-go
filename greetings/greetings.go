package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {

	//

	if name == "" {
		fmt.Printf("%T\n",errors.New("qqqq"))
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v", name)
	return message, nil

}
