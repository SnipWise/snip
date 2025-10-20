<!-- METADATA
{
  "title": "Crystal String Operations",
  "tags": [
    "crystal",
    "strings",
    "basics"
  ],
  "language": "crystal"
}
-->

## String Operations
Common string manipulations
```crystal
# Interpolation
name = "Alice"
greeting = "Hello, #{name}!"

# Concatenation
result = "Hello" + " " + "World"

# Multiline strings
text = "Line 1
Line 2
Line 3"

# Methods
str = "  hello world  "
puts str.upcase          # => "  HELLO WORLD  "
puts str.downcase        # => "  hello world  "
puts str.strip           # => "hello world"
puts str.size            # => 16

# Split and join
words = "a,b,c".split(",")
joined = words.join("-")  # => "a-b-c"

# Contains
puts "hello".includes?("ll")  # => true

# Replace
result = "hello".gsub("l", "L")  # => "heLLo"

# Substring
puts "hello"[1..3]  # => "ell"
```
