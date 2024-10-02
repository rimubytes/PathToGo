// A rate limiter controls how often a function or process can execute

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that limits the rate to 1 request per second
	rateLimiter := time.NewTicker(time.Second) // The time.NewTicker creates a ticker that ticks every second, controlling the rate of request handling
	defer rateLimiter.Stop()

	// Simulate 5 requests
	for i := 1; i <= 5; i++ {
		<-rateLimiter.c // Block until the ticker ticks
		go handleRequest(i) // Handle the request in a goroutine
	}

	// Block to ensure all goroutines finish
	time.Sleep(6 * time.Second)
}

func handleRequest(id int) {
	fmt.Printf("Handling request %d at %s\n", id, time.Now().Format(time.RFC3339))
}
