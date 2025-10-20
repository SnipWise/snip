<!-- METADATA
{
  "title": "Crystal Crypto Hashing",
  "tags": [
    "crystal",
    "crypto",
    "security"
  ],
  "language": "crystal"
}
-->

## Crypto Hashing
Cryptographic hash functions
```crystal
require "crypto/md5"
require "crypto/sha1"
require "crypto/sha256"
require "crypto/bcrypt"

# MD5 (not recommended for security)
md5 = Crypto::MD5.hexdigest("Hello, World!")
puts md5

# SHA1
sha1 = Crypto::SHA1.hexdigest("Hello, World!")
puts sha1

# SHA256
sha256 = Crypto::SHA256.hexdigest("Hello, World!")
puts sha256

# SHA512
require "crypto/sha512"
sha512 = Crypto::SHA512.hexdigest("password")

# Bcrypt for password hashing
password = "my_secure_password"
hashed = Crypto::Bcrypt::Password.create(password)
puts hashed

# Verify password
if Crypto::Bcrypt::Password.new(hashed).verify(password)
  puts "Password correct!"
end

# HMAC
require "crypto/hmac"
key = "secret_key"
message = "data to authenticate"
hmac = Crypto::HMAC.hexdigest(:sha256, key, message)
puts hmac
```
