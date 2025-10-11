<!-- METADATA
{
  "title": "Golang String Manipulation",
  "tags": [
    "go",
    "io",
    "strings"
  ],
  "language": "go"
}
-->

## String Manipulation
Common string operations
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    text := "  Hello, World!  "
    
    fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(text))
    fmt.Printf("Upper: %s\n", strings.ToUpper(text))
    fmt.Printf("Contains 'World': %t\n", strings.Contains(text, "World"))
    
    words := strings.Fields(strings.TrimSpace(text))
    fmt.Printf("Words: %v\n", words)
    
    // String conversion
    num, _ := strconv.Atoi("123")
    fmt.Printf("String to int: %d\n", num)
}
```