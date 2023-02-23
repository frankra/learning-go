package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

func testPanic() {
	var mySlice []int
	j := mySlice[0]
	log.Println("Hello, playground %d", j)
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, strings.ReplaceAll(string(debug.Stack()), "\n", ";;"))
		}
	}()

	//create your file with desired read/write permissions
	f, err := os.OpenFile("filename", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}

	//set output of logs to f
	testPanic()
	//defer to close when you're done with it, not because you think it's idiomatic!
	f.Close()
}
