# Challenge 1
- Yêu cầu bài toán tính số lượng tối đa bình nước phải uống
- Dữ liệu cho:
	+ số lượng bình nước đầy
	+ số lượng quy đổi từ bình nước rỗng ra bình nước đầy (giả sử sau khi uống hết 10 bình thì từ 10 bình rỗng đó ta có thể qui đổi ra bình nước đầy để thực hiện uống hết đến khi nào không đổi được nữa)

- Ví dụ: cho 10 bình nước đầy và số qui đổi là 3 (3 bình rỗng đổi được 1 bình đầy)
--> mô tả thao tác
+ uống hết 10 bình nước đầy được 10 bình nước rỗng  "Đã uống được 10 bình"
+ lấy 10 bình rỗng đổi ra 3 bình nước đầy còn dư 1 bình rỗng
+ tiếp tục uống hết 3 bình đầy "uống thêm 3 bình đầy -> 13 bình đầy"
+ lấy 3 bình rỗng bước trước + 1 bình rỗng lúc trước nữa là 4 bình rỗng đỏi đuọc 1 bình đầy dư 1 bình rỗng
+ uổng tếp 1 bình đầy đó "uống thêm 1 bình --> 13+ 1 =14 bình đầy" dư lại 1 bình rỗng éo đổi được nữa

* Ai giải được bài toán sẽ được tặng 1 ly cafe trị gía 19k. Kaka!

# Challenge 2
- Mô tả ban đầu: cho 1 mảng các số nguyên 
- Yêu cầu: Tìm tất cả các cặp số có value bằng nhau
Với mỗi cặp số có các vị trí i, j sao cho i < j

- Ví dụ: 
	input: nums = [1, 3, 5, 1, 5, 3, 8]
	output: 3 (cặp số như bên dưới) 
	ex(i, j): (0, 3); (1, 5); (2, 4) 
