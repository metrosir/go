package main

import "fmt"

//基础语法
// -- 指针 --

//go语言的参数传递是值传递

func main() {
	a, b := 1, 2
	example1(&a, &b)
}

func example1(a, b *int) {
	fmt.Println(*a, b) // 理解：a，b *int 存的是int类型的地址，当对指针类型的变量 *a 时，就是取出地址对应的值
}
