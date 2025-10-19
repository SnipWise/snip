<!-- METADATA
{
  "title": "Golang Interfaces",
  "tags": [
    "go",
    "interfaces"
  ],
  "language": "go"
}
-->

## Interfaces
Defining and implementing interfaces
```go
package main

import "fmt"

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    var s Shape = Rectangle{Width: 5, Height: 3}
    fmt.Printf("Area: %.2f\n", s.Area())
}
```