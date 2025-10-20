<!-- METADATA
{
  "title": "Crystal Custom Iterators",
  "tags": [
    "crystal",
    "iterators",
    "enumerables"
  ],
  "language": "crystal"
}
-->

## Custom Iterators
Creating custom enumerable types
```crystal
# Custom iterable class
class Countdown
  include Enumerable(Int32)

  def initialize(@start : Int32)
  end

  def each(&block : Int32 -> _)
    current = @start
    while current > 0
      yield current
      current -= 1
    end
  end
end

# Usage
countdown = Countdown.new(5)
countdown.each { |n| puts n }

# Now has all Enumerable methods
puts countdown.to_a           # => [5, 4, 3, 2, 1]
puts countdown.sum            # => 15
puts countdown.select(&.even?) # => [4, 2]

# Iterator class
class FibonacciIterator
  include Iterator(Int32)

  def initialize(@limit : Int32)
    @a = 0
    @b = 1
    @count = 0
  end

  def next
    return stop if @count >= @limit
    value = @a
    @a, @b = @b, @a + @b
    @count += 1
    value
  end
end

fib = FibonacciIterator.new(10)
puts fib.to_a  # => [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
```
