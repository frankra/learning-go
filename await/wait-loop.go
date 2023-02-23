package main

import (
	"sync"
	"time"
)

type EventToBeSent struct {
	ID string
	Target string
	EventType
}

func main() {
	go notificationWorker()
	time.Sleep(1 * time.Hour)
}

func notificationWorker() {
	for {
		var wg sync.WaitGroup
		wg.Add(2)

	}
}

func getEventsToBeSent()