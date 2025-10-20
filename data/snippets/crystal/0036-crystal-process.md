<!-- METADATA
{
  "title": "Crystal Process",
  "tags": [
    "crystal",
    "system",
    "process"
  ],
  "language": "crystal"
}
-->

## Process
Running external commands
```crystal
# Run command and get output
output = `ls -la`
puts output

# Process.run with capture
result = Process.run("echo", ["Hello"], output: Process::Redirect::Pipe)

# Capture output
process = Process.new("ls", ["-la"], output: Process::Redirect::Pipe)
output = process.output.gets_to_end
status = process.wait
puts "Exit code: #{status.exit_code}"

# Pipe between processes
Process.run("ls", output: Process::Redirect::Pipe) do |proc|
  Process.run("grep", ["test"], input: proc.output)
end

# Run in shell
Process.run("cat file.txt | grep pattern", shell: true)

# With error capture
process = Process.new(
  "command",
  output: Process::Redirect::Pipe,
  error: Process::Redirect::Pipe
)

# Environment variables
Process.run("printenv", env: {"MY_VAR" => "value"})

# Current process info
puts Process.pid
```
