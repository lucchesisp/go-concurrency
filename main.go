package main

import (
	"fmt"
	"time"
)

func worker(workerId int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		time.Sleep(1 * time.Second)
		results <- fmt.Sprintf("THREAD_%d: Finished file%d.txt", workerId, job)
	}
}

func main() {
	startProcess := time.Now()

	concurrent := 1
	files := 60

	jobs := make(chan int, files)
	results := make(chan string)

	// Iniciando as goroutines
	for workerId := 0; workerId < concurrent; workerId++ {
		go worker(workerId+1, jobs, results)
	}

	// Enviando as tarefas para as goroutines
	for i := 0; i < files; i++ {
		jobs <- i
	}

	// Encerrando as tarefas
	close(jobs)

	for i := 0; i < files; i++ {
		fmt.Println(<-results)
	}

	close(results)

	endProccess := time.Now()
	fmt.Println("Total time in seconds:", endProccess.Sub(startProcess).Seconds())
}