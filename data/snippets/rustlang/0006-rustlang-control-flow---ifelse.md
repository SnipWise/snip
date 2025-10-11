<!-- METADATA
{
  "title": "Rustlang Control Flow   Ifelse",
  "tags": [
    "rust",
    "control-flow"
  ],
  "language": "rust"
}
-->

## Control Flow - If/Else
Conditional statements
```rust
fn main() {
    let age = 25;
    
    if age >= 18 {
        println!("Adult");
    } else {
        println!("Minor");
    }
    
    // If as expression
    let status = if age >= 18 { "adult" } else { "minor" };
    println!("Status: {}", status);
    
    let score = 85;
    let grade = if score >= 90 { "A" } else if score >= 80 { "B" } else { "C" };
    println!("Grade: {}", grade);
}
```