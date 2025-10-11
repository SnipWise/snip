<!-- METADATA
{
  "title": "Rustlang Regular Expressions",
  "tags": [
    "rust",
    "io",
    "regex"
  ],
  "language": "rust"
}
-->

## Regular Expressions
Pattern matching with regex
```rust
use regex::Regex;

fn main() {
    let email_regex = Regex::new(r"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$").unwrap();
    
    let emails = vec!["user@example.com", "invalid.email", "test@domain.co.uk"];
    
    for email in emails {
        if email_regex.is_match(email) {
            println!("{}: Valid", email);
        } else {
            println!("{}: Invalid", email);
        }
    }
    
    // Find all matches
    let text = "Contact us at info@company.com or support@company.com";
    let matches: Vec<&str> = email_regex.find_iter(text).map(|m| m.as_str()).collect();
    println!("Found emails: {:?}", matches);
    
    // Replace matches
    let replaced = email_regex.replace_all(text, "[EMAIL]");
    println!("Replaced: {}", replaced);
}
```