<!-- METADATA
{
  "title": "Golang Maps",
  "tags": [
    "go",
    "maps"
  ],
  "language": "go"
}
-->

## Maps
Creating and manipulating maps
```go
package main

import "fmt"

func main() {
    fruits := map[string]int{"orange": 10, "grape": 15}
    fruits["apple"] = 5
    
    value, exists := fruits["apple"]
    fmt.Printf("Apple exists: %t, value: %d\n", exists, value)
    
    for key, value := range fruits {
        fmt.Printf("%s: %d\n", key, value)
    }
}
```