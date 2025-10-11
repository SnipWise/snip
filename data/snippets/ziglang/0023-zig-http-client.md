<!-- METADATA
{
  "title": "Zig HTTP Client",
  "tags": [
    "zig",
    "http",
    "networking"
  ],
  "language": "zig"
}
-->

## HTTP Client
Making HTTP requests
```zig
const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // Create HTTP client
    var client = std.http.Client{ .allocator = allocator };
    defer client.deinit();

    // Prepare request
    const uri = try std.Uri.parse("https://api.github.com/users/ziglang");
    var header_buffer: [8192]u8 = undefined;
    var request = try client.open(.GET, uri, .{ .server_header_buffer = &header_buffer });
    defer request.deinit();

    // Send request
    try request.send();
    try request.wait();

    // Read response
    var response_buffer: [4096]u8 = undefined;
    const bytes_read = try request.readAll(&response_buffer);

    std.debug.print("Response: {s}\n", .{response_buffer[0..bytes_read]});
}
```
