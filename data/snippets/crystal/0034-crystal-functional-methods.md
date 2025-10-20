<!-- METADATA
{
  "title": "Crystal Functional Methods",
  "tags": [
    "crystal",
    "functional",
    "collections"
  ],
  "language": "crystal"
}
-->

## Functional Methods
Higher-order functions for collections
```crystal
numbers = [1, 2, 3, 4, 5]

# Map - transform elements
squares = numbers.map { |n| n ** 2 }
# => [1, 4, 9, 16, 25]

# Select/Filter - keep matching elements
evens = numbers.select { |n| n.even? }
# => [2, 4]

# Reject - remove matching elements
odds = numbers.reject { |n| n.even? }
# => [1, 3, 5]

# Reduce/Inject - accumulate
sum = numbers.reduce(0) { |acc, n| acc + n }
# => 15

product = numbers.reduce { |acc, n| acc * n }
# => 120

# Partition - split by condition
evens, odds = numbers.partition(&.even?)

# Any? / All? / None?
has_even = numbers.any?(&.even?)      # => true
all_positive = numbers.all?(&.> 0)    # => true
has_negative = numbers.none?(&.< 0)   # => true

# Find
first_even = numbers.find { |n| n.even? }  # => 2

# Take/Drop
first_three = numbers.take(3)  # => [1, 2, 3]
skip_two = numbers.skip(2)     # => [3, 4, 5]

# Zip
letters = ['a', 'b', 'c']
zipped = numbers.zip(letters)
# => [{1, 'a'}, {2, 'b'}, {3, 'c'}]
```
