package main

import (
	"fmt"
	"math"
	"runtime/debug"
)

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, string(debug.Stack()))
		}
	}()
	var c *Circle
	fmt.Printf("(%v, %T)\n", c, c)
	fmt.Println(c.area())
}
