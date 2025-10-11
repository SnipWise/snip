<!-- METADATA
{
  "title": "Golang File Io",
  "tags": [
    "go",
    "file-io",
    "io"
  ],
  "language": "go"
}
-->

## File I/O
Reading and writing files
```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    content := "Hello, World!\n"
    err := ioutil.WriteFile("test.txt", []byte(content), 0644)
    if err != nil {
        fmt.Printf("Error writing: %v\n", err)
        return
    }
    
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Printf("Error reading: %v\n", err)
        return
    }
    
    fmt.Printf("File content: %s", string(data))
    os.Remove("test.txt")
}
```