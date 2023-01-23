package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func receive(receive bool, anotherDone chan bool) {
	if receive {
		time.Sleep(10 * time.Second)
		fmt.Println("from receive")
		anotherDone <- true
	}
}

func main() {
	done := make(chan bool)
	anotherDone := make(chan bool)

	go hello(done)
	toSend := <-done
	go receive(toSend, anotherDone)
	<-anotherDone
	fmt.Println("main function")
}
