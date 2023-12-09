// arraySorting.go
package controllers

import (
	"fmt"
	"goLangAssignment/models"
	"sort"
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

// ConcurrentSorting performs concurrent sorting on the input arrays.
// It logs information about the input array size before sorting.
func ConcurrentSorting(payload *models.Payload) *models.Response {
	// Log information about the input array size
	fmt.Printf("Concurrent Sorting: Input array size %d\n", len(payload.ToSort))

	// Initialize a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Record the start time
	start := time.Now()

	// Perform concurrent sorting using goroutines
	for i := range payload.ToSort {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sort.Ints(payload.ToSort[i])
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Record the end time
	end := time.Now()

	// Return the response with sorted arrays and time taken
	return &models.Response{
		TimeNS:       end.Sub(start).Nanoseconds(),
		SortedArrays: payload.ToSort,
	}
}
