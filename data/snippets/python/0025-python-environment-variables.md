## Environment Variables
Working with environment variables and configuration
```python
import os

# Basic environment variable operations
debug_mode = os.getenv('DEBUG', 'False').lower() == 'true'
database_url = os.getenv('DATABASE_URL', 'sqlite:///app.db')
port = int(os.getenv('PORT', 5000))

# Required environment variable (raises KeyError if not set)
try:
    api_key = os.environ['API_KEY']
except KeyError:
    api_key = "default-api-key"

# Set environment variable
os.environ['TEMP_VAR'] = 'temporary_value'

# Check if variable exists
if 'HOME' in os.environ:
    print(f"Home: {os.environ['HOME']}")

# Configuration class
class Config:
    DEBUG = debug_mode
    DATABASE_URL = database_url
    PORT = port
    API_KEY = api_key

print(f"Debug mode: {Config.DEBUG}")
print(f"Port: {Config.PORT}")
print(f"Total env vars: {len(os.environ)}")
```