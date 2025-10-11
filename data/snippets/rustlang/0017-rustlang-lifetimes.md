<!-- METADATA
{
  "title": "Rustlang Lifetimes",
  "tags": [
    "rust",
    "control-flow",
    "time",
    "lifetimes"
  ],
  "language": "rust"
}
-->

## Lifetimes
Working with lifetime annotations
```rust
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

struct ImportantExcerpt<'a> {
    part: &'a str,
}

fn main() {
    let string1 = String::from("long string");
    let string2 = "short";
    
    let result = longest(&string1, string2);
    println!("Longest: {}", result);
    
    let excerpt = ImportantExcerpt { part: &string1 };
    println!("Excerpt: {}", excerpt.part);
}
```