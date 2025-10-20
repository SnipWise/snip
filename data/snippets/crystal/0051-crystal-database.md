<!-- METADATA
{
  "title": "Crystal Database",
  "tags": [
    "crystal",
    "database",
    "sql"
  ],
  "language": "crystal"
}
-->

## Database
Database operations with crystal-db
```crystal
require "db"
require "sqlite3"  # or "pg" for PostgreSQL

# Open connection
DB.open "sqlite3://./data.db" do |db|

  # Execute DDL
  db.exec "CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name TEXT,
    age INTEGER
  )"

  # Insert data
  db.exec "INSERT INTO users (name, age) VALUES (?, ?)",
    "Alice", 30

  # Insert with named parameters
  db.exec "INSERT INTO users (name, age) VALUES (:name, :age)",
    name: "Bob", age: 25

  # Query single row
  name = db.query_one "SELECT name FROM users WHERE id = ?",
    1, as: String

  # Query multiple rows
  db.query "SELECT name, age FROM users" do |rs|
    rs.each do
      name = rs.read(String)
      age = rs.read(Int32)
      puts "#{name}: #{age}"
    end
  end

  # Query with mapping
  users = db.query_all "SELECT name, age FROM users",
    as: {String, Int32}

  users.each do |(name, age)|
    puts "#{name}: #{age}"
  end
end

# Connection pool
DB.open "postgres://user:pass@localhost/mydb?max_pool_size=10"
```
