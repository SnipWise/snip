<!-- METADATA
{
  "title": "Python Generators And Iterators",
  "tags": [
    "python",
    "generators",
    "iterators"
  ],
  "language": "python"
}
-->

## Generators and Iterators
Creating generators and custom iterators
```python
# Generator function
def fibonacci(n):
    a, b = 0, 1
    for _ in range(n):
        yield a
        a, b = b, a + b

# Generator expression
squares = (x**2 for x in range(10))

# Custom iterator class
class CountDown:
    def __init__(self, start):
        self.start = start

    def __iter__(self):
        return self

    def __next__(self):
        if self.start <= 0:
            raise StopIteration
        self.start -= 1
        return self.start + 1

# Usage
print("Fibonacci sequence:")
for num in fibonacci(10):
    print(num, end=" ")

print("\nSquares:")
for square in squares:
    if square > 50:
        break
    print(square, end=" ")

print("\nCountdown:")
for num in CountDown(5):
    print(num, end=" ")

# Infinite generator
def infinite_sequence():
    num = 0
    while True:
        yield num
        num += 1

# Take first 5 from infinite sequence
gen = infinite_sequence()
first_five = [next(gen) for _ in range(5)]
print(f"\nFirst five: {first_five}")
```