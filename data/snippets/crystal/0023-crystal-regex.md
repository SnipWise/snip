<!-- METADATA
{
  "title": "Crystal Regular Expressions",
  "tags": [
    "crystal",
    "regex",
    "strings"
  ],
  "language": "crystal"
}
-->

## Regular Expressions
Pattern matching with regex
```crystal
# Regex literal
pattern = /hello/i  # Case insensitive

# Match
text = "Hello, World!"
if text =~ /hello/i
  puts "Matched!"
end

# Capture groups
email = "user@example.com"
if email =~ /(\w+)@(\w+)\.(\w+)/
  puts $1  # => "user"
  puts $2  # => "example"
  puts $3  # => "com"
end

# Match method
match = "abc123".match(/([a-z]+)(\d+)/)
if match
  puts match[1]  # => "abc"
  puts match[2]  # => "123"
end

# Replace
result = "hello world".gsub(/o/, "0")
puts result  # => "hell0 w0rld"

# Split
words = "a,b,c".split(/,/)
puts words  # => ["a", "b", "c"]

# Scan
"a1b2c3".scan(/[a-z](\d)/) do |match|
  puts match[1]
end
```
