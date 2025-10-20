<!-- METADATA
{
  "title": "Crystal CLI Application",
  "tags": [
    "crystal",
    "cli",
    "application"
  ],
  "language": "crystal"
}
-->

## CLI Application
Building command-line apps
```crystal
require "option_parser"

# CLI with subcommands
class CLI
  property verbose = false
  property command : String?

  def parse_args
    OptionParser.parse do |parser|
      parser.banner = "Usage: myapp [command] [options]"

      parser.on("init", "Initialize project") do
        @command = "init"
        parser.banner = "Usage: myapp init [options]"
      end

      parser.on("build", "Build project") do
        @command = "build"
      end

      parser.on("-v", "--verbose", "Verbose output") do
        @verbose = true
      end

      parser.on("-h", "--help", "Show help") do
        puts parser
        exit
      end
    end
  end

  def run
    case @command
    when "init"
      init_project
    when "build"
      build_project
    else
      puts "Unknown command"
      exit 1
    end
  end

  private def init_project
    puts "Initializing..." if @verbose
    puts "Project initialized!"
  end

  private def build_project
    puts "Building..."
  end
end

cli = CLI.new
cli.parse_args
cli.run
```
