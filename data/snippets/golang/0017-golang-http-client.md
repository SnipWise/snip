<!-- METADATA
{
  "title": "Golang Http Client",
  "tags": [
    "go",
    "http",
    "networking"
  ],
  "language": "go"
}
-->

## HTTP Client
Making HTTP requests
```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, err := http.Get("https://httpbin.org/json")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading response: %v\n", err)
        return
    }
    
    fmt.Printf("Response: %s\n", string(body))
}
```