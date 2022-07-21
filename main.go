package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())

	concurrent := 50
	jobs := make(chan int, concurrent)
	results := make(chan string)
	files := 100

	// Start the workers
	for workerId := 0; workerId < concurrent; workerId++ {
		go worker(workerId+1, jobs, results)
	}

	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())

	// Send jobs to the workers
	for i := 0; i < files; i++ {
		jobs <- i
	}

	// Close the jobs channel
	close(jobs)

	for i := 0; i < files; i++ {
		fmt.Println(<-results)
	}

	close(results)
}

func worker(workerId int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		start := time.Now()
		random := rand.Intn(2)
		time.Sleep(time.Duration(random) * time.Second)
		end := time.Now()

		results <- fmt.Sprintf("THREAD_%d: Finished file%d.txt in %d milliseconds", workerId, job, end.Sub(start).Milliseconds())
	}
}
