<!-- METADATA
{
  "title": "Crystal UDP Sockets",
  "tags": [
    "crystal",
    "networking",
    "udp"
  ],
  "language": "crystal"
}
-->

## UDP Sockets
UDP client and server
```crystal
require "socket"

# UDP Server
server = UDPSocket.new
server.bind("localhost", 9000)

spawn do
  loop do
    message, client_addr = server.receive
    puts "Received from #{client_addr}: #{message}"
    server.send("ACK", client_addr)
  end
end

# UDP Client
client = UDPSocket.new
client.connect("localhost", 9000)
client.send("Hello, UDP!")

# Receive response
response, addr = client.receive
puts "Response: #{response}"

client.close
server.close

# Broadcast
socket = UDPSocket.new
socket.setsockopt(Socket::SOL_SOCKET, Socket::SO_BROADCAST, 1)
socket.send("Broadcast message", Socket::IPAddress.new("255.255.255.255", 9000))

# Multicast
multicast_addr = "239.0.0.1"
socket = UDPSocket.new
socket.join_group(Socket::IPAddress.new(multicast_addr, 0))
```
