package __go

import "fmt"

//基础语法
// -- 常量定义 --

//常量可以作为各种类型调用

func main() {
	//常量可以作为各种类型调用，这里即可以用作int类型，也可以用作float
	const filename = "ab"
	const a, b = 3, 4
	const (
		c = 1
		d = 2
		e = 3
	)

	const (
		f = iota //自增值，初始为0
		g
		h
	)

	fmt.Println(f, g, h)
}
