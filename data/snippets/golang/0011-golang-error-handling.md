<!-- METADATA
{
  "title": "Golang Error Handling",
  "tags": [
    "go",
    "error-handling"
  ],
  "language": "go"
}
-->

## Error Handling
Proper error handling patterns
```go
package main

import (
    "errors"
    "fmt"
)

var ErrDivisionByZero = errors.New("division by zero")

func safeDivide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, ErrDivisionByZero
    }
    return a / b, nil
}

func main() {
    result, err := safeDivide(10, 0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Result: %f\n", result)
    }
}
```