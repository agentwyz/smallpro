package main

import (
	"flag"
	"fmt"
	//"os"
)

func main() {
	// contents, err := os.ReadFile("test.txt")

	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// 	return
	// }

	// fmt.Printf("contents of file %s", string(contents))
	readC()
}

//we read a file from commmand line
func readC() {
	//第一个是以fpath为开头
	//第二个是获取默认值
	//第三个是用户usage
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
}