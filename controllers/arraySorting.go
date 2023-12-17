// arraySorting.go
package controllers

import (
	"fmt"
	"app/models"
	"sort"
	"runtime"
	"sync"
	"time"
)

// SequentialSorting performs sequential sorting on the input arrays.
// It logs information about the input array size before sorting.
func SequentialSorting(payload *models.Payload) *models.Response {
	// Log information about the input array size
	fmt.Printf("Sequential Sorting: Input array size %d\n", len(payload.ToSort))

	// Record the start time
	start := time.Now()

	// Perform sequential sorting
	for i := range payload.ToSort {
		sort.Ints(payload.ToSort[i])
	}

	// Record the end time
	end := time.Now()

	// Return the response with sorted arrays and time taken
	return &models.Response{
		TimeNS:       end.Sub(start).Nanoseconds(),
		SortedArrays: payload.ToSort,
	}
}

//Batch Processing using conncurrent workers
func ConcurrentSorting(payload *models.Payload) *models.Response {
	// Log information about the input array size
	fmt.Printf("Concurrent Sorting: Input array size %d\n", len(payload.ToSort))

	// Record the start time
	start := time.Now()

	// Determine the number of workers (goroutines)
	numWorkers := runtime.NumCPU()

	// Create a channel to communicate the batches of subarrays to be sorted
	batchChannel := make(chan [][]int, numWorkers)

	// Initialize a WaitGroup to wait for all batches to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func() {
			for batch := range batchChannel {
				for _, subarray := range batch {
					sort.Ints(subarray)
				}
				wg.Done()
			}
		}()
	}

	// Queue tasks in batches
	batchSize := len(payload.ToSort) / numWorkers
	for i := 0; i < len(payload.ToSort); i += batchSize {
		wg.Add(1)
		end := i + batchSize
		if end > len(payload.ToSort) {
			end = len(payload.ToSort)
		}
		batchChannel <- payload.ToSort[i:end]
	}

	// Close the channel to signal workers to exit
	close(batchChannel)

	// Wait for all batches to finish
	wg.Wait()

	// Record the end time
	end := time.Now()

	// Return the response with sorted arrays and time taken
	return &models.Response{
		TimeNS:       end.Sub(start).Nanoseconds(),
		SortedArrays: payload.ToSort,
	}
}

