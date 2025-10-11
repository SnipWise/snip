<!-- METADATA
{
  "title": "Rustlang File Io",
  "tags": [
    "rust",
    "file-io",
    "io"
  ],
  "language": "rust"
}
-->

## File I/O
Reading and writing files
```rust
use std::fs;
use std::io::Write;

fn main() -> std::io::Result<()> {
    // Write to file
    let content = "Hello, Rust!";
    fs::write("test.txt", content)?;
    
    // Read from file
    let read_content = fs::read_to_string("test.txt")?;
    println!("File content: {}", read_content);
    
    // Append to file
    let mut file = std::fs::OpenOptions::new()
        .append(true)
        .open("test.txt")?;
    writeln!(file, "Appended line")?;
    
    // Clean up
    fs::remove_file("test.txt")?;
    Ok(())
}
```