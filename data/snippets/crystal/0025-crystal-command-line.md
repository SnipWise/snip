<!-- METADATA
{
  "title": "Crystal Command Line",
  "tags": [
    "crystal",
    "cli",
    "argv"
  ],
  "language": "crystal"
}
-->

## Command Line
Command line arguments and options
```crystal
# Access arguments
ARGV.each do |arg|
  puts "Argument: #{arg}"
end

# Simple option parsing
require "option_parser"

name = "World"
verbose = false

OptionParser.parse do |parser|
  parser.banner = "Usage: program [options]"

  parser.on("-n NAME", "--name=NAME", "Specify name") do |n|
    name = n
  end

  parser.on("-v", "--verbose", "Enable verbose mode") do
    verbose = true
  end

  parser.on("-h", "--help", "Show help") do
    puts parser
    exit
  end
end

puts "Hello, #{name}!"
puts "Verbose mode" if verbose

# Environment variables
home = ENV["HOME"]?
puts "Home: #{home}" if home
```
