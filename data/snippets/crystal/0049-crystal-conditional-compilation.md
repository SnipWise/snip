<!-- METADATA
{
  "title": "Crystal Conditional Compilation",
  "tags": [
    "crystal",
    "compilation",
    "flags"
  ],
  "language": "crystal"
}
-->

## Conditional Compilation
Compile-time flags and conditions
```crystal
# Check compilation flag
{% if flag?(:debug) %}
  puts "Debug mode enabled"
{% end %}

# Platform-specific code
{% if flag?(:darwin) %}
  lib LibMac
    # macOS-specific code
  end
{% elsif flag?(:linux) %}
  lib LibLinux
    # Linux-specific code
  end
{% end %}

# Custom flags (compile with -Dproduction)
{% if flag?(:production) %}
  LOG_LEVEL = "error"
{% else %}
  LOG_LEVEL = "debug"
{% end %}

# Environment checks
{% if env("BUILD_ENV") == "test" %}
  ENABLE_MOCKS = true
{% end %}

# Crystal version checks
{% if Crystal::VERSION >= "1.0.0" %}
  # Use newer features
{% end %}

# Multiple conditions
{% if flag?(:linux) && flag?(:production) %}
  OPTIMIZE = true
{% end %}

# Compile-time configuration
macro configure_logger
  {% if flag?(:verbose) %}
    Logger.level = Logger::DEBUG
  {% else %}
    Logger.level = Logger::INFO
  {% end %}
end

# Run: crystal build --release -Dproduction -Dverbose app.cr
```
