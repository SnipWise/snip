<!-- METADATA
{
  "title": "Golang Configuration Management",
  "tags": [
    "go",
    "io",
    "configuration"
  ],
  "language": "go"
}
-->

## Configuration Management
Managing application configuration
```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Config struct {
    Server struct {
        Host string `json:"host"`
        Port int    `json:"port"`
    } `json:"server"`
    Database struct {
        Host     string `json:"host"`
        Username string `json:"username"`
    } `json:"database"`
}

func loadConfig(filename string) (*Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var config Config
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    return &config, err
}

func main() {
    config, err := loadConfig("config.json")
    if err != nil {
        fmt.Printf("Error loading config: %v\n", err)
        return
    }
    
    fmt.Printf("Server: %s:%d\n", config.Server.Host, config.Server.Port)
}
```