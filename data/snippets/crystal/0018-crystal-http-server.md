<!-- METADATA
{
  "title": "Crystal HTTP Server",
  "tags": [
    "crystal",
    "http",
    "server"
  ],
  "language": "crystal"
}
-->

## HTTP Server
Creating a basic HTTP server
```crystal
require "http/server"

# Simple server
server = HTTP::Server.new do |context|
  context.response.content_type = "text/plain"
  context.response.print "Hello, Crystal!"
end

# Routing
server = HTTP::Server.new do |context|
  case context.request.path
  when "/"
    context.response.print "Home"
  when "/about"
    context.response.print "About"
  when "/api/users"
    context.response.content_type = "application/json"
    context.response.print %({"users": []})
  else
    context.response.status_code = 404
    context.response.print "Not Found"
  end
end

address = server.bind_tcp(8080)
puts "Listening on http://#{address}"
server.listen
```
