<!-- METADATA
{
  "title": "Python Lists And List Comprehensions",
  "tags": [
    "python",
    "io",
    "comprehensions"
  ],
  "language": "python"
}
-->

## Lists and List Comprehensions
Working with lists and comprehensions
```python
numbers = [1, 2, 3, 4, 5]
names = ["Alice", "Bob", "Charlie"]

# List operations
numbers.append(6)
numbers.insert(0, 0)
first = numbers.pop(0)

# List comprehensions
squares = [x**2 for x in numbers]
evens = [x for x in numbers if x % 2 == 0]
pairs = [(x, y) for x in range(3) for y in range(3)]

print(f"Squares: {squares}")
print(f"Evens: {evens}")
```