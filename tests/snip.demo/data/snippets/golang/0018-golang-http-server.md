<!-- METADATA
{
  "title": "Golang Http Server",
  "tags": [
    "go",
    "http",
    "networking"
  ],
  "language": "go"
}
-->

## HTTP Server
Creating a simple HTTP server
```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"message": "Hello, JSON!"}`)
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/json", jsonHandler)
    
    fmt.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```