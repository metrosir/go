package main

import "fmt"

//向slices 添加元素

//添加元素时如果超越 cap， 系统会重新分配更大的底层数组，如果 arr 最终没有被使用到系统会通过垃圾回收机制吧arr 回收掉；
//由于值传递的关系，必须接收append的返回值
func main()  {
	fmt.Println("Creating slice")
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println(s1,s2,s3,s4,s5)
	e()

	fmt.Println("Copying slice")
	e1()

	fmt.Println("Deleting slice")
	e2()
}

func e2() {
	//delete
	s1 := []int{2,4,6,8}
	s2 := make([]int, 16)
	copy(s2, s1)
	s2 = append(s2[:3], s2[4:]...)
	PrintSlices(s2)
}


func e1()  {
	//copy
	s1 := []int{2,4,6,8}
	s2 := make([]int, 16)
	copy(s2, s1)
	PrintSlices(s2)
}

func e()  {
	var s []int
	for i := 0; i < 100 ; i++ {
		PrintSlices(s)
		s = append(s, i * 2 +1)
	}

	fmt.Println(s);
	s1 := []int{2,4,6,8}
	PrintSlices(s1)

	//初始化一个len=10 的slices
	s2 := make([]int, 10)
	PrintSlices(s2)
	//初始化一个len=10 cap=32 的slices
	s3 := make([]int, 10, 32)
	PrintSlices(s3)
}

func PrintSlices(s []int)  {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}