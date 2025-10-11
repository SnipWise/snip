<!-- METADATA
{
  "title": "Python Loops",
  "tags": [
    "python",
    "loops"
  ],
  "language": "python"
}
-->

## Loops
Different types of loops
```python
# For loop with range
for i in range(5):
    print(f"Number: {i}")

# For loop with list
fruits = ["apple", "banana", "cherry"]
for fruit in fruits:
    print(f"Fruit: {fruit}")

# For loop with enumerate
for index, fruit in enumerate(fruits):
    print(f"{index}: {fruit}")

# While loop
count = 0
while count < 3:
    print(f"Count: {count}")
    count += 1

# Loop with else
for i in range(3):
    print(i)
else:
    print("Loop completed")
```