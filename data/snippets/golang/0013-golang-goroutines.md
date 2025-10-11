<!-- METADATA
{
  "title": "Golang Goroutines",
  "tags": [
    "go",
    "concurrency"
  ],
  "language": "go"
}
-->

## Goroutines
Concurrent programming with goroutines
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(100 * time.Millisecond)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    wg.Wait()
}
```