<!-- METADATA
{
  "title": "Zig Threads",
  "tags": [
    "zig",
    "threads",
    "concurrency"
  ],
  "language": "zig"
}
-->

## Threads
Working with threads for concurrency
```zig
const std = @import("std");

fn worker(id: usize) void {
    std.debug.print("Thread {d} started\n", .{id});
    std.time.sleep(1 * std.time.ns_per_s);
    std.debug.print("Thread {d} finished\n", .{id});
}

pub fn main() !void {
    var threads: [3]std.Thread = undefined;

    // Spawn threads
    for (&threads, 0..) |*thread, i| {
        thread.* = try std.Thread.spawn(.{}, worker, .{i});
    }

    // Wait for all threads to complete
    for (threads) |thread| {
        thread.join();
    }

    std.debug.print("All threads completed\n", .{});
}
```
