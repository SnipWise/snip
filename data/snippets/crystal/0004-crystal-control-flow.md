<!-- METADATA
{
  "title": "Crystal Control Flow",
  "tags": [
    "crystal",
    "basics",
    "control-flow"
  ],
  "language": "crystal"
}
-->

## Control Flow
If, unless, case statements
```crystal
# If statement
age = 18
if age >= 18
  puts "Adult"
elsif age >= 13
  puts "Teenager"
else
  puts "Child"
end

# Unless (opposite of if)
unless age < 18
  puts "Can vote"
end

# Case statement
day = "Monday"
case day
when "Monday"
  puts "Start of week"
when "Friday"
  puts "Almost weekend"
else
  puts "Regular day"
end

# Ternary operator
status = age >= 18 ? "adult" : "minor"
```
