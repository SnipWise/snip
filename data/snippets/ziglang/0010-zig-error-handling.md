<!-- METADATA
{
  "title": "Zig Error Handling",
  "tags": [
    "zig",
    "errors",
    "error-handling"
  ],
  "language": "zig"
}
-->

## Error Handling
Error sets and error handling
```zig
const std = @import("std");

const MathError = error{
    DivisionByZero,
    Overflow,
};

fn divide(a: f64, b: f64) MathError!f64 {
    if (b == 0) {
        return MathError.DivisionByZero;
    }
    return a / b;
}

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    // Try-catch pattern
    const result = divide(10, 2) catch |err| {
        try stdout.print("Error: {}\n", .{err});
        return;
    };

    try stdout.print("Result: {d}\n", .{result});

    // Or propagate error with try
    const result2 = try divide(20, 4);
    try stdout.print("Result2: {d}\n", .{result2});
}
```
