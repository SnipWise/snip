<!-- METADATA
{
  "title": "Crystal Granite ORM",
  "tags": [
    "crystal",
    "orm",
    "database"
  ],
  "language": "crystal"
}
-->

## Granite ORM
Object-relational mapping
```crystal
require "granite/adapter/sqlite"

Granite::Connections << Granite::Adapter::Sqlite.new(
  name: "sqlite",
  url: "sqlite3://./db/data.db"
)

# Define model
class User < Granite::Base
  connection sqlite
  table users

  column id : Int64, primary: true
  column name : String
  column email : String
  column age : Int32
  column created_at : Time

  has_many :posts

  validates :name, presence: true
  validates :email, uniqueness: true
end

class Post < Granite::Base
  connection sqlite
  table posts

  column id : Int64, primary: true
  column title : String
  column user_id : Int64

  belongs_to :user
end

# Create
user = User.new
user.name = "Alice"
user.email = "alice@example.com"
user.save

# Find
user = User.find!(1)
users = User.all

# Query
adults = User.where(age: 18..100)
alice = User.find_by(name: "Alice")

# Update
user.age = 31
user.save

# Delete
user.destroy

# Associations
posts = user.posts
```
