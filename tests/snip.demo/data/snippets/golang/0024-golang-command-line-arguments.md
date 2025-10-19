<!-- METADATA
{
  "title": "Golang Command Line Arguments",
  "tags": [
    "go",
    "cli"
  ],
  "language": "go"
}
-->

## Command Line Arguments
Parsing command line arguments
```go
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    var name = flag.String("name", "World", "Name to greet")
    var age = flag.Int("age", 0, "Age of person")
    var verbose = flag.Bool("verbose", false, "Verbose output")
    
    flag.Parse()
    
    if *verbose {
        fmt.Printf("Name: %s, Age: %d\n", *name, *age)
    }
    
    fmt.Printf("Hello, %s!\n", *name)
    
    // Remaining arguments
    args := flag.Args()
    if len(args) > 0 {
        fmt.Printf("Extra args: %v\n", args)
    }
}
```