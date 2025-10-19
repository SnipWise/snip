<!-- METADATA
{
  "title": "Golang Functions",
  "tags": [
    "go",
    "functions",
    "io"
  ],
  "language": "go"
}
-->

## Functions
Function declaration and return values
```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func main() {
    sum := add(5, 3)
    result, err := divide(10, 2)
    fmt.Printf("Sum: %d, Division: %f, Error: %v\n", sum, result, err)
}
```