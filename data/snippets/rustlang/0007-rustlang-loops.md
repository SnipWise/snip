<!-- METADATA
{
  "title": "Rustlang Loops",
  "tags": [
    "rust",
    "loops"
  ],
  "language": "rust"
}
-->

## Loops
Different types of loops
```rust
fn main() {
    // Loop with break
    let mut counter = 0;
    loop {
        if counter >= 3 {
            break;
        }
        println!("Loop: {}", counter);
        counter += 1;
    }
    
    // While loop
    let mut n = 3;
    while n > 0 {
        println!("While: {}", n);
        n -= 1;
    }
    
    // For loop
    for i in 0..3 {
        println!("For: {}", i);
    }
}
```