<!-- METADATA
{
  "title": "Crystal Methods",
  "tags": [
    "crystal",
    "basics",
    "functions"
  ],
  "language": "crystal"
}
-->

## Methods
Method definition and usage
```crystal
# Basic method
def greet(name)
  "Hello, #{name}!"
end

# Method with type restrictions
def add(a : Int32, b : Int32) : Int32
  a + b
end

# Default parameters
def repeat(text, times = 3)
  text * times
end

# Named arguments
def create_user(name : String, age : Int32, role = "user")
  {name: name, age: age, role: role}
end

user = create_user(name: "Alice", age: 30)

# Block parameters
def execute
  puts "Before"
  yield
  puts "After"
end

execute { puts "Inside block" }
```
