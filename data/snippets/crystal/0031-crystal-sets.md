<!-- METADATA
{
  "title": "Crystal Sets",
  "tags": [
    "crystal",
    "data-structures",
    "sets"
  ],
  "language": "crystal"
}
-->

## Sets
Unordered collection of unique elements
```crystal
require "set"

# Create set
numbers = Set{1, 2, 3, 4, 5}
words = Set(String).new

# Add elements
numbers << 6
numbers.add(7)

# Check membership
puts numbers.includes?(3)  # => true

# Set operations
a = Set{1, 2, 3, 4}
b = Set{3, 4, 5, 6}

union = a | b           # => Set{1, 2, 3, 4, 5, 6}
intersection = a & b    # => Set{3, 4}
difference = a - b      # => Set{1, 2}

# Convert from array
array_set = [1, 2, 2, 3, 3].to_set  # => Set{1, 2, 3}

# Iterate
numbers.each { |n| puts n }

# Size and empty check
puts numbers.size
puts numbers.empty?  # => false
```
