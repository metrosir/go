package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//for 的条件里面不需要括号
//for 的条件里面可以忽略初始条件，结束条件，递增表达式

//数字转二进制表达式
func convertToBin(n int) string  {
	result := ""
	//省略起始条件
	for ; n>0; n /= 2 {
		lsb := n % 2
		//strconv.Itoa() 数字转字符串
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//循环输出文件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		//报错
		panic(err)
	}
	//读取文件
	scanner := bufio.NewScanner(file)
	//省略起始条件和递增条件，go语言中没有while，for中不传起始条件和递增条件也能实现while
	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
}

func foever() {
	//for条件中什么都不写，就是一个死循环
	for  {
		fmt.Println("abc")
	}
}

func main()  {
	fmt.Println(
			convertToBin(5),
			convertToBin(13),
			convertToBin(19123921),
			convertToBin(0))
	printFile("abc.txt")
	foever()
}

//for，if后面的条件没有括号
//if条件里也可定义变量
//没有while 被for包括
//switch 不需要break，可以直接switch多个条件