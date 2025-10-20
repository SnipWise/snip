<!-- METADATA
{
  "title": "Crystal Hashes",
  "tags": [
    "crystal",
    "data-structures",
    "hashes"
  ],
  "language": "crystal"
}
-->

## Hashes
Hash creation and manipulation
```crystal
# Hash literal
person = {
  "name" => "Alice",
  "age"  => 30,
  "city" => "Paris"
}

# Symbol keys
config = {
  host: "localhost",
  port: 3000,
  debug: true
}

# Access
person["name"]  # => "Alice"
config[:port]   # => 3000

# Methods
person.has_key?("name")  # => true
person.keys              # => ["name", "age", "city"]
person.values            # => ["Alice", 30, "Paris"]
```
