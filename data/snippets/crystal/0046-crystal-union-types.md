<!-- METADATA
{
  "title": "Crystal Union Types",
  "tags": [
    "crystal",
    "types",
    "unions"
  ],
  "language": "crystal"
}
-->

## Union Types
Multiple possible types
```crystal
# Union type variable
value : Int32 | String = 42
value = "hello"  # Valid

# Method with union return type
def parse_value(input : String) : Int32 | String | Nil
  return nil if input.empty?
  input.to_i? || input
end

result = parse_value("123")  # => 123
result = parse_value("abc")  # => "abc"
result = parse_value("")     # => nil

# Type checking
value = parse_value("42")

if value.is_a?(Int32)
  puts "Integer: #{value + 10}"
elsif value.is_a?(String)
  puts "String: #{value.upcase}"
end

# Case with union types
case value
when Int32
  puts "Got integer: #{value}"
when String
  puts "Got string: #{value}"
when Nil
  puts "Got nil"
end

# Union in class
class Response
  property data : String | Array(String) | Nil

  def initialize(@data)
  end
end

# Narrowing types
def process(val : Int32 | String)
  if val.is_a?(String)
    val.upcase  # Compiler knows val is String
  else
    val * 2     # Compiler knows val is Int32
  end
end
```
