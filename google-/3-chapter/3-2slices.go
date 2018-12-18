package main

import "fmt"

//go 语言 slice 切片
//go 语言中我们一般不用数组，我们一般使用slice
//go 语言中 slice 本身是没有数据的，是对底层array的一个view
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[2:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("After updateSlice(s1) ")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2) ")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
}

//go 语言slice是 array数组的一个view，一个视图
func updateSlice(s []int) {
	s[0] = 100
}
