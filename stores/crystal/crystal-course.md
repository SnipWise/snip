# Complete Crystal Programming Course

## Table of Contents

1. [Introduction to Crystal](#introduction-to-crystal)
2. [Getting Started](#getting-started)
3. [Basic Syntax](#basic-syntax)
4. [Variables and Data Types](#variables-and-data-types)
5. [Control Flow](#control-flow)
6. [Methods](#methods)
7. [Classes and Objects](#classes-and-objects)
8. [Modules](#modules)
9. [Structs](#structs)
10. [Enums](#enums)
11. [Arrays and Tuples](#arrays-and-tuples)
12. [Hashes](#hashes)
13. [Ranges](#ranges)
14. [Blocks and Procs](#blocks-and-procs)
15. [Type System](#type-system)
16. [Generics](#generics)
17. [Exception Handling](#exception-handling)
18. [File I/O](#file-io)
19. [Concurrency](#concurrency)
20. [Macros](#macros)
21. [C Bindings](#c-bindings)
22. [Testing](#testing)
23. [Project Structure](#project-structure)
24. [Advanced Topics](#advanced-topics)

---

## Introduction to Crystal

Crystal is a statically-typed, compiled programming language with syntax inspired by Ruby. It combines the elegance and productivity of Ruby with the performance and type safety of compiled languages like C.

### Key Features:
- **Ruby-like syntax** - Clean, readable, and expressive
- **Statically typed** - Type errors caught at compile time
- **Type inference** - Minimal type annotations needed
- **Fast execution** - Compiled to native code via LLVM
- **Null reference checking** - Prevents null pointer errors
- **Concurrency** - Built-in support via fibers
- **C bindings** - Easy integration with C libraries
- **Macro system** - Powerful metaprogramming capabilities

### Use Cases:
- Web applications and APIs
- Command-line tools
- System programming
- Microservices
- Network applications
- Data processing
- High-performance applications

---

## Getting Started

### Installation

#### macOS
```bash
brew install crystal
```

#### Ubuntu/Debian
```bash
curl -fsSL https://crystal-lang.org/install.sh | sudo bash
```

#### Arch Linux
```bash
pacman -S crystal
```

#### From Source
```bash
git clone https://github.com/crystal-lang/crystal
cd crystal
make
```

### Verify Installation
```bash
crystal --version
```

### Creating Your First Program

Create a file `hello.cr`:
```crystal
puts "Hello, Crystal!"
```

Run it:
```bash
crystal hello.cr
```

Or compile and run:
```bash
crystal build hello.cr
./hello
```

### Crystal REPL
```bash
crystal i
```

---

## Basic Syntax

### Comments

```crystal
# Single-line comment

=begin
Multi-line
comment
=end
```

### Printing Output

```crystal
puts "Hello"           # With newline
print "Hello"          # Without newline
p [1, 2, 3]           # Debug print
pp {name: "Alice"}    # Pretty print
```

### String Interpolation

```crystal
name = "Crystal"
puts "Hello, #{name}!"
puts "2 + 2 = #{2 + 2}"
```

### Command Output

```crystal
result = `ls -la`
puts result
```

---

## Variables and Data Types

### Variables

```crystal
# Local variables
name = "Alice"
age = 30

# Instance variables
@count = 0

# Class variables
@@total = 100

# Constants
MAX_SIZE = 1000
PI = 3.14159
```

### Type Annotations

```crystal
# Explicit type annotation
name : String = "Bob"
age : Int32 = 25
price : Float64 = 19.99

# Type inference
x = 42          # Int32
y = 3.14        # Float64
z = "text"      # String
```

### Integer Types

```crystal
# Signed integers
i8 : Int8 = 127
i16 : Int16 = 32767
i32 : Int32 = 2147483647
i64 : Int64 = 9223372036854775807
i128 : Int128 = 170141183460469231731687303715884105727

# Unsigned integers
u8 : UInt8 = 255
u16 : UInt16 = 65535
u32 : UInt32 = 4294967295
u64 : UInt64 = 18446744073709551615
u128 : UInt128 = 340282366920938463463374607431768211455

# Number separators
million = 1_000_000
```

### Float Types

```crystal
f32 : Float32 = 3.14_f32
f64 : Float64 = 3.14159265359

# Scientific notation
large = 1.5e10
small = 1.5e-10
```

### Boolean Type

```crystal
is_active = true
is_valid = false
```

### Character Type

```crystal
char = 'a'
unicode = '💎'
```

### String Type

```crystal
# String literals
str1 = "Hello"
str2 = "World"

# Multi-line strings
multiline = "This is
a multi-line
string"

# Heredoc
text = <<-TEXT
  This is a heredoc.
  It preserves formatting.
TEXT

# String operations
concatenated = str1 + " " + str2
repeated = "ha" * 3              # "hahaha"
substring = "Hello"[0..2]        # "Hel"
```

### Symbol Type

```crystal
symbol = :hello
status = :active

# Symbols are interned strings
:foo.object_id == :foo.object_id  # true
```

### Nil Type

```crystal
value = nil
value : String?  # Nilable type
```

---

## Control Flow

### if/elsif/else

```crystal
age = 18

if age < 13
  puts "Child"
elsif age < 20
  puts "Teenager"
else
  puts "Adult"
end

# Inline if
puts "Adult" if age >= 18

# Ternary operator
status = age >= 18 ? "adult" : "minor"

# if as expression
result = if age >= 18
  "Can vote"
else
  "Cannot vote"
end
```

### unless

```crystal
unless age < 18
  puts "Adult"
end

# Inline unless
puts "Valid" unless email.empty?
```

### case/when

```crystal
grade = 'B'

case grade
when 'A'
  puts "Excellent"
when 'B', 'C'
  puts "Good"
when 'D'
  puts "Needs improvement"
else
  puts "Failed"
end

# Case with ranges
score = 85

case score
when 90..100
  "A"
when 80..89
  "B"
when 70..79
  "C"
else
  "F"
end

# Case with types
value = 42

case value
when Int32
  "Integer"
when String
  "String"
when Array
  "Array"
end
```

### while Loop

```crystal
count = 0

while count < 5
  puts count
  count += 1
end

# while as modifier
count += 1 while count < 10
```

### until Loop

```crystal
count = 0

until count == 5
  puts count
  count += 1
end
```

### Loop

```crystal
loop do
  print "Continue? (y/n): "
  answer = gets
  break if answer == "n"
end
```

### Iteration

```crystal
# Each
[1, 2, 3].each do |num|
  puts num
end

# With index
["a", "b", "c"].each_with_index do |char, i|
  puts "#{i}: #{char}"
end

# Times
5.times do |i|
  puts i
end

# Upto
1.upto(5) do |i|
  puts i
end

# Downto
5.downto(1) do |i|
  puts i
end

# Step
0.step(to: 10, by: 2) do |i|
  puts i
end
```

### Break and Next

```crystal
# Break
5.times do |i|
  break if i == 3
  puts i
end

# Next (continue)
5.times do |i|
  next if i == 2
  puts i
end

# Break with value
result = loop do
  break 42
end
puts result  # 42
```

---

## Methods

### Method Definition

```crystal
def greet
  "Hello!"
end

def greet(name)
  "Hello, #{name}!"
end

# Method with type annotations
def add(a : Int32, b : Int32) : Int32
  a + b
end

# Default parameters
def greet(name = "World")
  "Hello, #{name}!"
end

# Named arguments
def create_user(name : String, age : Int32, email : String = "")
  # ...
end

create_user(name: "Alice", age: 30)
create_user(age: 25, name: "Bob", email: "bob@example.com")
```

### Return Values

```crystal
def add(a, b)
  a + b  # Implicit return
end

def divide(a, b)
  return 0 if b == 0
  a / b
end

# Multiple return values
def min_max(array)
  {array.min, array.max}
end

min, max = min_max([3, 1, 4, 1, 5])
```

### Splat Arguments

```crystal
def sum(*numbers)
  numbers.sum
end

sum(1, 2, 3, 4, 5)  # 15

# Named splat
def create(**options)
  pp options
end

create(name: "Alice", age: 30)
```

### Method Overloading

```crystal
def process(x : Int32)
  x * 2
end

def process(x : String)
  x.upcase
end

process(5)        # 10
process("hello")  # "HELLO"
```

### Operator Methods

```crystal
class Point
  property x : Int32
  property y : Int32

  def initialize(@x, @y)
  end

  def +(other : Point)
    Point.new(@x + other.x, @y + other.y)
  end

  def -(other : Point)
    Point.new(@x - other.x, @y - other.y)
  end
end

p1 = Point.new(1, 2)
p2 = Point.new(3, 4)
p3 = p1 + p2  # Point(4, 6)
```

### Yield and Blocks

```crystal
def repeat(times)
  times.times do
    yield
  end
end

repeat(3) do
  puts "Hello"
end

# Yield with arguments
def each_char(str)
  str.each_char do |c|
    yield c
  end
end

each_char("abc") do |c|
  puts c
end
```

---

## Classes and Objects

### Class Definition

```crystal
class Person
  # Getter and setter
  property name : String

  # Getter only
  getter age : Int32

  # Setter only
  setter email : String?

  def initialize(@name, @age, @email = nil)
  end

  def greet
    "Hello, I'm #{@name}"
  end
end

person = Person.new("Alice", 30)
puts person.name
puts person.greet
```

### Instance Variables

```crystal
class Counter
  @count : Int32

  def initialize
    @count = 0
  end

  def increment
    @count += 1
  end

  def value
    @count
  end
end
```

### Class Variables and Methods

```crystal
class User
  @@count = 0

  def initialize
    @@count += 1
  end

  def self.total_users
    @@count
  end

  def self.create(name)
    new(name)
  end
end

User.new("Alice")
User.new("Bob")
puts User.total_users  # 2
```

### Inheritance

```crystal
class Animal
  property name : String

  def initialize(@name)
  end

  def speak
    "Some sound"
  end
end

class Dog < Animal
  def speak
    "Woof!"
  end

  def fetch
    "Fetching..."
  end
end

class Cat < Animal
  def speak
    "Meow!"
  end
end

dog = Dog.new("Rex")
puts dog.speak  # "Woof!"
```

### Method Visibility

```crystal
class MyClass
  def public_method
    "Public"
  end

  private def private_method
    "Private"
  end

  protected def protected_method
    "Protected"
  end
end
```

### Abstract Classes and Methods

```crystal
abstract class Shape
  abstract def area : Float64
  abstract def perimeter : Float64
end

class Circle < Shape
  property radius : Float64

  def initialize(@radius)
  end

  def area : Float64
    Math::PI * @radius ** 2
  end

  def perimeter : Float64
    2 * Math::PI * @radius
  end
end

class Rectangle < Shape
  property width : Float64
  property height : Float64

  def initialize(@width, @height)
  end

  def area : Float64
    @width * @height
  end

  def perimeter : Float64
    2 * (@width + @height)
  end
end
```

---

## Modules

### Module Definition

```crystal
module Logging
  def log(message)
    puts "[#{Time.local}] #{message}"
  end
end

class Application
  include Logging

  def run
    log("Application started")
  end
end

app = Application.new
app.run
```

### Extend vs Include

```crystal
module Greetings
  def hello
    "Hello!"
  end
end

# Include - adds instance methods
class Person
  include Greetings
end

Person.new.hello  # "Hello!"

# Extend - adds class methods
class Company
  extend Greetings
end

Company.hello  # "Hello!"
```

### Namespace Modules

```crystal
module MyApp
  module Database
    class Connection
      def connect
        "Connected to database"
      end
    end
  end

  module Services
    class UserService
      def create_user
        "User created"
      end
    end
  end
end

conn = MyApp::Database::Connection.new
service = MyApp::Services::UserService.new
```

---

## Structs

### Struct Definition

```crystal
struct Point
  property x : Int32
  property y : Int32

  def initialize(@x, @y)
  end

  def distance_from_origin
    Math.sqrt(@x ** 2 + @y ** 2)
  end
end

# Structs are value types
p1 = Point.new(3, 4)
p2 = p1
p2.x = 10
puts p1.x  # 3 (not affected)
```

### Struct vs Class

```crystal
# Structs are:
# - Passed by value
# - Allocated on the stack
# - Faster for small data
# - Immutable by design (best practice)

# Classes are:
# - Passed by reference
# - Allocated on the heap
# - Better for larger data
# - Can be inherited

struct Color
  getter r : UInt8
  getter g : UInt8
  getter b : UInt8

  def initialize(@r, @g, @b)
  end
end

class Image
  property width : Int32
  property height : Int32
  property pixels : Array(Color)

  def initialize(@width, @height)
    @pixels = Array(Color).new(@width * @height, Color.new(0, 0, 0))
  end
end
```

---

## Enums

### Enum Definition

```crystal
enum Color
  Red
  Green
  Blue
end

color = Color::Red

case color
when Color::Red
  puts "Red"
when Color::Green
  puts "Green"
when Color::Blue
  puts "Blue"
end

# Enum with values
enum Status
  Active   = 1
  Inactive = 2
  Pending  = 3
end

# Flags enum
@[Flags]
enum Permission
  Read
  Write
  Execute
end

perms = Permission::Read | Permission::Write
puts perms.includes?(Permission::Read)  # true
```

### Enum Methods

```crystal
enum Direction
  North
  South
  East
  West

  def opposite
    case self
    when North then South
    when South then North
    when East  then West
    when West  then East
    end
  end
end

dir = Direction::North
puts dir.opposite  # South
```

---

## Arrays and Tuples

### Arrays

```crystal
# Array creation
numbers = [1, 2, 3, 4, 5]
empty : Array(Int32) = []
strings = Array(String).new
filled = Array.new(5, 0)  # [0, 0, 0, 0, 0]

# Array operations
numbers << 6              # Append
numbers.push(7)
numbers.unshift(0)        # Prepend
numbers.pop               # Remove last
numbers.shift             # Remove first

# Access
first = numbers[0]
last = numbers[-1]
slice = numbers[1..3]

# Iteration
numbers.each { |n| puts n }
numbers.map { |n| n * 2 }
numbers.select { |n| n > 3 }
numbers.reject { |n| n < 3 }

# Methods
numbers.size
numbers.empty?
numbers.includes?(3)
numbers.sum
numbers.min
numbers.max
numbers.sort
numbers.reverse
numbers.uniq

# Multi-dimensional arrays
matrix = [[1, 2], [3, 4], [5, 6]]
puts matrix[1][0]  # 3
```

### Tuples

```crystal
# Tuple creation
tuple = {1, "hello", true}
named = {name: "Alice", age: 30}

# Access
first = tuple[0]
name = named[:name]

# Destructuring
x, y, z = {10, 20, 30}
{a, b} = {name: "Bob", age: 25}

# Tuples are immutable
# tuple[0] = 2  # Error!

# Named tuples
person = {name: "Alice", age: 30, city: "NYC"}
puts person[:name]
puts person[:age]
```

---

## Hashes

### Hash Creation

```crystal
# Hash literal
hash = {"name" => "Alice", "age" => 30}

# Hash with symbol keys
options = {:debug => true, :verbose => false}

# Alternative syntax
options = {debug: true, verbose: false}

# Typed hash
scores : Hash(String, Int32) = {} of String => Int32
scores["Alice"] = 100
scores["Bob"] = 95

# Hash creation
Hash(String, Int32).new
Hash(String, Int32).new(default_value: 0)
```

### Hash Operations

```crystal
hash = {name: "Alice", age: 30}

# Access
hash[:name]
hash[:name]?  # Returns nil if not found

# Modify
hash[:age] = 31
hash[:city] = "NYC"

# Delete
hash.delete(:age)

# Check
hash.has_key?(:name)
hash.has_value?(30)

# Iteration
hash.each do |key, value|
  puts "#{key}: #{value}"
end

hash.each_key { |k| puts k }
hash.each_value { |v| puts v }

# Methods
hash.keys
hash.values
hash.size
hash.empty?
hash.clear

# Merge
hash1 = {a: 1, b: 2}
hash2 = {b: 3, c: 4}
merged = hash1.merge(hash2)  # {a: 1, b: 3, c: 4}
```

---

## Ranges

### Range Types

```crystal
# Inclusive range
range1 = 1..10     # 1 to 10 (inclusive)
range2 = 'a'..'z'  # 'a' to 'z'

# Exclusive range
range3 = 1...10    # 1 to 9 (exclusive)

# Range methods
range1.includes?(5)  # true
range1.covers?(5)    # true
range1.begin         # 1
range1.end           # 10

# Iteration
(1..5).each { |i| puts i }

# Convert to array
array = (1..5).to_a  # [1, 2, 3, 4, 5]

# Use in case
age = 25
case age
when 0..12
  "Child"
when 13..19
  "Teenager"
when 20..64
  "Adult"
else
  "Senior"
end
```

---

## Blocks and Procs

### Blocks

```crystal
# Block with do...end
[1, 2, 3].each do |n|
  puts n * 2
end

# Block with braces
[1, 2, 3].each { |n| puts n * 2 }

# Method with block
def repeat(times)
  times.times do
    yield
  end
end

repeat(3) { puts "Hello" }

# Block with arguments
def transform(value)
  yield value
end

result = transform(5) { |x| x * 2 }  # 10
```

### Procs

```crystal
# Proc literal
square = ->(x : Int32) { x * x }
puts square.call(5)  # 25

# Proc.new
add = Proc(Int32, Int32, Int32).new do |a, b|
  a + b
end
puts add.call(3, 4)  # 7

# Passing procs
def apply(value, operation : Proc(Int32, Int32))
  operation.call(value)
end

double = ->(x : Int32) { x * 2 }
puts apply(5, double)  # 10

# Block to proc
def filter(array, &block : Int32 -> Bool)
  array.select(&block)
end

numbers = [1, 2, 3, 4, 5]
evens = filter(numbers) { |x| x % 2 == 0 }
```

---

## Type System

### Union Types

```crystal
# Variable can be multiple types
value : Int32 | String = 42
value = "hello"

# Handling union types
def process(value : Int32 | String)
  case value
  when Int32
    value * 2
  when String
    value.upcase
  end
end

# Nilable types
name : String? = nil  # Same as String | Nil
name = "Alice"
```

### Type Checking

```crystal
value = 42

if value.is_a?(Int32)
  puts "Integer: #{value}"
end

# Responds to
if value.responds_to?(:to_s)
  puts value.to_s
end

# Type assertion
str : String = value.as(String)  # Raises if not String

# Type cast
num = value.as?(String)  # Returns nil if not String
```

### Alias Types

```crystal
alias Name = String
alias Age = Int32
alias UserId = Int64

name : Name = "Alice"
age : Age = 30

# Complex type alias
alias StringOrInt = String | Int32
alias NamedTuple = {name: String, age: Int32}
```

### Typeof

```crystal
x = 42
y = typeof(x)  # Int32

# Multiple values
z = typeof(1, "hello", true)  # Int32 | String | Bool
```

---

## Generics

### Generic Classes

```crystal
class Box(T)
  @value : T

  def initialize(@value : T)
  end

  def get : T
    @value
  end

  def set(@value : T)
  end
end

int_box = Box(Int32).new(42)
str_box = Box(String).new("hello")

# Type inference
box = Box.new(42)  # Box(Int32)
```

### Generic Methods

```crystal
def first(array : Array(T)) : T forall T
  array[0]
end

puts first([1, 2, 3])        # 1
puts first(["a", "b", "c"])  # "a"

# Multiple type parameters
def pair(a : T, b : U) : {T, U} forall T, U
  {a, b}
end

p pair(1, "hello")  # {1, "hello"}
```

### Generic Constraints

```crystal
# Must respond to certain methods
def sum(values : Array(T)) : T forall T
  total = T.zero
  values.each do |v|
    total += v
  end
  total
end

# Must be a specific type or subtype
class Container(T)
  def initialize(@value : T)
  end

  def process where T <= Number
    @value * 2
  end
end
```

---

## Exception Handling

### Raising Exceptions

```crystal
# Raise exception
raise "Something went wrong"
raise ArgumentError.new("Invalid argument")

# Custom exception
class MyError < Exception
end

raise MyError.new("Custom error")
```

### Rescue

```crystal
begin
  result = 10 / 0
rescue ex : DivisionByZeroError
  puts "Cannot divide by zero"
  puts ex.message
rescue ex
  puts "An error occurred: #{ex}"
ensure
  puts "This always runs"
end

# Rescue inline
value = begin
  risky_operation
rescue
  default_value
end

# Rescue specific types
begin
  # code
rescue ex : ArgumentError | KeyError
  puts "Argument or Key error: #{ex}"
end
```

### Custom Exceptions

```crystal
class ValidationError < Exception
  getter field : String

  def initialize(@field, message)
    super(message)
  end
end

def validate_age(age)
  raise ValidationError.new("age", "Must be positive") if age < 0
  raise ValidationError.new("age", "Must be under 150") if age > 150
  age
end

begin
  validate_age(-5)
rescue ex : ValidationError
  puts "#{ex.field}: #{ex.message}"
end
```

---

## File I/O

### Reading Files

```crystal
# Read entire file
content = File.read("file.txt")
lines = File.read_lines("file.txt")

# Read with block
File.open("file.txt") do |file|
  file.each_line do |line|
    puts line
  end
end

# Read binary
data = File.read("image.png", mode: "rb")
```

### Writing Files

```crystal
# Write entire file
File.write("output.txt", "Hello, World!")

# Write with block
File.open("output.txt", "w") do |file|
  file.puts "Line 1"
  file.puts "Line 2"
  file.print "No newline"
end

# Append
File.open("output.txt", "a") do |file|
  file.puts "Appended line"
end
```

### File Operations

```crystal
# Check existence
File.exists?("file.txt")
Dir.exists?("directory")

# File info
File.size("file.txt")
File.info("file.txt").modification_time

# Delete
File.delete("file.txt")

# Rename/Move
File.rename("old.txt", "new.txt")

# Copy
File.copy("source.txt", "dest.txt")

# Create directory
Dir.mkdir("new_dir")
Dir.mkdir_p("path/to/dir")  # Create with parents

# List directory
Dir.entries(".")
Dir.glob("*.txt")

# Change directory
Dir.cd("new_dir") do
  # Work in directory
end
```

### Path Operations

```crystal
require "file_utils"

# Join paths
path = File.join("dir", "subdir", "file.txt")

# Basename and dirname
File.basename("/path/to/file.txt")  # "file.txt"
File.dirname("/path/to/file.txt")   # "/path/to"

# Extension
File.extname("file.txt")  # ".txt"

# Expand path
File.expand_path("~/file.txt")

# Temp files
File.tempfile("prefix") do |file|
  file.puts "Temporary content"
end
```

---

## Concurrency

### Fibers

```crystal
# Create fiber
fiber = spawn do
  5.times do |i|
    puts "Fiber: #{i}"
    Fiber.yield
  end
end

# Main fiber
5.times do |i|
  puts "Main: #{i}"
  Fiber.yield
end

# Automatic spawning
spawn do
  sleep 1
  puts "After 1 second"
end

spawn do
  sleep 2
  puts "After 2 seconds"
end

sleep 3  # Wait for fibers
```

### Channels

```crystal
# Create channel
channel = Channel(Int32).new

# Send to channel
spawn do
  5.times do |i|
    channel.send(i)
  end
  channel.close
end

# Receive from channel
spawn do
  while value = channel.receive?
    puts "Received: #{value}"
  end
end

sleep 1

# Buffered channel
buffered = Channel(String).new(10)

# Select
ch1 = Channel(Int32).new
ch2 = Channel(String).new

spawn { ch1.send(42) }
spawn { ch2.send("hello") }

select
when value = ch1.receive
  puts "Got int: #{value}"
when value = ch2.receive
  puts "Got string: #{value}"
end
```

### Parallel Execution

```crystal
# Parallel map
results = (1..10).parallel_map do |i|
  expensive_operation(i)
end

# Wait group pattern
done = Channel(Nil).new

10.times do |i|
  spawn do
    do_work(i)
    done.send(nil)
  end
end

10.times { done.receive }
```

---

## Macros

### Basic Macros

```crystal
macro define_method(name, content)
  def {{name}}
    {{content}}
  end
end

define_method greet, "Hello!"

puts greet  # "Hello!"

# Macro with arguments
macro create_property(name, type)
  @{{name}} : {{type}}

  def {{name}}
    @{{name}}
  end

  def {{name}}=(value : {{type}})
    @{{name}} = value
  end
end

class Person
  create_property name, String
  create_property age, Int32

  def initialize(@name, @age)
  end
end
```

### Compile-Time Reflection

```crystal
macro generate_methods
  {% for name in %w(foo bar baz) %}
    def {{name.id}}
      {{name}}
    end
  {% end %}
end

class MyClass
  generate_methods
end

# Type information at compile time
macro inspect_type(obj)
  puts {{obj.class_name}}
end

inspect_type(42)  # Prints "Int32" at compile time
```

### Macro Conditionals

```crystal
macro ifdef(flag, &block)
  {% if flag?(flag.id) %}
    {{block.body}}
  {% end %}
end

ifdef :debug do
  puts "Debug mode enabled"
end

# Run with: crystal build -Ddebug file.cr
```

### Hook Macros

```crystal
class Base
  macro inherited
    puts "{{@type}} inherits from Base"
  end
end

class Child < Base  # Prints "Child inherits from Base"
end

# Method added hook
macro method_added(method)
  puts "Added method: {{method.name}}"
end
```

---

## C Bindings

### Lib Definition

```crystal
@[Link("m")]
lib LibM
  fun sqrt(x : Float64) : Float64
  fun pow(x : Float64, y : Float64) : Float64
end

puts LibM.sqrt(16.0)  # 4.0
puts LibM.pow(2.0, 3.0)  # 8.0
```

### Struct Binding

```crystal
lib LibC
  struct TimeSpec
    tv_sec : Int64
    tv_nsec : Int64
  end

  fun clock_gettime(clk_id : Int32, tp : TimeSpec*) : Int32
end

time = LibC::TimeSpec.new
LibC.clock_gettime(0, pointerof(time))
puts time.tv_sec
```

### Callbacks

```crystal
lib LibEvent
  alias Callback = (Int32, Int16, Void*) -> Void

  fun event_set(ev : Void*, fd : Int32, events : Int16,
                callback : Callback, arg : Void*)
end

callback = ->(fd : Int32, events : Int16, arg : Void*) {
  puts "Event on fd #{fd}"
}
```

### Common C Bindings

```crystal
# String handling
lib LibC
  fun strlen(s : UInt8*) : UInt64
  fun strcmp(s1 : UInt8*, s2 : UInt8*) : Int32
end

# Memory operations
lib LibC
  fun malloc(size : UInt64) : Void*
  fun free(ptr : Void*)
  fun memcpy(dest : Void*, src : Void*, n : UInt64) : Void*
end

# File operations
lib LibC
  fun fopen(path : UInt8*, mode : UInt8*) : Void*
  fun fclose(file : Void*) : Int32
end
```

---

## Testing

### Spec Framework

```crystal
require "spec"

describe Array do
  describe "#size" do
    it "returns the number of elements" do
      [1, 2, 3].size.should eq(3)
    end

    it "returns 0 for empty array" do
      ([] of Int32).size.should eq(0)
    end
  end

  describe "#push" do
    it "adds element to end" do
      array = [1, 2]
      array.push(3)
      array.should eq([1, 2, 3])
    end
  end
end
```

### Expectations

```crystal
require "spec"

describe "Expectations" do
  it "tests equality" do
    1.should eq(1)
    1.should_not eq(2)
  end

  it "tests truthiness" do
    true.should be_true
    false.should be_false
    nil.should be_nil
  end

  it "tests inclusion" do
    [1, 2, 3].should contain(2)
  end

  it "tests exceptions" do
    expect_raises(DivisionByZeroError) do
      1 / 0
    end
  end

  it "tests types" do
    42.should be_a(Int32)
    "hello".should be_a(String)
  end
end
```

### Hooks

```crystal
require "spec"

describe "Hooks" do
  before_each do
    @value = 0
  end

  after_each do
    @value = nil
  end

  it "test 1" do
    @value.should eq(0)
  end

  it "test 2" do
    @value.should eq(0)
  end
end

# around_each
around_each do |example|
  setup
  example.run
  teardown
end
```

### Test Helpers

```crystal
require "spec"

# Pending tests
pending "implement this feature" do
  # Not yet implemented
end

# Focus on specific test
it "important test", focus: true do
  # This test will run
end

# Tags
it "slow test", tags: "slow" do
  # Can be filtered
end

# Run with: crystal spec --tag slow
```

---

## Project Structure

### Shards (Package Manager)

#### shard.yml
```yaml
name: my_project
version: 0.1.0

authors:
  - Your Name <your.email@example.com>

crystal: 1.18.2

license: MIT

dependencies:
  kemal:
    github: kemalcr/kemal
    version: ~> 1.0.0

development_dependencies:
  ameba:
    github: crystal-ameba/ameba
    version: ~> 1.0.0

targets:
  my_project:
    main: src/my_project.cr

scripts:
  postinstall: make setup
```

### Install Dependencies

```bash
shards install
shards update
shards list
```

### Project Structure

```
my_project/
├── shard.yml
├── shard.lock
├── README.md
├── LICENSE
├── src/
│   ├── my_project.cr       # Main entry point
│   ├── my_project/
│   │   ├── version.cr
│   │   ├── config.cr
│   │   └── models/
│   │       └── user.cr
│   └── views/
│       └── template.ecr
├── spec/
│   ├── spec_helper.cr
│   ├── my_project_spec.cr
│   └── models/
│       └── user_spec.cr
├── bin/
│   └── my_project
└── lib/                    # Dependencies (auto-generated)
```

### Main File Structure

```crystal
# src/my_project.cr
require "./my_project/**"

module MyProject
  VERSION = "0.1.0"

  def self.run
    # Application logic
  end
end

MyProject.run
```

### Version File

```crystal
# src/my_project/version.cr
module MyProject
  VERSION = {{ `shards version #{__DIR__}`.chomp.stringify }}
end
```

---

## Advanced Topics

### Annotations

```crystal
@[AlwaysInline]
def fast_method
  # Compiler will always inline this
end

@[NoInline]
def debug_method
  # Compiler will never inline this
end

@[Packed]
struct PackedData
  # No padding between fields
end

@[Flags]
enum FileMode
  Read
  Write
  Execute
end
```

### Pointers

```crystal
# Pointer basics
x = 42
ptr = pointerof(x)
value = ptr.value  # 42

# Pointer arithmetic
array = [1, 2, 3, 4, 5]
ptr = array.to_unsafe
ptr[2]  # 3
(ptr + 2).value  # 3

# Null pointer
null_ptr = Pointer(Int32).null
null_ptr.null?  # true
```

### Unsafe Code

```crystal
# Unsafe memory allocation
ptr = Pointer(Int32).malloc(10)
ptr[0] = 42
ptr[0]  # 42
ptr.free  # Don't forget to free!

# Uninitialized value
value = uninitialized Int32
value = 42

# Cast pointers
int_ptr = Pointer(Int32).malloc(1)
byte_ptr = int_ptr.as(Pointer(UInt8))
```

### Compile-Time Constants

```crystal
# Compile-time execution
SIZE = {{ 10 * 20 }}  # 200

# Read file at compile time
CONFIG = {{ read_file("config.json") }}

# Run command at compile time
GIT_COMMIT = {{ `git rev-parse HEAD`.chomp.stringify }}
```

### Method Missing

```crystal
class DynamicObject
  def method_missing(call)
    puts "Called: #{call.name} with #{call.args}"
  end
end

obj = DynamicObject.new
obj.some_method(1, 2, 3)  # Called: some_method with (1, 2, 3)
```

### Operator Overloading

```crystal
class Vector
  property x : Float64
  property y : Float64

  def initialize(@x, @y)
  end

  def +(other : Vector)
    Vector.new(@x + other.x, @y + other.y)
  end

  def -(other : Vector)
    Vector.new(@x - other.x, @y - other.y)
  end

  def *(scalar : Float64)
    Vector.new(@x * scalar, @y * scalar)
  end

  def [](index : Int32)
    index == 0 ? @x : @y
  end

  def []=(index : Int32, value : Float64)
    index == 0 ? @x = value : @y = value
  end
end
```

---

## Resources and Best Practices

### Essential Tools

- **crystal** - Compiler
- **shards** - Dependency manager
- **crystal tool** - Built-in tools (format, context, hierarchy)
- **ameba** - Linter
- **crystal spec** - Test runner

### Useful Commands

```bash
# Format code
crystal tool format

# Check code without compilation
crystal build --no-codegen file.cr

# Show type hierarchy
crystal tool hierarchy

# Show implementations
crystal tool implementations

# Generate documentation
crystal docs

# Install dependencies
shards install

# Run tests
crystal spec

# Lint code
ameba
```

### Learning Resources

- [Official Crystal Documentation](https://crystal-lang.org/reference/)
- [Crystal API Documentation](https://crystal-lang.org/api/)
- [Crystal by Example](https://crystal-lang.org/examples/)
- [Crystal Forum](https://forum.crystal-lang.org/)
- [Crystal Shards](https://crystalshards.org/)

### Popular Shards

#### Web Frameworks
- **kemal** - Sinatra-like web framework
- **lucky** - Full-stack web framework
- **amber** - Productive web framework

#### Database
- **crystal-db** - Common database API
- **crystal-pg** - PostgreSQL driver
- **crystal-mysql** - MySQL driver
- **crystal-sqlite3** - SQLite3 driver

#### HTTP Clients
- **crest** - HTTP and REST client
- **halite** - HTTP requests

#### JSON/Serialization
- **json** - Built-in JSON support
- **yaml** - YAML parser
- **xml** - XML parser

#### Testing
- **spec** - Built-in testing framework
- **webmock** - HTTP mocking

#### Utilities
- **dotenv** - Environment variables
- **time** - Date and time handling
- **log** - Logging

This course provides a comprehensive foundation in Crystal programming. Practice by building projects and exploring the Crystal ecosystem!
