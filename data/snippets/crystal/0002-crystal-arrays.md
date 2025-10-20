<!-- METADATA
{
  "title": "Crystal Arrays",
  "tags": [
    "crystal",
    "data-structures",
    "arrays"
  ],
  "language": "crystal"
}
-->

## Arrays
Array creation and manipulation
```crystal
# Array literal
numbers = [1, 2, 3, 4, 5]
names = ["Alice", "Bob", "Charlie"]

# Array with explicit type
values = Array(Int32).new
values << 10 << 20 << 30

# Array methods
numbers.size         # => 5
numbers.first        # => 1
numbers.last         # => 5
numbers[1]           # => 2
numbers.includes?(3) # => true

# Iteration
numbers.each { |n| puts n }
squares = numbers.map { |n| n * n }
```
