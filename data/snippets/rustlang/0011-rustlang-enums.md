<!-- METADATA
{
  "title": "Rustlang Enums",
  "tags": [
    "rust",
    "enums"
  ],
  "language": "rust"
}
-->

## Enums
Defining and using enums
```rust
enum Color {
    Red,
    Green,
    Blue,
    RGB(u8, u8, u8),
}

fn main() {
    let red = Color::Red;
    let custom = Color::RGB(255, 128, 0);
    
    match red {
        Color::Red => println!("It's red!"),
        Color::Green => println!("It's green!"),
        Color::Blue => println!("It's blue!"),
        Color::RGB(r, g, b) => println!("RGB: {}, {}, {}", r, g, b),
    }
}
```