<!-- METADATA
{
  "title": "Crystal WebSockets",
  "tags": [
    "crystal",
    "websockets",
    "networking"
  ],
  "language": "crystal"
}
-->

## WebSockets
WebSocket server and client
```crystal
require "http/web_socket"

# WebSocket Server
server = HTTP::Server.new do |context|
  ws = HTTP::WebSocket.new(context)

  ws.on_message do |message|
    puts "Received: #{message}"
    ws.send "Echo: #{message}"
  end

  ws.on_close do
    puts "Client disconnected"
  end

  ws.run
end

address = server.bind_tcp(8080)
puts "WebSocket server on ws://#{address}"
spawn { server.listen }

# WebSocket Client
HTTP::WebSocket.new("ws://localhost:8080") do |ws|
  ws.send("Hello, WebSocket!")

  ws.on_message do |message|
    puts "Server says: #{message}"
  end

  sleep 1
end

# Broadcast to multiple clients
clients = [] of HTTP::WebSocket

server = HTTP::Server.new do |context|
  ws = HTTP::WebSocket.new(context)
  clients << ws

  ws.on_message do |msg|
    clients.each { |client| client.send(msg) }
  end

  ws.run
end
```
