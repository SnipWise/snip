## Variable Declaration
Different ways to declare and initialize variables
```go
package main

import "fmt"

func main() {
    var name string = "John"    // Explicit type
    var age = 30               // Type inference
    city := "New York"         // Short declaration
    var x, y int = 1, 2       // Multiple variables
    
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```