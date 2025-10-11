## Pointers
Working with pointers
```go
package main

import "fmt"

func modifyValue(x *int) {
    *x = 100
}

func main() {
    value := 42
    ptr := &value
    
    fmt.Printf("Value: %d, Address: %p\n", value, ptr)
    fmt.Printf("Value via pointer: %d\n", *ptr)
    
    modifyValue(&value)
    fmt.Printf("Modified value: %d\n", value)
}
```