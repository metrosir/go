package main

import "fmt"

//基础语法
// -- 变量定义 --
//变量要先声明，再赋值

//变量定义
func main() {
	//声明
	var a int    //声明int类型的变量
	var b [3]int //声明int类型的数组
	var c []int  //声明int类型的切片【slice】
	var d *int   //声明int类型的指针
	fmt.Println(a, b, c, d)

	//赋值
	a = 10
	b[0] = 10
	fmt.Println(a, b)

	//同时声明和赋值
	var e = 10
	f := 10
	g, h, i, j := 1, 2, true, "def"
	fmt.Println(e, f, g, h, i, j)
}
