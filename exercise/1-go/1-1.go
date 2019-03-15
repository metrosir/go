package __go

import (
	"fmt"
	"sort"
	"strings"
)

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
	sorts()
	bytestest()
	string1();
}

type User struct {
	Id int
	Age  int
	Name string
}

//sort 排序应用
func sorts()  {
	data := []User {
		{Id: 1, Age: 10, Name:"a1"},
		{Id: 4, Age: 20, Name:"a2"},
		{Id: 3, Age: 11, Name:"a3"},
		{Id: 2, Age: 19, Name:"a4"},
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Age < data[j].Age
	})

	for _, v := range  data {
		fmt.Println(v.Id, v.Age, v.Name, '\n')
	}
}

func bytestest()  {
	str := "cat"
	s := []byte(str)
	strs := string(s)
	fmt.Printf("%T===%#v\n", str, s, strs)
	
}


//golang pkg  https://golang.org/pkg/
//常用包 md5 utf8 bytes strings sort big

func string1()  {
	s1 := "a哈吃的"
	s1_count := strings.Count(s1, "")
	fmt.Println("\n", s1_count)

}


