package main

import "fmt"

type Animal interface {
	MakeNoise() string
}

type Cat struct{}

func (c *Cat) MakeNoise() string {
	return "meow"
}

type Dog struct{}

func (d *Dog) MakeNoise() string {
	return "woof"
}

func makeNoise(a Animal) {
	fmt.Println(a.MakeNoise())
}
func main() {
	cat := Cat{}
	dog := Dog{}

	makeNoise(&cat)
	makeNoise(&dog)
}
