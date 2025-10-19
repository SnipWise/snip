<!-- METADATA
{
  "title": "Golang Arrays And Slices",
  "tags": [
    "go",
    "arrays",
    "slices"
  ],
  "language": "go"
}
-->

## Arrays and Slices
Working with arrays and slices
```go
package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    slice := []string{"apple", "banana", "cherry"}
    slice = append(slice, "date")
    
    fmt.Printf("Array: %v\n", arr)
    fmt.Printf("Slice: %v, Len: %d\n", slice, len(slice))
}
```