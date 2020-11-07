package maxConsecutiveOnes

import (
	"fmt"
)

func FindMaxConsecutiveOnes(nums []int) int {
	genMap := [][]int{} // khoi tao mang 2 chieu voi khoi tao bang rong
	m := []int{}        //khoi tao mang 1 chieu ben trong
	maxx := 0
	for i := 0; i < len(nums); i++ { //duyen tu dau den cuoi cua mang
		if nums[i] == 0 { //neu value = 0 : chot mang tru truong hop i=0
			if i == 0 {
				continue
			}
			genMap = append(genMap, m)
			if maxx < len(m) {
				maxx = len(m) //lay max len cua tung mang
			}
			m = []int{} //chot xong khoi tao tao mang con moi

		} else {

			m = append(m, 1)
			if i == len(nums)-1 {
				genMap = append(genMap, m)
				if maxx < len(m) {
					maxx = len(m) //lay max len cua tung mang
				}
			}
		}
	}

	fmt.Printf("Output: %v\n", genMap)
	return maxx
}
