## Match Statements
Pattern matching with match
```rust
fn main() {
    let number = 3;
    
    match number {
        1 => println!("One"),
        2 | 3 => println!("Two or Three"),
        4..=10 => println!("Four to Ten"),
        _ => println!("Something else"),
    }
    
    let option_value = Some(5);
    match option_value {
        Some(x) => println!("Got value: {}", x),
        None => println!("No value"),
    }
}
```