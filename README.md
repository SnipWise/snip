# snip
**S.N.I.P.**: Smart Neural Intelligence Partner

```bash
docker pull k33g/snip:0.1.0
```

The Vectore Store is populated at the first startup from the snippets in the `data` folder. 

## Main Commands

```bash
./snip

# Serve mode (the default)
./snip --serve

# Chat mode (interactive CLI)
./snip --chat


# Indexation with 2 arguments (path to file/folder and path to store)
./snip --index ./my-data ./my-store/snippets.json

# Add a file to an existing index
./snip --add-to-index ./file.md ./store/snippets.json
```
