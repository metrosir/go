package main

import "fmt"

//变量是一种使用方便的占位符，用于引用计算机内存地址
//go 语言的指针不能运算
//go 语言的值
//go 语言是值传递还是引用传递呢？c和c++可以值传递、也可以引用传递
//go 语言只有值传递一种方式，不用考虑引用传递的方式
//go 语言中值传递和指针是怎么配合的？

//func swap(a, b *int)  {
//	*b, *a = *a, *b
//}
//
////func swap(a, b int) (int , int)  {
////	return  b, a
////}
//
//func main()  {
//	a, b := 3, 4
//
//	a, b = swap(&a, &b)
//	fmt.Println(a, b)
//}

//func swap(a, b *int) {
//	*b, *a = *a, *b
//}
//
//func main() {
//	a, b := 3, 4
//	swap (&a, &b)
//	fmt.Println(a, b)
//}

//func swap(a, b int)  (int, int) {
//	return  b, a
//}
//
//func main()  {
//	a, b := 5, 6
//	a, b = swap(a, b)
//	fmt.Println(a, b)
//}

//go 语言的取地址符是&，放到一个变量前面就会返回相应变量的内存地址
func main() {
	a := 10
	fmt.Printf("变量的地址：%x\n", &a)
}
