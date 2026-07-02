/* Mode 1: Sequential   → one request at a time → pure latency
Mode 2: Concurrent   → many at once → throughput + connection reuse
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	url           = "http://localhost:8080/users/1"
	totalRequests = 1000
	concurrency   = 1000 //for concurrent mode
)

func main() {
	fmt.Println("====HTTP/1.1 Benchmark====")
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Total Requests: %d\n", totalRequests)

	runSequentialMode()
	fmt.Println("====================================")
	runConcurrentMode()

}

func runSequentialMode() {
	fmt.Println("Running in Sequential Mode... [1 at a time → pure latency]")
	// Implement sequential HTTP requests here

	latencies := make([]time.Duration, 0, totalRequests)
	start := time.Now()

	client := &http.Client{}
	for i := 0; i < totalRequests; i++ {
		reqStart := time.Now()
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		latencies = append(latencies, time.Since(reqStart))

	}

	total := time.Since(start)
	report(latencies, total)
}

func runConcurrentMode() {

	fmt.Printf("====Concurrent (%d at a time)=====\n", concurrency)

	latencies := make([]time.Duration, totalRequests)
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	client := http.Client{}
	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(idx int) {
			defer wg.Done()
			defer func() { <-sem }()

			reqStart := time.Now()
			resp, err := client.Get(url)
			if err != nil {
				fmt.Println("Error making request:", err)
				return

			}

			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()

			latencies[idx] = time.Since(reqStart)
		}(i)
	}

	wg.Wait()
	total := time.Since(start)
	report(latencies, total)
}

func report(latencies []time.Duration, total time.Duration) {

	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	var sum time.Duration
	for _, l := range latencies {
		sum += l
	}

	avg := sum / time.Duration(len(latencies))

	p50 := latencies[len(latencies)*50/100]
	p90 := latencies[len(latencies)*90/100]
	p99 := latencies[len(latencies)*99/100]
	min := latencies[0]
	max := latencies[len(latencies)-1]

	throughput := float64(len(latencies)) / total.Seconds()

	fmt.Println("Total Time:", total)
	fmt.Println("Throughput (requests/sec):", throughput)
	fmt.Println("Average Latency:", avg)
	fmt.Println("P50 Latency:", p50)
	fmt.Println("P90 Latency:", p90)
	fmt.Println("P99 Latency:", p99)
	fmt.Println("Min Latency:", min)
	fmt.Println("Max Latency:", max)
}
