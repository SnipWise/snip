<!-- METADATA
{
  "title": "Rustlang Generics",
  "tags": [
    "rust",
    "generics"
  ],
  "language": "rust"
}
-->

## Generics
Using generic types and functions
```rust
fn find_largest<T: PartialOrd + Copy>(list: &[T]) -> T {
    let mut largest = list[0];
    for &item in list {
        if item > largest {
            largest = item;
        }
    }
    largest
}

struct Point<T> {
    x: T,
    y: T,
}

fn main() {
    let numbers = vec![34, 50, 25, 100, 65];
    let largest = find_largest(&numbers);
    println!("Largest number: {}", largest);
    
    let point = Point { x: 5, y: 10 };
    println!("Point: ({}, {})", point.x, point.y);
}
```