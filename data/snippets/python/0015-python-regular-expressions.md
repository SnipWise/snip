## Regular Expressions
Pattern matching with re module
```python
import re

text = "Contact us at info@company.com or support@company.com. Phone: (555) 123-4567"

# Email pattern
email_pattern = r"[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}"
emails = re.findall(email_pattern, text)
print(f"Found emails: {emails}")

# Phone pattern
phone_pattern = r"\((\d{3})\)\s(\d{3})-(\d{4})"
phone_match = re.search(phone_pattern, text)
if phone_match:
    area_code, prefix, number = phone_match.groups()
    print(f"Phone: ({area_code}) {prefix}-{number}")

# Replace pattern
masked_text = re.sub(email_pattern, "[EMAIL]", text)
print(f"Masked: {masked_text}")

# Split by pattern
parts = re.split(r"[.!?]", "Hello. How are you? I'm fine!")
print(f"Sentences: {[p.strip() for p in parts if p.strip()]}")

# Compile pattern for reuse
email_regex = re.compile(email_pattern)
matches = email_regex.finditer(text)
for match in matches:
    print(f"Email found at position {match.start()}-{match.end()}: {match.group()}")
```