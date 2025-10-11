<!-- METADATA
{
  "title": "Golang Json Handling",
  "tags": [
    "go",
    "json"
  ],
  "language": "go"
}
-->

## JSON Handling
Marshal and unmarshal JSON data
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

func main() {
    person := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
    
    jsonData, _ := json.Marshal(person)
    fmt.Printf("JSON: %s\n", string(jsonData))
    
    var newPerson Person
    json.Unmarshal(jsonData, &newPerson)
    fmt.Printf("Unmarshaled: %+v\n", newPerson)
}
```