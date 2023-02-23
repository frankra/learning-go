package main

import "fmt"

type Person struct {
	name *string
}

func (p *Person) GetName() *string {
	return p.name
}

func main() {
	name := "Rafa"
	p1 := Person{
		name: &name,
	}

	name = "Bharath"

	fmt.Println(*p1.GetName())
}
