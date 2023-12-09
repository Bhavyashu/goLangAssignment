// controllers/arraySorting.go
package controllers

import (
   "goLangAssignment/models"
   "sort"
   "sync"
   "fmt"
   "time"
)

// Logic for sequential sorting
func SequentialSorting(payload *models.Payload) *models.Response {
    // Log information about the input array size
    fmt.Printf("SequentialSorting Sorting: Input array size %d\n", len(payload.ToSort))


   start := time.Now()
   for i := range payload.ToSort {
      sort.Ints(payload.ToSort[i])
   }
   end := time.Now()

   return &models.Response{
      TimeNS:       end.Sub(start).Nanoseconds(),
      SortedArrays: payload.ToSort,
   }
}

// Logic for concurrent sorting
func ConcurrentSorting(payload *models.Payload) *models.Response {
   fmt.Printf("Concurrent Sorting: Input array size %d\n", len(payload.ToSort))
   var wg sync.WaitGroup

   start := time.Now()
   for i := range payload.ToSort {
      wg.Add(1)
      go func(i int) {
         defer wg.Done()
         sort.Ints(payload.ToSort[i])
      }(i)
   }
   wg.Wait()
   end := time.Now()

   return &models.Response{
      TimeNS:       end.Sub(start).Nanoseconds(),
      SortedArrays: payload.ToSort,
   }
}
