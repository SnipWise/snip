<!-- METADATA
{
  "title": "Crystal Benchmarking",
  "tags": [
    "crystal",
    "performance",
    "benchmarking"
  ],
  "language": "crystal"
}
-->

## Benchmarking
Performance measurement
```crystal
require "benchmark"

# Basic benchmark
elapsed = Benchmark.measure do
  1_000_000.times { |i| i * 2 }
end
puts "Elapsed: #{elapsed}"

# Compare implementations
Benchmark.ips do |x|
  x.report("map") do
    (1..1000).map { |i| i * 2 }
  end

  x.report("each") do
    result = [] of Int32
    (1..1000).each { |i| result << i * 2 }
  end
end

# Memory benchmark
Benchmark.memory do |x|
  x.report("array") do
    Array.new(10_000) { |i| i }
  end

  x.report("deque") do
    Deque.new(10_000) { |i| i }
  end
end

# Real time measurement
time = Time.measure do
  sleep 1.second
end
puts "Duration: #{time}"

# Multiple runs
Benchmark.bm do |x|
  x.report("string concat") { "hello" + "world" }
  x.report("interpolation") { "hello #{"world"}" }
  x.report("string builder") do
    String.build { |io| io << "hello" << "world" }
  end
end
```
