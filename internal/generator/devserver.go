package generator

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// StartDevServer starts a development server with live reload
func StartDevServer() error {
	// First build the site
	if err := BuildSite(); err != nil {
		return fmt.Errorf("failed to build site: %w", err)
	}
	updateModTime()

	// Start file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create file watcher: %w", err)
	}
	defer watcher.Close()

	// Watch for changes in source directories
	watchDirs := []string{"posts", "pages", "themes", "."}
	for _, dir := range watchDirs {
		if _, err := os.Stat(dir); err == nil {
			if err := watcher.Add(dir); err != nil {
				log.Printf("Warning: failed to watch directory %s: %v", dir, err)
			}
		}
	}

	// Explicitly watch the config file
	if _, err := os.Stat("bazel.toml"); err == nil {
		if err := watcher.Add("bazel.toml"); err != nil {
			log.Printf("Warning: failed to watch bazel.toml: %v", err)
		} else {
			log.Println("📝 Watching bazel.toml for configuration changes")
		}
	}

	// Start file system watcher in a goroutine
	go func() {
		debounceTimer := time.NewTimer(0)
		if !debounceTimer.Stop() {
			<-debounceTimer.C
		}

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				// Only rebuild on write events and ignore temporary files
				if event.Op&fsnotify.Write == fsnotify.Write {
					if filepath.Ext(event.Name) == ".tmp" ||
						filepath.Base(event.Name)[0] == '.' {
						continue
					}

					log.Printf("File changed: %s", event.Name)

					// Debounce rebuilds (wait 500ms after last change)
					debounceTimer.Reset(500 * time.Millisecond)
					go func() {
						<-debounceTimer.C
						log.Println("Rebuilding site...")
						if err := BuildSite(); err != nil {
							log.Printf("Build error: %v", err)
						} else {
							log.Println("Site rebuilt successfully")
							updateModTime()
						}
					}()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Watcher error: %v", err)
			}
		}
	}()

	// Set up HTTP server
	publicDir := "public"
	if _, err := os.Stat(publicDir); os.IsNotExist(err) {
		return fmt.Errorf("public directory not found - please build the site first")
	}

	// Create file server with live reload injection
	fs := http.FileServer(&liveReloadFS{http.Dir(publicDir)})
	http.Handle("/", fs)

	// Live reload endpoint
	http.HandleFunc("/live-reload", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Last-Modified", lastModTime.Format(time.RFC1123))
		w.Write([]byte("reload"))
	})

	port := "3000"
	log.Printf("🚀 Development server starting on http://localhost:%s", port)
	log.Println("📁 Serving files from public/")
	log.Println("👀 Watching for file changes...")
	log.Println("🔄 Live reload enabled")
	log.Println("Press Ctrl+C to stop")

	return http.ListenAndServe(":"+port, nil)
}

// liveReloadFS wraps http.FileSystem to inject live reload script
type liveReloadFS struct {
	fs http.FileSystem
}

func (lr *liveReloadFS) Open(name string) (http.File, error) {
	file, err := lr.fs.Open(name)
	if err != nil {
		return nil, err
	}

	// Only inject script into HTML files
	if filepath.Ext(name) == ".html" || name == "/" || name == "/index.html" {
		return &liveReloadFile{File: file}, nil
	}

	return file, nil
}

// liveReloadFile wraps http.File to inject live reload script
type liveReloadFile struct {
	http.File
	modifiedContent []byte
	readOffset      int
}

func (lr *liveReloadFile) Read(p []byte) (n int, err error) {
	// For non-HTML files, just pass through
	if lr.modifiedContent == nil {
		// Read and modify content on first access
		if err := lr.prepareContent(); err != nil {
			return lr.File.Read(p)
		}
	}

	// Read from our modified content
	if lr.readOffset >= len(lr.modifiedContent) {
		return 0, io.EOF
	}

	remaining := len(lr.modifiedContent) - lr.readOffset
	if remaining > len(p) {
		remaining = len(p)
	}

	copy(p, lr.modifiedContent[lr.readOffset:lr.readOffset+remaining])
	lr.readOffset += remaining
	return remaining, nil
}

func (lr *liveReloadFile) prepareContent() error {
	// Read all content first
	var content strings.Builder
	buf := make([]byte, 4096)

	for {
		n, err := lr.File.Read(buf)
		if n > 0 {
			content.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	htmlContent := content.String()

	// Only inject script if it's HTML content
	if strings.Contains(htmlContent, "<html") || strings.Contains(htmlContent, "<HTML") {
		liveReloadScript := `
<script>
// Simple live reload implementation
(function() {
	let lastCheck = Date.now();
	
	function checkForUpdates() {
		fetch('/live-reload?' + Date.now(), {
			method: 'GET',
			cache: 'no-cache'
		}).then(response => response.text())
		.then(data => {
			if (data.includes('reload')) {
				console.log('🔄 Reloading page...');
				window.location.reload();
			}
		}).catch(() => {
			// Server might be restarting, try again
		});
	}
	
	// Check every 1000ms
	setInterval(checkForUpdates, 1000);
})();
</script>
</body>`

		// Replace </body> with script + </body>
		if bodyIndex := strings.LastIndex(htmlContent, "</body>"); bodyIndex != -1 {
			htmlContent = htmlContent[:bodyIndex] + liveReloadScript
		} else if bodyIndex := strings.LastIndex(htmlContent, "</BODY>"); bodyIndex != -1 {
			htmlContent = htmlContent[:bodyIndex] + strings.ReplaceAll(liveReloadScript, "</body>", "</BODY>")
		} else {
			// If no </body> tag, append at the end
			htmlContent += liveReloadScript
		}
	}

	lr.modifiedContent = []byte(htmlContent)
	return nil
}

// Track last modification time for live reload
var lastModTime time.Time

func updateModTime() {
	lastModTime = time.Now()
}
