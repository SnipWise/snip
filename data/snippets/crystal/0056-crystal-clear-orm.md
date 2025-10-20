<!-- METADATA
{
  "title": "Crystal Clear ORM",
  "tags": [
    "crystal",
    "orm",
    "database"
  ],
  "language": "crystal"
}
-->

## Clear ORM
Advanced PostgreSQL ORM
```crystal
require "clear"

Clear::SQL.init("postgres://user:pass@localhost/mydb")

# Define model
class User
  include Clear::Model

  primary_key
  column name : String
  column email : String
  column age : Int32

  has_many posts : Post, foreign_key: "user_id"

  def validate
    ensure_than name, "must be present", &.!=("")
    ensure_than email, "must be valid", &.includes?("@")
  end
end

class Post
  include Clear::Model

  primary_key
  column title : String
  column content : String
  column user_id : Int64

  belongs_to user : User
end

# Create
user = User.create!(name: "Alice", email: "alice@example.com")

# Query with SQL builder
users = User.query
  .where { age > 18 }
  .order_by(name: :asc)
  .limit(10)

# Joins
posts_with_users = Post.query
  .join(:user) { user.id == post.user_id }
  .select("posts.*, users.name as user_name")

# Transactions
Clear::SQL.transaction do |tx|
  user = User.create!(name: "Bob")
  Post.create!(title: "First post", user_id: user.id)
end
```
