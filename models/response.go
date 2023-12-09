// response.go
package models

// Response represents the data structure for the response containing sorted arrays and time taken.
type Response struct {
	TimeNS       int64   `json:"time_ns"`       // TimeNS represents the time taken for sorting in nanoseconds.
	SortedArrays [][]int `json:"sorted_arrays"` // SortedArrays contains the sorted arrays.
}
