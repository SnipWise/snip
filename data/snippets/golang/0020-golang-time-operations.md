<!-- METADATA
{
  "title": "Golang Time Operations",
  "tags": [
    "go",
    "io",
    "time"
  ],
  "language": "go"
}
-->

## Time Operations
Working with time and date
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Printf("Current time: %v\n", now)
    fmt.Printf("Formatted: %s\n", now.Format("2006-01-02 15:04:05"))
    
    tomorrow := now.Add(24 * time.Hour)
    fmt.Printf("Tomorrow: %v\n", tomorrow)
    
    // Parse time string
    timeStr := "2023-12-25 10:30:00"
    parsed, _ := time.Parse("2006-01-02 15:04:05", timeStr)
    fmt.Printf("Parsed: %v\n", parsed)
}
```