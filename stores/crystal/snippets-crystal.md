## Hello World
Basic program structure and main execution
```crystal
puts "Hello, World!"
```

----------

## Variable Declaration
Different ways to declare and initialize variables
```crystal
# Immutable by default
name = "Alice"
age = 30
city = "New York"

# With type annotation
score : Int32 = 100
price : Float64 = 19.99
active : Bool = true

# Type inference
x = 42          # Int32
y = 3.14        # Float64
z = "text"      # String

puts "Name: #{name}, Age: #{age}, City: #{city}"
```

----------

## Constants
Declaring constants
```crystal
PI = 3.14159
MAX_USERS = 100
APP_NAME = "MyApp"

module Config
  VERSION = "1.0.0"
  DEBUG = true
end

puts "#{Config::APP_NAME} v#{Config::VERSION}"
```

----------

## Data Types
Working with basic data types
```crystal
# Integers
int8 : Int8 = 127
int32 : Int32 = 2147483647
int64 : Int64 = 9223372036854775807
uint32 : UInt32 = 4294967295

# Floats
float32 : Float32 = 3.14_f32
float64 : Float64 = 3.14159

# Boolean
is_active = true
is_valid = false

# Character
char = 'A'
emoji = '💎'

# String
text = "Hello, Crystal!"
multiline = "Line 1
Line 2"

# Symbol
status = :active
color = :red

# Nil
value = nil
nullable : String? = nil

puts "Types: #{int32}, #{float64}, #{is_active}, #{char}, #{text}"
```

----------

## Arrays
Creating and manipulating arrays
```crystal
# Array creation
numbers = [1, 2, 3, 4, 5]
strings = ["apple", "banana", "cherry"]
empty : Array(Int32) = []

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
doubled = numbers.map { |n| n * 2 }
evens = numbers.select { |n| n % 2 == 0 }

# Useful methods
puts "Size: #{numbers.size}"
puts "Sum: #{numbers.sum}"
puts "Max: #{numbers.max}"
puts "Includes 3: #{numbers.includes?(3)}"
```

----------

## Hashes
Creating and working with hashes
```crystal
# Hash creation
person = {"name" => "Alice", "age" => 30}
options = {debug: true, verbose: false}

# Typed hash
scores = {} of String => Int32
scores["Alice"] = 100
scores["Bob"] = 95

# Access
name = person["name"]
age = person["age"]?  # Returns nil if not found

# Modify
person["city"] = "NYC"
person.delete("age")

# Iteration
person.each do |key, value|
  puts "#{key}: #{value}"
end

# Methods
puts "Keys: #{person.keys}"
puts "Values: #{person.values}"
puts "Has key? #{person.has_key?("name")}"
```

----------

## Tuples
Working with tuples and named tuples
```crystal
# Tuple
coordinates = {10, 20, 30}
x, y, z = coordinates

# Named tuple
person = {name: "Alice", age: 30, city: "NYC"}
puts person[:name]
puts person[:age]

# Destructuring
{name, age} = {name: "Bob", age: 25}
puts "#{name} is #{age} years old"

# Tuples are immutable
# coordinates[0] = 15  # Error!
```

----------

## Ranges
Using ranges for sequences
```crystal
# Inclusive range
range1 = 1..10     # 1 to 10
range2 = 'a'..'z'  # 'a' to 'z'

# Exclusive range
range3 = 1...10    # 1 to 9

# Iteration
(1..5).each { |i| puts i }

# Convert to array
array = (1..5).to_a

# Check inclusion
puts (1..10).includes?(5)  # true

# Use in case
age = 25
category = case age
when 0..12   then "Child"
when 13..19  then "Teenager"
when 20..64  then "Adult"
else              "Senior"
end
```

----------

## String Manipulation
Common string operations
```crystal
text = "  Hello, World!  "

# Basic operations
puts text.upcase
puts text.downcase
puts text.capitalize
puts text.reverse
puts text.trim

# String methods
puts text.size
puts text.empty?
puts text.includes?("World")
puts text.starts_with?("Hello")
puts text.ends_with?("!")

# Split and join
words = text.split(',')
joined = words.join(" - ")

# Interpolation
name = "Crystal"
puts "Hello, #{name}!"

# Substring
puts "Crystal"[0..3]  # "Crys"

# Character access
puts "Hello"[0]  # 'H'
```

----------

## Control Flow - If/Else
Conditional statements
```crystal
age = 25

if age >= 18
  puts "Adult"
else
  puts "Minor"
end

# Inline if
puts "Adult" if age >= 18

# If/elsif/else
score = 85
if score >= 90
  puts "A"
elsif score >= 80
  puts "B"
elsif score >= 70
  puts "C"
else
  puts "F"
end

# Ternary operator
status = age >= 18 ? "adult" : "minor"

# Unless
unless age < 18
  puts "Adult"
end
```

----------

## Case Statements
Pattern matching with case/when
```crystal
grade = 'B'

case grade
when 'A'
  puts "Excellent"
when 'B', 'C'
  puts "Good"
when 'D'
  puts "Pass"
else
  puts "Fail"
end

# Case with ranges
score = 85
result = case score
when 90..100 then "A"
when 80..89  then "B"
when 70..79  then "C"
when 60..69  then "D"
else              "F"
end

# Case with types
value = 42
type_name = case value
when Int32   then "Integer"
when String  then "String"
when Array   then "Array"
else              "Unknown"
end
```

----------

## Loops
Different types of loops
```crystal
# While loop
count = 0
while count < 5
  puts count
  count += 1
end

# Until loop
count = 0
until count == 5
  puts count
  count += 1
end

# Loop with break
loop do
  print "Continue? (y/n): "
  answer = gets
  break if answer == "n"
end

# Times
5.times do |i|
  puts "Iteration #{i}"
end

# Upto/Downto
1.upto(5) { |i| puts i }
5.downto(1) { |i| puts i }

# Each
[1, 2, 3].each { |n| puts n }
```

----------

## Methods
Method definition and calling
```crystal
# Basic method
def greet
  "Hello!"
end

# Method with parameters
def greet(name)
  "Hello, #{name}!"
end

# With type annotations
def add(a : Int32, b : Int32) : Int32
  a + b
end

# Default parameters
def greet(name = "World")
  "Hello, #{name}!"
end

# Named arguments
def create_user(name : String, age : Int32, email = "")
  {name: name, age: age, email: email}
end

user = create_user(name: "Alice", age: 30)

# Multiple return values
def min_max(array)
  {array.min, array.max}
end

min, max = min_max([3, 1, 4, 1, 5])
```

----------

## Classes
Defining and using classes
```crystal
class Person
  property name : String
  getter age : Int32
  setter email : String?

  def initialize(@name, @age, @email = nil)
  end

  def greet
    "Hello, I'm #{@name}"
  end

  def birthday
    @age += 1
  end

  def self.species
    "Homo sapiens"
  end
end

# Create instance
person = Person.new("Alice", 30)
puts person.name
puts person.greet
person.birthday
puts Person.species
```

----------

## Inheritance
Class inheritance and method overriding
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
    "#{@name} is fetching"
  end
end

class Cat < Animal
  def speak
    "Meow!"
  end
end

dog = Dog.new("Rex")
puts dog.speak     # "Woof!"
puts dog.fetch     # "Rex is fetching"

cat = Cat.new("Whiskers")
puts cat.speak     # "Meow!"
```

----------

## Modules
Creating and using modules
```crystal
module Logging
  def log(message)
    puts "[#{Time.local}] #{message}"
  end
end

module Validation
  def valid_email?(email)
    email.includes?("@")
  end
end

class Application
  include Logging
  include Validation

  def run
    log("Application started")
    if valid_email?("user@example.com")
      log("Email is valid")
    end
  end
end

app = Application.new
app.run
```

----------

## Structs
Value types with structs
```crystal
struct Point
  property x : Int32
  property y : Int32

  def initialize(@x, @y)
  end

  def distance_from_origin
    Math.sqrt(@x ** 2 + @y ** 2)
  end

  def +(other : Point)
    Point.new(@x + other.x, @y + other.y)
  end
end

# Structs are value types
p1 = Point.new(3, 4)
p2 = Point.new(1, 2)
p3 = p1 + p2

puts "Point: (#{p3.x}, #{p3.y})"
puts "Distance: #{p1.distance_from_origin}"
```

----------

## Enums
Defining and using enums
```crystal
enum Color
  Red
  Green
  Blue
  Yellow
end

# With values
enum Status
  Active   = 1
  Inactive = 2
  Pending  = 3
end

# Using enums
color = Color::Red

case color
when Color::Red
  puts "Red color"
when Color::Green
  puts "Green color"
when Color::Blue
  puts "Blue color"
else
  puts "Other color"
end

# Enum methods
puts Status::Active.value  # 1

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

----------

## Blocks and Procs
Working with blocks and procs
```crystal
# Block with each
[1, 2, 3].each do |num|
  puts num * 2
end

# Block with curly braces
[1, 2, 3].map { |n| n * 2 }

# Proc literal
square = ->(x : Int32) { x * x }
puts square.call(5)  # 25

# Proc with multiple arguments
add = ->(a : Int32, b : Int32) { a + b }
puts add.call(3, 4)  # 7

# Method with yield
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

result = transform(5) { |x| x * 2 }
puts result  # 10
```

----------

## Iterators
Common iteration methods
```crystal
numbers = [1, 2, 3, 4, 5]

# Map
doubled = numbers.map { |n| n * 2 }

# Select (filter)
evens = numbers.select { |n| n % 2 == 0 }

# Reject
odds = numbers.reject { |n| n % 2 == 0 }

# Reduce
sum = numbers.reduce(0) { |acc, n| acc + n }
product = numbers.reduce(1) { |acc, n| acc * n }

# Each with index
numbers.each_with_index do |num, idx|
  puts "#{idx}: #{num}"
end

# Any/All
has_even = numbers.any? { |n| n % 2 == 0 }
all_positive = numbers.all? { |n| n > 0 }

# Find
first_even = numbers.find { |n| n % 2 == 0 }

# Count
even_count = numbers.count { |n| n % 2 == 0 }
```

----------

## Exception Handling
Handling errors with rescue
```crystal
def divide(a, b)
  raise "Division by zero" if b == 0
  a / b
end

# Basic rescue
begin
  result = divide(10, 0)
rescue ex
  puts "Error: #{ex.message}"
end

# Rescue specific types
begin
  # risky code
  File.read("missing.txt")
rescue ex : File::NotFoundError
  puts "File not found"
rescue ex : Exception
  puts "Other error: #{ex}"
ensure
  puts "This always runs"
end

# Inline rescue
value = begin
  risky_operation
rescue
  default_value
end

# Custom exception
class ValidationError < Exception
end

begin
  raise ValidationError.new("Invalid input")
rescue ex : ValidationError
  puts ex.message
end
```

----------

## File I/O
Reading and writing files
```crystal
# Write to file
File.write("test.txt", "Hello, Crystal!")

# Read entire file
content = File.read("test.txt")
puts content

# Read lines
lines = File.read_lines("test.txt")
lines.each { |line| puts line }

# Read with block
File.open("test.txt") do |file|
  file.each_line do |line|
    puts line
  end
end

# Append to file
File.open("test.txt", "a") do |file|
  file.puts "Appended line"
end

# Check if file exists
if File.exists?("test.txt")
  puts "File exists"
end

# File info
puts File.size("test.txt")

# Delete file
File.delete("test.txt") if File.exists?("test.txt")
```

----------

## JSON Handling
Working with JSON data
```crystal
require "json"

# Define struct for JSON
struct User
  include JSON::Serializable

  property name : String
  property age : Int32
  property email : String?
end

# Create object
user = User.new(name: "Alice", age: 30, email: "alice@example.com")

# Serialize to JSON
json_string = user.to_json
puts json_string

# Deserialize from JSON
json_data = %({ "name": "Bob", "age": 25, "email": "bob@example.com" })
parsed_user = User.from_json(json_data)
puts parsed_user.name

# Working with JSON::Any
json = JSON.parse(%({ "name": "Charlie", "scores": [85, 90, 95] }))
puts json["name"]
puts json["scores"][0]
```

----------

## HTTP Client
Making HTTP requests
```crystal
require "http/client"

# GET request
response = HTTP::Client.get("https://api.github.com")
puts response.status_code
puts response.body

# GET with headers
headers = HTTP::Headers{"User-Agent" => "Crystal"}
response = HTTP::Client.get("https://api.github.com", headers: headers)

# POST request
response = HTTP::Client.post(
  "https://httpbin.org/post",
  headers: HTTP::Headers{"Content-Type" => "application/json"},
  body: %({ "name": "Alice", "age": 30 })
)

# Using client instance
client = HTTP::Client.new("api.example.com", tls: true)
response = client.get("/users")
puts response.body
client.close
```

----------

## HTTP Server
Creating a simple HTTP server
```crystal
require "http/server"

server = HTTP::Server.new do |context|
  path = context.request.path
  method = context.request.method

  context.response.content_type = "text/plain"

  case {method, path}
  when {"GET", "/"}
    context.response.print "Hello, World!"
  when {"GET", /\/hello\/(.+)/}
    name = $1
    context.response.print "Hello, #{name}!"
  when {"POST", "/data"}
    body = context.request.body.try(&.gets_to_end)
    context.response.print "Received: #{body}"
  else
    context.response.status_code = 404
    context.response.print "Not Found"
  end
end

address = server.bind_tcp(8080)
puts "Listening on http://#{address}"
server.listen
```

----------

## Concurrency with Fibers
Concurrent execution with fibers
```crystal
# Spawn a fiber
spawn do
  5.times do |i|
    puts "Fiber 1: #{i}"
    sleep 0.1
  end
end

spawn do
  5.times do |i|
    puts "Fiber 2: #{i}"
    sleep 0.1
  end
end

# Wait for fibers
sleep 1

# Using channels
channel = Channel(Int32).new

spawn do
  5.times do |i|
    channel.send(i)
  end
  channel.close
end

while value = channel.receive?
  puts "Received: #{value}"
end
```

----------

## Channels
Communication between fibers
```crystal
# Create channel
channel = Channel(String).new

# Producer fiber
spawn do
  ["apple", "banana", "cherry"].each do |fruit|
    channel.send(fruit)
    sleep 0.1
  end
  channel.close
end

# Consumer fiber
spawn do
  while message = channel.receive?
    puts "Got: #{message}"
  end
end

sleep 1

# Buffered channel
buffered = Channel(Int32).new(capacity: 10)

10.times do |i|
  buffered.send(i)
end

10.times do
  puts buffered.receive
end
```

----------

## Testing
Unit tests with spec
```crystal
require "spec"

def add(a, b)
  a + b
end

def divide(a, b)
  raise "Division by zero" if b == 0
  a / b
end

describe "Math operations" do
  describe "#add" do
    it "adds two numbers" do
      add(2, 3).should eq(5)
    end

    it "handles negative numbers" do
      add(-1, 1).should eq(0)
    end
  end

  describe "#divide" do
    it "divides two numbers" do
      divide(10, 2).should eq(5)
    end

    it "raises on division by zero" do
      expect_raises(Exception, "Division by zero") do
        divide(10, 0)
      end
    end
  end
end
```

----------

## Regular Expressions
Pattern matching with regex
```crystal
# Regex literal
email_regex = /^[\w\.-]+@[\w\.-]+\.\w+$/

# Test match
email = "user@example.com"
if email =~ email_regex
  puts "Valid email"
end

# Match object
if match = email.match(email_regex)
  puts "Matched: #{match[0]}"
end

# Capture groups
phone_regex = /(\d{3})-(\d{3})-(\d{4})/
phone = "555-123-4567"

if match = phone.match(phone_regex)
  puts "Area code: #{match[1]}"
  puts "Exchange: #{match[2]}"
  puts "Number: #{match[3]}"
end

# Replace with regex
text = "Hello 123 World 456"
result = text.gsub(/\d+/, "X")  # "Hello X World X"

# Scan for all matches
numbers = text.scan(/\d+/)
puts numbers.inspect
```

----------

## Date and Time
Working with time
```crystal
require "time"

# Current time
now = Time.local
utc_now = Time.utc

puts "Local: #{now}"
puts "UTC: #{utc_now}"

# Create specific time
birthday = Time.local(1990, 6, 15, 10, 30, 0)
puts "Birthday: #{birthday}"

# Format time
puts now.to_s("%Y-%m-%d %H:%M:%S")
puts now.to_s("%B %d, %Y")

# Time arithmetic
tomorrow = now + 1.day
yesterday = now - 1.day
next_week = now + 1.week

# Time span
span = Time.measure do
  sleep 1
end
puts "Took: #{span.total_seconds} seconds"

# Compare times
if now > birthday
  puts "Birthday has passed"
end

# Parse time
parsed = Time.parse("2023-12-25 10:30:00", "%Y-%m-%d %H:%M:%S", Time::Location::UTC)
```

----------

## Command Line Arguments
Parsing CLI arguments
```crystal
# Access arguments
ARGV.each_with_index do |arg, i|
  puts "Argument #{i}: #{arg}"
end

# Using OptionParser
require "option_parser"

name = "World"
verbose = false
count = 1

OptionParser.parse do |parser|
  parser.banner = "Usage: myapp [arguments]"

  parser.on("-n NAME", "--name=NAME", "Specify name") do |n|
    name = n
  end

  parser.on("-v", "--verbose", "Enable verbose output") do
    verbose = true
  end

  parser.on("-c COUNT", "--count=COUNT", "Repeat count") do |c|
    count = c.to_i
  end

  parser.on("-h", "--help", "Show help") do
    puts parser
    exit
  end
end

count.times do
  puts verbose ? "Greeting: Hello, #{name}!" : "Hello, #{name}!"
end
```

----------

## Environment Variables
Working with environment variables
```crystal
# Get environment variable
home = ENV["HOME"]?
path = ENV["PATH"]

# Get with default value
debug = ENV["DEBUG"]? || "false"

# Set environment variable
ENV["MY_VAR"] = "my_value"

# Check if exists
if ENV.has_key?("HOME")
  puts "HOME is set"
end

# Iterate all
ENV.each do |key, value|
  puts "#{key} = #{value}"
end

# Delete
ENV.delete("MY_VAR")
```

----------

## Macros
Compile-time code generation
```crystal
# Simple macro
macro debug(expression)
  puts "DEBUG: {{expression}} = #{{{expression}}}"
end

x = 42
debug(x)  # Prints: DEBUG: x = 42

# Macro to define methods
macro define_property(name, type)
  @{{name}} : {{type}}

  def {{name}}
    @{{name}}
  end

  def {{name}}=(value : {{type}})
    @{{name}} = value
  end
end

class Person
  define_property name, String
  define_property age, Int32

  def initialize(@name, @age)
  end
end

# Generate multiple methods
macro define_methods
  {% for name in %w(foo bar baz) %}
    def {{name.id}}
      "{{name}}"
    end
  {% end %}
end

class MyClass
  define_methods
end

obj = MyClass.new
puts obj.foo  # "foo"
```

----------

## Type Checking
Working with union types and type checking
```crystal
# Union types
value : Int32 | String = 42
value = "hello"

# Check type
if value.is_a?(String)
  puts value.upcase
elsif value.is_a?(Int32)
  puts value * 2
end

# Nilable types
name : String? = nil
name = "Alice"

# Safe navigation
length = name.try(&.size)  # nil if name is nil

# Type assertion
str = value.as(String)  # Raises if not String

# Type cast
num = value.as?(Int32)  # Returns nil if not Int32

# Responds to
if value.responds_to?(:upcase)
  puts value.upcase
end

# Typeof
x = 42
puts typeof(x)  # Int32
```

----------

## Generics
Generic classes and methods
```crystal
# Generic class
class Box(T)
  @value : T

  def initialize(@value)
  end

  def get : T
    @value
  end

  def set(@value : T)
  end
end

int_box = Box(Int32).new(42)
str_box = Box(String).new("hello")

# Generic method
def first(array : Array(T)) : T forall T
  array[0]
end

puts first([1, 2, 3])        # 1
puts first(["a", "b", "c"])  # "a"

# Multiple type parameters
def pair(a : T, b : U) : {T, U} forall T, U
  {a, b}
end

result = pair(1, "hello")  # {1, "hello"}
```

----------

## C Bindings
Calling C functions
```crystal
# Link with math library
@[Link("m")]
lib LibM
  fun sqrt(x : Float64) : Float64
  fun pow(x : Float64, y : Float64) : Float64
  fun sin(x : Float64) : Float64
  fun cos(x : Float64) : Float64
end

puts LibM.sqrt(16.0)      # 4.0
puts LibM.pow(2.0, 3.0)   # 8.0
puts LibM.sin(0.0)        # 0.0

# Calling system functions
lib LibC
  fun getpid : Int32
  fun time(t : Int64*) : Int64
end

puts "Process ID: #{LibC.getpid}"
puts "Unix timestamp: #{LibC.time(nil)}"
```
