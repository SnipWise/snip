<!-- METADATA
{
  "title": "Golang Constants",
  "tags": [
    "go",
    "constants"
  ],
  "language": "go"
}
-->

## Constants
Declaring and using constants
```go
package main

import "fmt"

const (
    Pi       = 3.14159
    MaxUsers = 100
)

func main() {
    fmt.Printf("Pi: %f, Max Users: %d\n", Pi, MaxUsers)
}
```