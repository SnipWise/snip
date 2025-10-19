<!-- METADATA
{
  "title": "Golang For Loops",
  "tags": [
    "go",
    "loops"
  ],
  "language": "go"
}
-->

## For Loops
Different types of for loops
```go
package main

import "fmt"

func main() {
    // Basic for loop
    for i := 0; i < 3; i++ {
        fmt.Printf("i = %d\n", i)
    }
    
    // Range over slice
    names := []string{"Alice", "Bob"}
    for index, name := range names {
        fmt.Printf("%d: %s\n", index, name)
    }
}
```