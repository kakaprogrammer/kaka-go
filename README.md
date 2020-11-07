# Install go on linux
- cmd: tar -C /usr/local -xzf go1.15.4.linux-amd64.tar.gz
- cmd: export PATH=$PATH:/usr/local/go/bin
- test cmd:  go version


## Variable 
- Khai báo kiểu 1: 
>><code class="language-go">var a int
a = 10
</code>

- Khai báo biến kiểu 2(vừa khai báo vừa khởi tạo gía trị):
>><code class="language-go">var a string = "Kaka"</code>

- Khai báo biến kiểu 3():
>><code class="language-go">vara := "Kaka"</code>




## Constant
- 

- Nếu mà chúng ta khai báo 1 enum mà những cáci giá trị hằng số ở bên trong có 1 quy tắt nào đó kiểu tuyến tính thì ta nên sử dụng iota để khởi tạo nó

