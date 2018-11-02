package main

import (
	"fmt"
	"io/ioutil"
)

//条件语句
func main() {
	const filename  = "abc";
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
