<!-- METADATA
{
  "title": "Crystal Exception Handling",
  "tags": [
    "crystal",
    "exceptions",
    "error-handling"
  ],
  "language": "crystal"
}
-->

## Exception Handling
Error handling with rescue
```crystal
# Basic exception handling
begin
  result = 10 / 0
rescue DivisionByZeroError
  puts "Cannot divide by zero"
end

# Multiple rescue clauses
begin
  # Some risky operation
  raise ArgumentError.new("Invalid argument")
rescue DivisionByZeroError
  puts "Division error"
rescue ArgumentError => e
  puts "Argument error: #{e.message}"
rescue Exception => e
  puts "Other error: #{e}"
end

# Ensure (always executed)
begin
  file = File.open("data.txt")
  # Process file
ensure
  file.close if file
end

# Custom exception
class CustomError < Exception
end

raise CustomError.new("Something went wrong")
```
