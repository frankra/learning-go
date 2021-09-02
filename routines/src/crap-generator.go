package main

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func generateCrapWorker(size int, min int, max int, out chan []int, wg *sync.WaitGroup) {
	crap := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		crap[i] = rand.Intn(max - min) + min
	}
	out <- crap

	defer wg.Done()

	
}

func GenerateCrap(size int, min int, max int) []int {
	cpuCount := runtime.NumCPU()
	out := make(chan []int, cpuCount)
	var wg sync.WaitGroup

	workerSize := int(size / cpuCount)

	for i := 0; i < cpuCount; i++ {
		wg.Add(1)
		go generateCrapWorker(workerSize, min, max, out, &wg)
	}

	wg.Wait()
	close(out)

	result := make([]int, 0);
	for slice := range out {
		
		result = append(result, slice...)
	}
	
	return result
}

// func main(){
// 	start := time.Now()

// 	GenerateCrap(1000000, 10, 999999)

// 	fmt.Println(time.Now().Sub(start))
// }