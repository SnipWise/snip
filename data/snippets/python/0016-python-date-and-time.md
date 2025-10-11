<!-- METADATA
{
  "title": "Python Date And Time",
  "tags": [
    "python",
    "time"
  ],
  "language": "python"
}
-->

## Date and Time
Working with datetime module
```python
from datetime import datetime, date, time, timedelta
import time as time_module

# Current date and time
now = datetime.now()
today = date.today()
current_time = datetime.now().time()

print(f"Now: {now}")
print(f"Today: {today}")
print(f"Time: {current_time}")

# Formatting
formatted = now.strftime("%Y-%m-%d %H:%M:%S")
formatted_date = today.strftime("%B %d, %Y")

print(f"Formatted datetime: {formatted}")
print(f"Formatted date: {formatted_date}")

# Date arithmetic
tomorrow = today + timedelta(days=1)
next_week = now + timedelta(weeks=1)
one_hour_ago = now - timedelta(hours=1)

print(f"Tomorrow: {tomorrow}")
print(f"One hour ago: {one_hour_ago}")

# Parse date string
date_string = "2023-12-25 10:30:00"
parsed_date = datetime.strptime(date_string, "%Y-%m-%d %H:%M:%S")
print(f"Parsed: {parsed_date}")

# Timestamp
timestamp = now.timestamp()
from_timestamp = datetime.fromtimestamp(timestamp)
print(f"Timestamp: {timestamp}")
print(f"From timestamp: {from_timestamp}")
```