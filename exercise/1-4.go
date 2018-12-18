package main

import (
	"bufio"
	"fmt"
	"io"
)

//基础语法
// -- for循环 --

func main() {
	example()
	example2(10)
}

func example() {
	//赋值语句，判断语句，递减语句
	for i := 5; i > 0; i-- {
		fmt.Println(i)
	}
}

func example2(n int) {
	//无赋值
	for ; n > 0; n /= 2 {
		fmt.Println(n)
	}
}

func example3(file io.Reader) {
	//仅赋值
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text)
	}
}

func example4() {
	//死循环
	for {
		fmt.Println(1)
	}
}
