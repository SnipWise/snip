<!-- METADATA
{
  "title": "Crystal File I/O",
  "tags": [
    "crystal",
    "io",
    "files"
  ],
  "language": "crystal"
}
-->

## File I/O
Reading and writing files
```crystal
# Write to file
File.write("data.txt", "Hello, Crystal!")

# Read from file
content = File.read("data.txt")
puts content

# Read lines
File.each_line("data.txt") do |line|
  puts line
end

# Open with block (auto-close)
File.open("data.txt", "w") do |file|
  file.puts "Line 1"
  file.puts "Line 2"
end

# Append to file
File.open("data.txt", "a") do |file|
  file.puts "Line 3"
end

# Check existence
if File.exists?("data.txt")
  puts "File exists"
end

# File info
info = File.info("data.txt")
puts info.size
```
