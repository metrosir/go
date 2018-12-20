package main

import "fmt"

//内建容器
// -- 数组 --

func main() {

	//定义数组
	var arr1 [3]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5} //用编译器去数数组中有多少个int
	arr4 := [2][4]int{}             //定义一个两行四列的数组
	fmt.Println(arr1, arr2, arr3, arr4)

	//遍历数组
	for _, v := range arr2 {
		fmt.Println(v)
	}

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	example9([5]int{1, 2, 3, 4, 5}, [10]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

//注意，这里的[5]int 和 [10]int是不同的类型，在go语言中一般不直接使用数组，而是切片slice
func example9(arr [5]int, arr1 [10]int) {
	for _, v := range arr {
		fmt.Println(v)
	}

	for _, v := range arr1 {
		fmt.Println(v)
	}
}
