<!-- METADATA
{
  "title": "Zig File I/O",
  "tags": [
    "zig",
    "file",
    "io"
  ],
  "language": "zig"
}
-->

## File I/O
Reading and writing files
```zig
const std = @import("std");

pub fn main() !void {
    // Write to file
    const file = try std.fs.cwd().createFile("test.txt", .{});
    defer file.close();

    try file.writeAll("Hello, Zig!\n");

    // Read from file
    const read_file = try std.fs.cwd().openFile("test.txt", .{});
    defer read_file.close();

    var buffer: [100]u8 = undefined;
    const bytes_read = try read_file.readAll(&buffer);

    std.debug.print("Read {d} bytes: {s}", .{bytes_read, buffer[0..bytes_read]});

    // Delete file
    try std.fs.cwd().deleteFile("test.txt");
}
```
