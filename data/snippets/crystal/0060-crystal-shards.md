<!-- METADATA
{
  "title": "Crystal Shards",
  "tags": [
    "crystal",
    "dependencies",
    "shards"
  ],
  "language": "crystal"
}
-->

## Shards
Dependency management
```yaml
# shard.yml - Project configuration
name: myapp
version: 0.1.0

authors:
  - Your Name <you@example.com>

crystal: 1.0.0

license: MIT

dependencies:
  kemal:
    github: kemalcr/kemal
    version: ~> 1.0.0

  db:
    github: crystal-lang/crystal-db

  pg:
    github: will/crystal-pg
    branch: master

development_dependencies:
  spec-kemal:
    github: kemalcr/spec-kemal

targets:
  myapp:
    main: src/myapp.cr

scripts:
  postinstall: echo "Dependencies installed"
```

```crystal
# Install: shards install
# Update: shards update
# Build: shards build

# Require in code
require "kemal"
require "db"
require "pg"
```
