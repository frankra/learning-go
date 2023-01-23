package main

import (
	"fmt"
	"time"
)

func worker(ch chan string) {
	i := 1
	for {
		time.Sleep(500 * time.Millisecond)
		ch <- fmt.Sprintf("working%d", i)
		i++
	}
}

func main() {
	ch := make(chan string)
	go worker(ch)
	go observe(ch, func(i string) {
		fmt.Println("callback called", i)
	})
	fmt.Println("continue after ")
	time.Sleep(1 * time.Hour)
}

func observe(ch chan string, fn func(string)) {
	fn(<-ch)
	observe(ch, fn)
}
