<!-- METADATA
{
  "title": "Zig C Interop",
  "tags": [
    "zig",
    "c",
    "interop"
  ],
  "language": "zig"
}
-->

## C Interop
Calling C functions from Zig
```zig
const std = @import("std");
const c = @cImport({
    @cInclude("stdio.h");
    @cInclude("stdlib.h");
});

pub fn main() !void {
    // Call C functions
    _ = c.printf("Hello from C printf!\n");

    // Allocate with C malloc
    const ptr = c.malloc(100);
    defer c.free(ptr);

    // Use C types
    var x: c.int = 42;
    _ = c.printf("C int: %d\n", x);

    // String conversion
    const zig_str = "Hello from Zig";
    _ = c.printf("String: %s\n", zig_str.ptr);
}
```
