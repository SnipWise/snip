## Structs
Defining and using structs
```rust
struct Person {
    name: String,
    age: u32,
}

impl Person {
    fn new(name: String, age: u32) -> Person {
        Person { name, age }
    }
    
    fn greet(&self) -> String {
        format!("Hello, I'm {}", self.name)
    }
}

fn main() {
    let person = Person::new("Alice".to_string(), 30);
    println!("Person: {} ({})", person.name, person.age);
    println!("{}", person.greet());
}
```