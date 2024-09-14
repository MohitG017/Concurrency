package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Job struct to represent a unit of work
type Job struct {
	ID     int
	Number int
}

// Result struct to represent the result of a job
type Result struct {
	Job    Job
	Square int
}

// Worker function to process jobs and send results
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		time.Sleep(time.Second) // Simulate time-consuming work
		output := Result{Job: job, Square: job.Number * job.Number}
		results <- output
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
	}
}

func main() {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	// Create 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Create 5 jobs
	for j := 1; j <= 5; j++ {
		job := Job{ID: j, Number: rand.Intn(10) + 1}
		jobs <- job
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Job ID: %d, Number: %d, Square: %d\n", result.Job.ID, result.Job.Number, result.Square)
	}
}
