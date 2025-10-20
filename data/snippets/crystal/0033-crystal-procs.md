<!-- METADATA
{
  "title": "Crystal Procs and Blocks",
  "tags": [
    "crystal",
    "functional",
    "procs"
  ],
  "language": "crystal"
}
-->

## Procs and Blocks
First-class functions and closures
```crystal
# Proc literal
add = ->(x : Int32, y : Int32) { x + y }
puts add.call(3, 4)  # => 7

# Proc with multiple lines
greet = ->(name : String) {
  message = "Hello, #{name}!"
  puts message
  message
}

# Proc.new
multiply = Proc(Int32, Int32, Int32).new do |x, y|
  x * y
end

# Closure
def make_counter
  count = 0
  ->{
    count += 1
    count
  }
end

counter = make_counter
puts counter.call  # => 1
puts counter.call  # => 2

# Passing procs
def apply(x, y, operation : Proc(Int32, Int32, Int32))
  operation.call(x, y)
end

result = apply(5, 3, add)
puts result  # => 8

# Block to proc
def execute(&block : Int32 -> Int32)
  block.call(10)
end

execute { |x| x * 2 }  # => 20
```
