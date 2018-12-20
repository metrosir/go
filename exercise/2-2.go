package main

import "fmt"

func main() {

	//前开后闭 ?（取前不取后）
	arr := [5]int{1, 2, 3, 4, 5}

	//arrs := []int{}

	arr1 := arr[1:2]
	arr2 := arr[:4]
	arr3 := arr[3:]
	arr4 := arr[:]
	fmt.Println(arr, arr1, arr2, arr3, arr4)

	//视图 切片是数组的"视图"，即"引用"
	example10(arr[1:2])
	fmt.Println(arr)

	//切片的切片依然是对一个数组的引用
	example10(arr2)
	fmt.Println(arr, arr2)

	fmt.Println(arr2, len(arr2), cap(arr2))

	//直接创建切片
	s := []int{1, 2, 3}
	var s1 []int
	s2 := make([]int, 16, 32)
	fmt.Println(s, s1, s2)

	//arrss := [5] int{1,2,3,4,5}
	//arrs1 := arrs[1:3]
	//fmt.Println(arrs[1],arrs1)
	b := [...]string{"x", "a"}
	fmt.Println(b[0])
	c := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Println(c[1:4])
	cc := make([]byte, 5)
	fmt.Println(cc)

	example11()
	example12()

	p := example13([]int{1, 2, 3}, func(i int) bool {
		if i%2 == 0 {
			fmt.Println(i)
			return true
		}
		return false
	})
	fmt.Println(p)
}

//[]中不写具体的大小，表示是切片，引用传递
func example10(arr []int) {
	arr[0] = 111
}

func example11() {
	s := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	t := make([]byte, len(s), (cap(s)+1)*2)
	//for i := range s {
	//	t[i] = s[i]
	//}
	copy(t, s)
	s = t
	fmt.Printf("", s)
}

//将一个切片追加到另一个切片的尾部
func example12() {
	a := []string{"a", "b", "c"}
	b := []string{"d", "e", "f"}
	a = append(a, b...)
	fmt.Println("\n", a)
}

func example13(s []int, fn func(int) bool) []int {

	var p []int
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
