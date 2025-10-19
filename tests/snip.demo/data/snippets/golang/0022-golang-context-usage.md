<!-- METADATA
{
  "title": "Golang Context Usage",
  "tags": [
    "go",
    "context"
  ],
  "language": "go"
}
-->

## Context Usage
Using context for cancellation and timeouts
```go
package main

import (
    "context"
    "fmt"
    "time"
)

func longTask(ctx context.Context) error {
    select {
    case <-time.After(2 * time.Second):
        fmt.Println("Task completed")
        return nil
    case <-ctx.Done():
        fmt.Printf("Task cancelled: %v\n", ctx.Err())
        return ctx.Err()
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    longTask(ctx)
}
```