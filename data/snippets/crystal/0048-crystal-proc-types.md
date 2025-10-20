<!-- METADATA
{
  "title": "Crystal Proc Types",
  "tags": [
    "crystal",
    "types",
    "procs"
  ],
  "language": "crystal"
}
-->

## Proc Types
Type-safe function pointers
```crystal
# Proc type definition
calculator : Proc(Int32, Int32, Int32)
calculator = ->(a : Int32, b : Int32) { a + b }

result = calculator.call(5, 3)  # => 8

# Proc with no arguments
greeter : Proc(String)
greeter = ->{ "Hello!" }
puts greeter.call

# Proc with no return
action : Proc(String, Nil)
action = ->(msg : String) { puts msg; nil }
action.call("Hello")

# Storing procs in collections
operations = [] of Proc(Int32, Int32, Int32)
operations << ->(a : Int32, b : Int32) { a + b }
operations << ->(a : Int32, b : Int32) { a * b }

operations.each do |op|
  puts op.call(10, 5)
end

# Proc as class member
class Calculator
  property operation : Proc(Int32, Int32, Int32)

  def initialize(@operation)
  end

  def compute(a, b)
    @operation.call(a, b)
  end
end

calc = Calculator.new(->(x : Int32, y : Int32) { x + y })
puts calc.compute(7, 3)

# Partial application
def partial(f : Proc(Int32, Int32, Int32), x : Int32)
  ->(y : Int32) { f.call(x, y) }
end

add = ->(a : Int32, b : Int32) { a + b }
add_5 = partial(add, 5)
puts add_5.call(3)  # => 8
```
