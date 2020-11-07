package main

import (
	"fmt"

	"kaka.com/maxConsecutiveOnes"
	"kaka.com/greetings"
)

func main() {
	message, error := greetings.Hello("")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(message)
	}

	maxx := maxConsecutiveOnes.FindMaxConsecutiveOnes([]int{0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1})
	fmt.Println(maxx)
}
