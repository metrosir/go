package __go

import "fmt"

//基础语法
// -- 函数 --

func main() {
	a, b := exampl1(1, 2)
	x, y := exampl2(8, 9)
	i, j := exampl3(100, 1)

	fmt.Println(a, b, x, y, i, j)
}

//有多个返回值
func exampl1(a, b int) (int, int) {
	return a + b, a * b
}

//为多个返回值起名字 （仅用于简单函数）
func exampl2(a, b int) (x, y int) {
	x = a + b
	y = a * b
	return
}

//输出错误
func exampl3(a, b int) (int, error) {
	if a+b > 10 {
		return a + b, fmt.Errorf("%s", "error")
	} else {
		return a + b, nil
	}
}
