## Testing with unittest
Unit testing framework
```python
import unittest
from unittest.mock import patch, MagicMock

def add(a, b):
    return a + b

def divide(a, b):
    if b == 0:
        raise ValueError("Cannot divide by zero")
    return a / b

class Calculator:
    def multiply(self, a, b):
        return a * b

class TestMathOperations(unittest.TestCase):
    def setUp(self):
        self.calculator = Calculator()

    def test_add(self):
        self.assertEqual(add(2, 3), 5)
        self.assertEqual(add(-1, 1), 0)

    def test_divide(self):
        self.assertEqual(divide(10, 2), 5.0)
        with self.assertRaises(ValueError):
            divide(10, 0)

    def test_multiply(self):
        self.assertEqual(self.calculator.multiply(3, 4), 12)

    @patch('requests.get')
    def test_mocked_request(self, mock_get):
        mock_response = MagicMock()
        mock_response.json.return_value = {"status": "ok"}
        mock_get.return_value = mock_response
        
        # Simulate function that uses requests
        import requests
        response = requests.get("https://api.example.com")
        self.assertEqual(response.json()["status"], "ok")

if __name__ == "__main__":
    unittest.main()
```