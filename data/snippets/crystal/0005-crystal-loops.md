<!-- METADATA
{
  "title": "Crystal Loops",
  "tags": [
    "crystal",
    "basics",
    "loops"
  ],
  "language": "crystal"
}
-->

## Loops
Various loop constructs
```crystal
# While loop
i = 0
while i < 5
  puts i
  i += 1
end

# Until loop
j = 0
until j >= 5
  puts j
  j += 1
end

# Loop with break
loop do
  puts "Enter 'quit' to exit"
  break if gets == "quit"
end

# Times
5.times do |i|
  puts "Iteration #{i}"
end

# Range iteration
(1..5).each { |n| puts n }

# Array iteration
["a", "b", "c"].each_with_index do |val, idx|
  puts "#{idx}: #{val}"
end
```
