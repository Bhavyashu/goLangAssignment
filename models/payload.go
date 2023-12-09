// payload.go
package models

// Payload represents the data structure for the input payload containing arrays to be sorted.
type Payload struct {
	ToSort [][]int `json:"to_sort"`
}
