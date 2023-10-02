package main

import (
	"fmt"
	"reflect"
)

type SomeInterface interface {
}

type SomeStruct struct {
}

func ReturnInterface() SomeInterface {
	return nil
}

func ReturnStruct() *SomeStruct {
	return nil
}

func main() {
	returnType := reflect.TypeOf(ReturnInterface).Out(0)
	fmt.Print(returnType)
	returnType = reflect.TypeOf(ReturnStruct).Out(0)
	fmt.Print(returnType)
}
