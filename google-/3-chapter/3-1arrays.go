package main

import "fmt"

//定义数组的几种方法
//在go语言中一般不直接使用数组 go语言使用最常见的是切片
//go 语言中定义数组数字要在类型的前面
//go 语言中[10]int 和 [20]int 是不同的类型
//调用 func f(arr [10]int)会拷贝数组，因为go语言中只有值传递的方式
//数组是值类型
func main() {
	//1.规定数组中共有5个int
	var arr1 [5]int
	//2.这种方式定义需要一个初值
	arr2 := [3]int{1, 3, 5}
	//3.用编译器去数数组中有多少个int
	arr3 := [...]int{2, 4, 6, 8, 10}
	//4.二维数字，数组中有4行5列，4个长度为5的长度数组
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//一般遍历 使用range 关键字，range关键字能获取数组的下标
	//range意义明确，美观
	//
	//遍历下标来获得值
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	//遍历下标和值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//只遍历值
	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Println("printArray(arr1)")
	printArray(arr1)
	printArrayP(&arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)
	printArrayP(&arr3)

	//因为数组是值类型，所以这里arr1 的零号元素被改变了但以下的运行结果没有被改变
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)

}

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
	//arr[0] = 100
}

//使用指向数组的一个指针
func printArrayP(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
