package main

import (
	"fmt"
	"kaka.com/maxConsecutiveOnes"
)

func main() {
	input := []int {0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0}
	//maxx := maxConsecutiveOnes.FindMaxConsecutiveOnes(input)
	//maxx := maxConsecutiveOnes.FindMaxxOnes(input)
	maxx := maxConsecutiveOnes.FindTemp(input)
	fmt.Println(maxx)
}
