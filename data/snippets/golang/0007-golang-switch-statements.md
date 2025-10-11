<!-- METADATA
{
  "title": "Golang Switch Statements",
  "tags": [
    "go",
    "control-flow"
  ],
  "language": "go"
}
-->

## Switch Statements
Switch case statements
```go
package main

import "fmt"

func main() {
    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Other")
    }
}
```