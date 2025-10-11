## Regular Expressions
Pattern matching with regex
```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    
    emails := []string{"user@example.com", "invalid.email", "test@domain.co.uk"}
    
    for _, email := range emails {
        isValid := emailRegex.MatchString(email)
        fmt.Printf("%s: %t\n", email, isValid)
    }
    
    text := "Contact us at info@company.com"
    matches := emailRegex.FindAllString(text, -1)
    fmt.Printf("Found emails: %v\n", matches)
}
```