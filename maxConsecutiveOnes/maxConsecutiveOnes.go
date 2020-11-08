package maxConsecutiveOnes

import (
	"fmt"
)

//Tham khao
/**
Mô tả: 
- ta khai báo 2 biến max tạm với biến đếm count và khởi tạo gía trị bằng 0
- ta dùng vòng lặp duyêt tất cả cac value phần tư của mảng
	+ trương hợp biến count + value (tăng lên 1 đv) mà lớn hơn tổng đó thì gán cho tổng đó
	+ trường ko thoãn đk bỏ qua
- biến count giống như biến max nhưng nhân với value phần tử để value trong trường hợp 0 sẽ reset về 0
*/
func FindTemp(a []int) int {
	best, cnt := 0, 0
	for _, v := range a {
		if cnt + v > best { best = cnt + v} //neu v  = 1 thi cong don cho best 
		cnt = (cnt + v)*v
	}
	return best

	//Phan tich
	//TESTCASE 1: [0,1,1,0,0]
	//step 1: best = 0 + 0 = 0  ~~~ crt = (crt + v)*v = 0
	//step 2: best = 0 + 1 = 1  ~~~ crt = (0 + 1)*1 = 1
	//step 3: best = 1 + 1 = 2  ~~~ crt = (1+1)* 1 = 2
	//step 4: best = 2 + 0 = 2  ~~~ crt = (2+0)* 0 = 0
//==> Thuật toán ở đây gần giống với thuật toán dưới nhưng ko khai báo cờ reset biến count mà nhân count với v khi = 0 sẽ reset lại biến count 
}

func FindMaxxOnes(nums []int) int {
	var maxx int = 0
	var flag bool = false //co bao chot xong 1 chuoi
	var count = 0
	//duyet array
	for i := 0; i < len(nums); i++ {
		//
		//if nums[0] == 0 {
			//flag = true
			//continue
		//} //gia tri dau tien cua mang = 0 thi bo qua ko tinh

		if nums[i] == 0 && flag == false {
			flag = true

			if i == 0 {
				continue
			}
			if maxx < count { maxx = count }
			count = 0
		}

		if nums[i] == 1 {
			count++
			flag = false
		}

		fmt.Printf("debug Count %v\n", count)
	}

	return maxx
}

func FindMaxConsecutiveOnes(nums []int) int {
	genMap := [][]int{} // khoi tao mang 2 chieu voi khoi tao bang rong
	m := []int{}        //khoi tao mang 1 chieu ben trong
	maxx := 0
	for i := 0; i < len(nums); i++ { //duyen tu dau den cuoi cua mang
		if nums[i] == 0 { //neu value = 0 : chot mang tru truong hop i=0
			if i == 0 {
				continue
			}
			if len(m) > 0 { genMap = append(genMap, m)}
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
