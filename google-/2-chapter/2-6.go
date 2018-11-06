package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数
// go 语言是函数式编程语言
// 函数返回多个值时可以起名字
// 仅用于非常简单的函数
// 对于调用者而言没有区别
// go 语言有可变参数列表，没有默认参数，可选参数
func eval(a, b int, op string)  (int, error)  {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// _ 跳过第二个返回值
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("op error: " + op)
		//panic("op error: " + op)
	}
}

//函数可以有两个返回值
func div(a, b int) (q int, r int)  {
	return a / b, a % b
}

// 函数式编程，函数的参数也可以是一个函数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	//获取函数 op 的名称
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args " + "(%d, %d)\n", opName, a, b)
	return op(a, b)
}

//重写 math.Pow 函数
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//go 语言有可变参数列表 ... 表示函数的参数个数由调用者来决定
func sum(numbers ...int) int  {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return  s
}

func main()  {
	//两个返回值的常用的场景，我们通常情况第二个返回值是return error的情况
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	//fmt.Println(eval(3, 4, "x"))
	a, b := div(3, 4)
	fmt.Println(a, b)

	//fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(func(i int, i2 int) int {
		return int(math.Pow(float64(i), float64(i2)))
	}, 3, 4))

	fmt.Println(sum(1,2,3,4,5,6,9))
}