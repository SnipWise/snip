## Logging
Different logging approaches
```go
package main

import (
    "log"
    "log/slog"
    "os"
)

func main() {
    // Standard logging
    log.Println("Standard log message")
    log.Printf("User %s logged in", "Alice")
    
    // Custom logger
    logger := log.New(os.Stdout, "APP: ", log.LstdFlags)
    logger.Println("Custom logger message")
    
    // Structured logging (Go 1.21+)
    slog.Info("User action", 
        "user_id", 123, 
        "action", "login")
    
    slog.Error("Database error", 
        "error", "connection timeout",
        "retry_count", 3)
}
```