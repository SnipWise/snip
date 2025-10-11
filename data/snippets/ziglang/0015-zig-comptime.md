<!-- METADATA
{
  "title": "Zig Comptime",
  "tags": [
    "zig",
    "comptime",
    "metaprogramming"
  ],
  "language": "zig"
}
-->

## Comptime
Compile-time execution and metaprogramming
```zig
const std = @import("std");

fn fibonacci(comptime n: u32) u32 {
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

fn generic(comptime T: type, value: T) T {
    return value * 2;
}

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    // Computed at compile time
    const fib10 = comptime fibonacci(10);
    try stdout.print("Fibonacci(10): {d}\n", .{fib10});

    // Generic function
    const x = generic(i32, 21);
    const y = generic(f64, 3.14);
    try stdout.print("Generic i32: {d}, f64: {d}\n", .{x, y});
}
```
