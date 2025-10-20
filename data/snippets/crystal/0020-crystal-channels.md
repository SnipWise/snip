<!-- METADATA
{
  "title": "Crystal Channels",
  "tags": [
    "crystal",
    "concurrency",
    "channels"
  ],
  "language": "crystal"
}
-->

## Channels
Communication between fibers
```crystal
# Unbuffered channel
channel = Channel(String).new

# Producer fiber
spawn do
  3.times do |i|
    channel.send("Message #{i}")
    sleep 0.5
  end
  channel.close
end

# Consumer fiber
spawn do
  loop do
    value = channel.receive?
    break unless value
    puts "Received: #{value}"
  end
end

sleep 2.seconds

# Buffered channel
buffered = Channel(Int32).new(5)

spawn do
  10.times { |i| buffered.send(i) }
  buffered.close
end

buffered.each { |val| puts val }

# Select from multiple channels
ch1 = Channel(String).new
ch2 = Channel(Int32).new

select
when msg = ch1.receive
  puts "From ch1: #{msg}"
when num = ch2.receive
  puts "From ch2: #{num}"
end
```
