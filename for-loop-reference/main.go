package main

import "fmt"

type TestStruct struct {
	Value string
}

func main() {
	testData := []TestStruct{
		{
			Value: "test 1",
		},
		{
			Value: "test 2",
		},
		{
			Value: "test 3",
		},
	}

	copy := []*string{}
	for _, elem := range testData {
		fmt.Println(elem.Value)
		copy = append(copy, &elem.Value)
	}

	fmt.Println("Check what we copied")
	for _, str := range copy {
		fmt.Println(*str)
	}
}
