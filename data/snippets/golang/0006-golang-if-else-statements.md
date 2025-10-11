## If-Else Statements
Conditional statements
```go
package main

import "fmt"

func main() {
    age := 25
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }
    
    // If with short variable declaration
    if num := 42; num > 40 {
        fmt.Printf("Number %d > 40\n", num)
    }
}
```