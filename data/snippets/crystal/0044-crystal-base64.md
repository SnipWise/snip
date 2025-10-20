<!-- METADATA
{
  "title": "Crystal Base64",
  "tags": [
    "crystal",
    "encoding",
    "base64"
  ],
  "language": "crystal"
}
-->

## Base64
Base64 encoding and decoding
```crystal
require "base64"

# Encode string
text = "Hello, World!"
encoded = Base64.encode(text)
puts encoded  # => "SGVsbG8sIFdvcmxkIQ=="

# Decode string
decoded = Base64.decode_string(encoded)
puts decoded  # => "Hello, World!"

# Encode bytes
bytes = "binary data".to_slice
encoded_bytes = Base64.encode(bytes)

# Decode to bytes
decoded_bytes = Base64.decode(encoded_bytes)
puts String.new(decoded_bytes)

# URL-safe encoding
url_encoded = Base64.urlsafe_encode("data+with/special=chars")
url_decoded = Base64.urlsafe_decode(url_encoded)

# Strict encoding (no line breaks)
strict_encoded = Base64.strict_encode("data")

# File encoding
file_data = File.read("image.png")
base64_file = Base64.encode(file_data)

# Decode and write file
File.write("output.png", Base64.decode(base64_file))
```
