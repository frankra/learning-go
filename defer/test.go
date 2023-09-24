package main

import "fmt"

func testDefer() int {
	if true {
		return 2
	}
	defer fmt.Print("ok")
	return 1
}
func main() {
	testDefer()
}
