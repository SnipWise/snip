## Ownership and Borrowing
Understanding ownership rules
```rust
fn take_ownership(s: String) {
    println!("Took ownership of: {}", s);
}

fn borrow_string(s: &String) {
    println!("Borrowed: {}", s);
}

fn main() {
    let s1 = String::from("Hello");
    borrow_string(&s1); // Borrow
    println!("Still have: {}", s1);
    
    take_ownership(s1); // Move
    // println!("{}", s1); // Would cause compile error
}
```