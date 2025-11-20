package main

import (
	"fmt"
)

type FetchResult struct {
	URL        string
	StatusCode int
	Size       int
	Error      error
}

func worker(id int, jobs <-chan string, results chan<- FetchResult) {
	defer wg.Done()
	// TODO: fetch the URL
	// TODO: send result struct to results channel
	// hint: use resp, err := http.Get(url)
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://uottawa.ca",
		"https://github.com",
		"https://httpbin.org/get",
	}

	numWorkers := 3

	jobs := make(chan string, len(urls))
	results := make(chan FetchResult, len(urls))

	//start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// send jobs
	for j := 1; j <= len(urls); j++ {
		jobs <- urls[j-1]
	}
	close(jobs)

	// collect results
	for i := 1; i <= len(urls); i++ {
		fmt.Println("Result:", <-results)
	}

	fmt.Println("\n Scraping complete!")
}
