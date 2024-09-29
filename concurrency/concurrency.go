package main

import (
	"fmt"
	"time"
)

// Producer generates numbers and sends them to the `jobs` channel
func Producer(jobs chan<- int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("Producer: produced %d\n", i)
		jobs <- i                          // Send a job to the jobs channel
		time.Sleep(time.Millisecond * 500) // Simulate production delay
	}
	close(jobs) // Close the jobs channel to indicate no more jobs
}

// Consumer reads from the `jobs` channel and processes the data
func Consumer(jobs <-chan int, done chan<- bool) {
	for job := range jobs {
		fmt.Printf("Consumer: consumed %d\n", job)
		time.Sleep(time.Millisecond * 1000) // Simulate processing delay
	}
	done <- true // Notify that consuming is done
}

func main() {
	// Create a channel for jobs
	jobs := make(chan int, 10) // Buffered channel to handle jobs
	done := make(chan bool)    // Done channel to notify when consuming is finished

	// Start the Producer and Consumer goroutines
	go Producer(jobs, 5) // Produces 5 jobs
	go Consumer(jobs, done)

	// Wait for the consumer to finish processing
	<-done
	fmt.Println("All jobs processed. Exiting.")
}
