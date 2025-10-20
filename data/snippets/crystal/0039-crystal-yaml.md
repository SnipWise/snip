<!-- METADATA
{
  "title": "Crystal YAML",
  "tags": [
    "crystal",
    "yaml",
    "serialization"
  ],
  "language": "crystal"
}
-->

## YAML
YAML serialization and parsing
```crystal
require "yaml"

# Define serializable class
class Config
  include YAML::Serializable

  property host : String
  property port : Int32
  property debug : Bool
  property tags : Array(String)

  def initialize(@host, @port, @debug, @tags)
  end
end

# Serialize to YAML
config = Config.new("localhost", 3000, true, ["web", "api"])
yaml_str = config.to_yaml
puts yaml_str

# Deserialize from YAML
yaml_data = <<-YAML
host: localhost
port: 3000
debug: true
tags:
  - web
  - api
YAML

loaded = Config.from_yaml(yaml_data)
puts loaded.host  # => "localhost"

# Parse raw YAML
data = YAML.parse(yaml_data)
puts data["host"]  # => "localhost"

# Write to file
File.write("config.yaml", config.to_yaml)

# Read from file
loaded_config = Config.from_yaml(File.read("config.yaml"))
```
