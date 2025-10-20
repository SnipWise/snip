<!-- METADATA
{
  "title": "Crystal Kemal Web Framework",
  "tags": [
    "crystal",
    "web",
    "kemal"
  ],
  "language": "crystal"
}
-->

## Kemal Web Framework
Building web apps with Kemal
```crystal
require "kemal"

# Basic route
get "/" do
  "Hello, Kemal!"
end

# Route with parameters
get "/users/:id" do |env|
  id = env.params.url["id"]
  "User ID: #{id}"
end

# Query parameters
get "/search" do |env|
  query = env.params.query["q"]?
  "Searching for: #{query}"
end

# POST request
post "/users" do |env|
  name = env.params.json["name"].as(String)
  {id: 1, name: name}.to_json
end

# JSON response
get "/api/users" do |env|
  env.response.content_type = "application/json"
  [{id: 1, name: "Alice"}].to_json
end

# Before filter
before_get "/admin/*" do |env|
  halt env, status_code: 401 unless authorized?(env)
end

# Static files
serve_static({"gzip" => true})

# Error handling
error 404 do
  "Page not found"
end

# Start server
Kemal.run
```
