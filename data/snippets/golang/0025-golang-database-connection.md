## Database Connection
Basic database operations with PostgreSQL
```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
)

type User struct {
    ID    int
    Name  string
    Email string
}

func main() {
    db, err := sql.Open("postgres", "host=localhost dbname=test sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    var user User
    err = db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", 1).
        Scan(&user.ID, &user.Name, &user.Email)
    
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("User: %+v\n", user)
    }
}
```