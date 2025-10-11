## HTTP Requests
Making HTTP requests with requests library
```python
import requests
import json

# GET request
response = requests.get("https://httpbin.org/json")
print(f"Status: {response.status_code}")
print(f"JSON: {response.json()}")

# POST request with JSON
data = {"name": "Alice", "age": 30}
headers = {"Content-Type": "application/json"}

response = requests.post(
    "https://httpbin.org/post",
    json=data,
    headers=headers
)

if response.status_code == 200:
    result = response.json()
    print(f"POST response: {result['json']}")

# GET with parameters
params = {"page": 1, "limit": 10}
response = requests.get("https://httpbin.org/get", params=params)
print(f"URL with params: {response.url}")

# Error handling
try:
    response = requests.get("https://httpbin.org/status/404", timeout=5)
    response.raise_for_status()
except requests.exceptions.RequestException as e:
    print(f"Request error: {e}")
```