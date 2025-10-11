## Error Handling
Proper error handling with ? operator
```rust
use std::fs;
use std::io;

fn read_file_content(path: &str) -> Result<String, io::Error> {
    let content = fs::read_to_string(path)?;
    Ok(content.trim().to_string())
}

fn main() {
    match read_file_content("test.txt") {
        Ok(content) => println!("File content: {}", content),
        Err(e) => println!("Error reading file: {}", e),
    }
}
```