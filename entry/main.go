package main

import  (
	"fmt"

	"kaka.com/maxConsecutiveOnes"
)

func main() {
	input := []int{0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1}

	maxx := maxConsecutiveOnes.FindMaxConsecutiveOnes(input)
	fmt.Printf("Maxx: %v\n", maxx)
}
