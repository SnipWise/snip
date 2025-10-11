<!-- METADATA
{
  "title": "Rustlang Option And Result",
  "tags": [
    "rust",
    "io",
    "option",
    "result"
  ],
  "language": "rust"
}
-->

## Option and Result
Working with Option and Result types
```rust
fn find_user(id: u32) -> Option<String> {
    if id == 1 {
        Some("Alice".to_string())
    } else {
        None
    }
}

fn parse_number(s: &str) -> Result<i32, std::num::ParseIntError> {
    s.parse()
}

fn main() {
    match find_user(1) {
        Some(name) => println!("Found user: {}", name),
        None => println!("User not found"),
    }
    
    match parse_number("42") {
        Ok(n) => println!("Parsed: {}", n),
        Err(e) => println!("Parse error: {}", e),
    }
}
```