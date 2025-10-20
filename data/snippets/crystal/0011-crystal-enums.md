<!-- METADATA
{
  "title": "Crystal Enums",
  "tags": [
    "crystal",
    "enums",
    "types"
  ],
  "language": "crystal"
}
-->

## Enums
Enumeration types
```crystal
# Basic enum
enum Color
  Red
  Green
  Blue
end

# Enum with values
enum HttpStatus
  OK           = 200
  NotFound     = 404
  ServerError  = 500
end

# Flags enum (bitwise operations)
@[Flags]
enum Permission
  Read
  Write
  Execute
end

# Usage
color = Color::Red
status = HttpStatus::OK

puts status.value  # => 200

perms = Permission::Read | Permission::Write
puts perms.includes?(Permission::Read)   # => true
puts perms.includes?(Permission::Execute) # => false
```
