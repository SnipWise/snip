package watcher

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"snip/embeddings"
	"snip/rag"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/fsnotify/fsnotify"
)

// StartFileWatcher watches the snippets folder and generates embeddings for new .md files
func StartFileWatcher(ctx context.Context, snippetsFolder string, embedder ai.Embedder, store *rag.MemoryVectorStore, genkitInstance any) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}

	// Start watching in a goroutine
	go func() {
		defer watcher.Close()

		for {
			select {
			case <-ctx.Done():
				log.Println("🛑 File watcher stopped")
				return

			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				// Only process Create and Write events for .md files
				if event.Op&(fsnotify.Create|fsnotify.Write) != 0 {
					if strings.HasSuffix(event.Name, ".md") {
						log.Printf("📝 Detected new/modified file: %s", event.Name)

						// Generate embeddings for the new file
						if err := embeddings.GenerateForFile(ctx, event.Name, embedder, store, genkitInstance); err != nil {
							log.Printf("😡 Error generating embeddings for %s: %v", event.Name, err)
						} else {
							log.Printf("✅ Successfully processed file: %s", event.Name)
						}
					}
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("⚠️  Watcher error: %v", err)
			}
		}
	}()

	// Add the snippets folder to watch
	if err := watcher.Add(snippetsFolder); err != nil {
		return fmt.Errorf("failed to watch folder %s: %w", snippetsFolder, err)
	}

	// Also watch subdirectories
	err = filepath.Walk(snippetsFolder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Add directories to watcher
		if info != nil && info.IsDir() {
			if err := watcher.Add(path); err != nil {
				log.Printf("⚠️  Failed to watch subdirectory %s: %v", path, err)
			} else {
				log.Printf("👀 Watching directory: %s", path)
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk directory tree: %w", err)
	}

	log.Printf("👀 File watcher started for: %s", snippetsFolder)
	return nil
}
