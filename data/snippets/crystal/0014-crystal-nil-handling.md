<!-- METADATA
{
  "title": "Crystal Nil Handling",
  "tags": [
    "crystal",
    "nil",
    "types"
  ],
  "language": "crystal"
}
-->

## Nil Handling
Working with nullable types
```crystal
# Nullable type
name : String? = "Alice"
empty : String? = nil

# Nil check
if name
  puts name.upcase  # Safe to use
end

# Safe navigation
puts name.try(&.upcase)  # => "ALICE"
puts empty.try(&.upcase) # => nil

# Not-nil assertion (!)
def get_value : String?
  "hello"
end

value = get_value
puts value.not_nil!.upcase  # Raises if nil

# Default value with ||
result = empty || "default"
puts result  # => "default"

# Compact mapping
values = [1, nil, 3, nil, 5]
non_nil = values.compact  # => [1, 3, 5]
```
