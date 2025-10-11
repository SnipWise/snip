<!-- METADATA
{
  "title": "Rustlang Hash Maps",
  "tags": [
    "rust",
    "maps"
  ],
  "language": "rust"
}
-->

## Hash Maps
Creating and manipulating hash maps
```rust
use std::collections::HashMap;

fn main() {
    let mut map = HashMap::new();
    map.insert("apple", 5);
    map.insert("banana", 3);
    
    // Check if key exists
    match map.get("apple") {
        Some(value) => println!("Apple: {}", value),
        None => println!("Apple not found"),
    }
    
    for (key, value) in &map {
        println!("{}: {}", key, value);
    }
}
```