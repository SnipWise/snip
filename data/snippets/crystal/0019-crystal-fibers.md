<!-- METADATA
{
  "title": "Crystal Fibers",
  "tags": [
    "crystal",
    "concurrency",
    "fibers"
  ],
  "language": "crystal"
}
-->

## Fibers
Lightweight concurrency with fibers
```crystal
# Spawn a fiber
spawn do
  puts "Inside fiber"
  sleep 1.second
  puts "Fiber done"
end

puts "Main continues"
sleep 2.seconds

# Multiple fibers
10.times do |i|
  spawn do
    sleep rand(0.1..1.0)
    puts "Fiber #{i} completed"
  end
end

sleep 2.seconds

# Channel communication
channel = Channel(Int32).new

spawn do
  5.times do |i|
    channel.send(i)
    sleep 0.1
  end
  channel.close
end

while value = channel.receive?
  puts "Received: #{value}"
end
```
