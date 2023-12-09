// models/response.go
package models

type Response struct {
   TimeNS       int64   `json:"time_ns"`
   SortedArrays [][]int `json:"sorted_arrays"`
}
