<!-- METADATA
{
  "title": "Golang Middleware Pattern",
  "tags": [
    "go",
    "patterns"
  ],
  "language": "go"
}
-->

## Middleware Pattern
HTTP middleware implementation
```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)
        next(w, r)
        log.Printf("Completed in %v", time.Since(start))
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", loggingMiddleware(homeHandler))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```