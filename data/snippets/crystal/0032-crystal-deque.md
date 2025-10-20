<!-- METADATA
{
  "title": "Crystal Deque",
  "tags": [
    "crystal",
    "data-structures",
    "deque"
  ],
  "language": "crystal"
}
-->

## Deque
Double-ended queue for efficient push/pop
```crystal
require "deque"

# Create deque
deque = Deque(Int32).new
deque = Deque{1, 2, 3, 4, 5}

# Add to front
deque.unshift(0)  # => Deque{0, 1, 2, 3, 4, 5}

# Add to back
deque.push(6)     # => Deque{0, 1, 2, 3, 4, 5, 6}
deque << 7

# Remove from front
first = deque.shift  # => 0

# Remove from back
last = deque.pop     # => 7

# Access elements
puts deque[0]        # First element
puts deque[-1]       # Last element

# Peek without removing
puts deque.first?
puts deque.last?

# Other operations
puts deque.size
puts deque.empty?
deque.clear
```
