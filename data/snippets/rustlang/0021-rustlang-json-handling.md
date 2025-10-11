## JSON Handling
Serialize and deserialize JSON with serde
```rust
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
struct Person {
    name: String,
    age: u32,
    email: Option<String>,
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let person = Person {
        name: "Alice".to_string(),
        age: 30,
        email: Some("alice@example.com".to_string()),
    };
    
    // Serialize to JSON
    let json = serde_json::to_string(&person)?;
    println!("JSON: {}", json);
    
    // Deserialize from JSON
    let parsed: Person = serde_json::from_str(&json)?;
    println!("Parsed: {:?}", parsed);
    
    Ok(())
}
```