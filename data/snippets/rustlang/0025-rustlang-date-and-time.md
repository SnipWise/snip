## Date and Time
Working with chrono for date/time operations
```rust
use chrono::{DateTime, Utc, Local, Duration};

fn main() {
    let now: DateTime<Utc> = Utc::now();
    let local_now: DateTime<Local> = Local::now();
    
    println!("UTC now: {}", now);
    println!("Local now: {}", local_now);
    println!("Formatted: {}", now.format("%Y-%m-%d %H:%M:%S"));
    
    // Date arithmetic
    let tomorrow = now + Duration::days(1);
    let one_hour_ago = now - Duration::hours(1);
    
    println!("Tomorrow: {}", tomorrow);
    println!("One hour ago: {}", one_hour_ago);
    
    // Parse date string
    let date_str = "2023-12-25T10:30:00Z";
    match DateTime::parse_from_rfc3339(date_str) {
        Ok(parsed) => println!("Parsed: {}", parsed),
        Err(e) => println!("Parse error: {}", e),
    }
}
```