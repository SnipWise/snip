<!-- METADATA
{
  "title": "Golang Channels",
  "tags": [
    "go",
    "concurrency"
  ],
  "language": "go"
}
-->

## Channels
Communication between goroutines
```go
package main

import "fmt"

func producer(ch chan<- int) {
    for i := 1; i <= 3; i++ {
        ch <- i
        fmt.Printf("Sent: %d\n", i)
    }
    close(ch)
}

func main() {
    ch := make(chan int)
    go producer(ch)
    
    for value := range ch {
        fmt.Printf("Received: %d\n", value)
    }
}
```