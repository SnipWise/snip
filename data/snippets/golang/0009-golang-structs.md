## Structs
Defining and using structs
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.Name)
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    fmt.Printf("Person: %+v\n", person)
    fmt.Println(person.Greet())
}
```