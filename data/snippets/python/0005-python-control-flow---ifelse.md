<!-- METADATA
{
  "title": "Python Control Flow   Ifelse",
  "tags": [
    "python",
    "control-flow"
  ],
  "language": "python"
}
-->

## Control Flow - If/Else
Conditional statements
```python
age = 25
score = 85

if age >= 18:
    print("Adult")
elif age >= 13:
    print("Teenager")
else:
    print("Child")

# Ternary operator
status = "adult" if age >= 18 else "minor"

# Multiple conditions
if 80 <= score <= 100:
    grade = "A"
elif 70 <= score < 80:
    grade = "B"
else:
    grade = "C"

print(f"Status: {status}, Grade: {grade}")
```