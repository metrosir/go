package main

import "fmt"

// if 的条件里可以赋值
// if 的条件里赋值的变量作用域就在这个if语句里

//func main()  {
//	const filename  = "abc.txt"
//	contents, err := ioutil.ReadFile(filename)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Printf("%s\n", contents)
//	}
//}

//func main()  {
//	const filename  = "abc.txt"
//	if contents,err :=  ioutil.ReadFile(filename); err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Printf("%s\n", contents)
//	}
//}

// switch 会自动break，除非使用fallthrough
// default 中的panic 相当于报错，停止程序的运行

func grade(score int ) string {
	g := ""
	// switch 里面可以不用表达式，不用表达式的时候在case 里面加入条件就可以了
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))

	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main()  {
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		grade(-3),
		)
}