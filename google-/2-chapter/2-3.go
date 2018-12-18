package main

import (
	"fmt"
	"io/ioutil"
)

//条件语句
func main() {
	const filename = "abc.txt"
	//if contents, err := ioutil.ReadFile(filename); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	contents, err := file_get_contents(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func file_get_contents(filePath string) ([]byte, error) {
	if contents, err := ioutil.ReadFile(filePath); err != nil {
		return nil, fmt.Errorf("file path error: " + filePath)
	} else {
		return contents, nil
	}
}
