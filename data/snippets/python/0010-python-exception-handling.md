<!-- METADATA
{
  "title": "Python Exception Handling",
  "tags": [
    "python",
    "error-handling",
    "io"
  ],
  "language": "python"
}
-->

## Exception Handling
Try-catch blocks and custom exceptions
```python
class CustomError(Exception):
    def __init__(self, message):
        self.message = message
        super().__init__(self.message)

def divide_numbers(a, b):
    if b == 0:
        raise CustomError("Cannot divide by zero")
    return a / b

def safe_conversion(value):
    try:
        result = int(value)
        return result
    except ValueError:
        print(f"Cannot convert '{value}' to integer")
        return None
    except Exception as e:
        print(f"Unexpected error: {e}")
        return None
    finally:
        print("Conversion attempt completed")

# Usage
try:
    result = divide_numbers(10, 2)
    print(f"Result: {result}")
    
    bad_result = divide_numbers(10, 0)
except CustomError as e:
    print(f"Custom error: {e.message}")

number = safe_conversion("123")
invalid = safe_conversion("abc")
```