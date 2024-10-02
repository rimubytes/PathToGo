// A worker pool is a pattern where you have a fixed number of workers (goroutines) processing tasks concurrently from a task queue

package main 

import (
	"fmt"
	"time"
)

// Worker function
func worker(id int, jobs <-chan int, results chan <- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d \n", id, j)
		time.Sleep(time.Second) // simulate time taken by the job
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2 
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= numJobs; a++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}