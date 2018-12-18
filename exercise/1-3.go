package main

import (
	"fmt"
	"io/ioutil"
)

//基础语法
// -- 条件语句与循环 --

func main() {
	result := ifelse1(1)
	fmt.Println(result)

	filename := "abc.txt"
	content, err := readFile(filename)
	if err {
		fmt.Printf("%s\n", content)
	} else {
		fmt.Println("get filename: " + filename + " error")
	}
}

//
func ifelse1(a int) int {
	if a > 100 {
		return 100
	} else if a > 50 {
		return 50
	} else {
		return 0
	}
}

func ifelse2() {
	if a, b := 1, 2; a+b > 3 {
		fmt.Println(a, b)
	}

	//fmt.Println(a, b) // 这样定义是错误的，a, b 的是if条件里面定义的，作用域仅限于if中使用
}

//example 利用if语句判断读取文件时是否有异常
func readFile(filename string) ([]byte, bool) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return content, false
	} else {
		return content, true
	}
}
