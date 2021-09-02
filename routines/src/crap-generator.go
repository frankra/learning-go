package main

import (
    "fmt"
    "math/rand"
    "time"
	"sync"
	"runtime"
)

func generateCrapWorker(size int, min int, max int, out chan []int, wg *sync.WaitGroup) {
	crap := make([]int, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		crap[i] = rand.Intn(1000) + min
	}
	fmt.Println(crap)
	out <- crap

	defer wg.Done()
}

func generateCrap(size int, min int, max int) []int {
	cpuCount := runtime.NumCPU()
	out := make(chan []int, cpuCount)
	var wg sync.WaitGroup

	workerSize := int(size / cpuCount)

	for i := 0; i < cpuCount; i++ {
		wg.Add(1)
		go generateCrapWorker(workerSize, max, min, out, &wg)
	}

	wg.Wait()
	close(out)

	result := make([]int, 0);
	for slice := range out {
		
		result = append(result, slice...)
	}
	
	return result
}

func main(){
	start := time.Now()

	crap := generateCrap(1000000, 10, 30)
	fmt.Println(crap)

	fmt.Println(time.Now().Sub(start))
}