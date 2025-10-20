<!-- METADATA
{
  "title": "Crystal TCP Sockets",
  "tags": [
    "crystal",
    "networking",
    "tcp"
  ],
  "language": "crystal"
}
-->

## TCP Sockets
TCP client and server
```crystal
require "socket"

# TCP Server
server = TCPServer.new("localhost", 8080)
puts "Server listening on port 8080"

spawn do
  while client = server.accept?
    spawn handle_client(client)
  end
end

def handle_client(client)
  message = client.gets
  puts "Received: #{message}"
  client.puts "Echo: #{message}"
  client.close
end

# TCP Client
client = TCPSocket.new("localhost", 8080)
client.puts "Hello, Server!"
response = client.gets
puts "Server response: #{response}"
client.close

# Server with multiple connections
def run_server
  server = TCPServer.new(8080)

  loop do
    client = server.accept
    spawn do
      while line = client.gets
        client.puts line.upcase
      end
      client.close
    end
  end
end

# Non-blocking read
socket = TCPSocket.new("example.com", 80)
socket.read_timeout = 5.seconds
```
