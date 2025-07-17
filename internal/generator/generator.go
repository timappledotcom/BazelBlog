package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/yourusername/bazel_blog/internal/config"
	"github.com/yourusername/bazel_blog/internal/registry"
)

func NewSite(name string) error {
	if err := os.MkdirAll(name, 0755); err != nil {
		return fmt.Errorf("failed to create site directory: %w", err)
	}
	fmt.Println("Site directory created at", filepath.Join(name))

	// Initialize default site structure
	dirs := []string{"posts", "pages", "themes"}
	for _, dir := range dirs {
		path := filepath.Join(name, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create subdirectory %s: %w", dir, err)
		}
	}

	// Create default config
	configPath := filepath.Join(name, "bazel.toml")
	configFile, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer configFile.Close()

	configContent := `site_name = "` + name + `"
title = "Bazel Blog!"
description = "Welcome to Bazel! This is your new static site. Edit this description in bazel.toml to make it your own."
base_url = "https://example.com"
editor = "auto"

[theme]
color_scheme = "pika-beach"
font = "pika-serif"

[socials]`
	_, err = configFile.WriteString(configContent)
	if err != nil {
		return fmt.Errorf("failed to write to config file: %w", err)
	}

	// Create default about page with Bazel information (as Markdown)
	aboutPath := filepath.Join(name, "pages", "about.md")
	aboutFile, err := os.Create(aboutPath)
	if err != nil {
		return fmt.Errorf("failed to create about page: %w", err)
	}
	defer aboutFile.Close()

	aboutContent := `---
title: About Bazel
date: ` + time.Now().Format("January 2, 2006") + `
---

# About Bazel

Welcome to your new static site generated with **Bazel**, a fast and simple static site generator.

## Getting Started

Your site is now set up and ready to use. Here's how to get started:

### Creating Posts

To create a new blog post, run:

` + "```bash" + `
bazel post
` + "```" + `

This will open an interactive menu for creating and managing posts in Markdown format.

### Creating Pages

To create a new static page, run:

` + "```bash" + `
bazel page
` + "```" + `

This will open an interactive menu for creating and managing pages in Markdown format.

### Building Your Site

To build your site for production, run:

` + "```bash" + `
bazel build
` + "```" + `

This will generate all HTML files and assets in the ` + "`public/`" + ` directory with organized structure:

- ` + "`public/posts/`" + ` - Blog posts
- ` + "`public/pages/`" + ` - Static pages
- ` + "`public/`" + ` - Homepage, CSS, and RSS feed

### Development Server

To preview your site locally while developing, run:

` + "```bash" + `
bazel serve
` + "```" + `

This will start a local server at ` + "`http://localhost:3000`" + ` where you can preview your site.

### Configuration

You can customize your site by editing the ` + "`bazel.toml`" + ` file in your site's root directory. This includes:

- **Site title and description**
- **Theme colors and fonts**
- **Base URL for deployment**
- **Social media links**

### Themes

Bazel comes with built-in themes including:

- ` + "`pika-beach`" + ` - A warm, beach-inspired theme (default)
- ` + "`catppuccin-latte`" + ` - A light, elegant theme
- ` + "`catppuccin-mocha`" + ` - A dark, modern theme
- ` + "`3li7e`" + ` - A retro green-on-black CRT monitor theme

Change the theme using the interactive configuration menu or by editing ` + "`bazel.toml`" + `.

## Content Format

Both posts and pages are written in **Markdown** with frontmatter:

` + "```markdown" + `
---
title: Your Page Title
date: January 1, 2025
---

# Your Content Here

Write your content in Markdown format with full support for:

- **Bold** and *italic* text
- [Links](https://example.com)
- Lists and tables
- Code blocks
- And much more!
` + "```" + `

## Next Steps

1. **Customize this about page** by editing ` + "`pages/about.md`" + `
2. **Create your first post** with ` + "`bazel post`" + `
3. **Configure your site** with ` + "`bazel config`" + `
4. **Build and deploy** with ` + "`bazel build`" + `

Happy blogging! 🎉`

	_, err = aboutFile.WriteString(aboutContent)
	if err != nil {
		return fmt.Errorf("failed to write about page content: %w", err)
	}

	// Create sample posts
	samplePosts := []struct {
		Title   string
		Date    string
		Content string
	}{
		{"First Steps", "July 1, 2025", "## Getting Started with Bazel\n\nBazel is a simple and fast static site generator. Start by creating posts!"},
		{"Design Ideas", "June 20, 2025", "## Designing a Great Static Site\n\nConsider theme and layout for your site's content. Use Bazel's options for customization."},
	}

	for _, post := range samplePosts {
		filename := filepath.Join(name, "posts", post.Title+".md")
		postFile, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("failed to create sample post: %w", err)
		}
		defer postFile.Close()

		postContent := fmt.Sprintf("---\ntitle: %s\ndate: %s\n---\n\n%s\n", post.Title, post.Date, post.Content)
		_, err = postFile.WriteString(postContent)
		if err != nil {
			return fmt.Errorf("failed to write sample post content: %w", err)
		}
	}

	// Register the site in the registry
	sitePath, _ := filepath.Abs(name)
	reg, err := registry.LoadRegistry()
	if err != nil {
		return fmt.Errorf("failed to load site registry: %w", err)
	}

	description := "Welcome to Bazel! This is your new static site."
	err = reg.AddSite(name, sitePath, description)
	if err != nil {
		return fmt.Errorf("failed to register site: %w", err)
	}

	fmt.Printf("Site '%s' registered in bazel registry\n", name)
	fmt.Printf("You can now run bazel commands from anywhere!\n")

	return nil
}

func NewPost(title string) error {
	// Validate post title
	if strings.TrimSpace(title) == "" {
		return fmt.Errorf("post title cannot be empty")
	}

	// Validate title length and characters
	if len(title) > 100 {
		return fmt.Errorf("post title too long (max 100 characters)")
	}

	// Check for invalid characters in title
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range invalidChars {
		if strings.Contains(title, char) {
			return fmt.Errorf("post title contains invalid character: %s", char)
		}
	}

	// Create posts directory if it doesn't exist
	postsDir := "posts"
	if err := os.MkdirAll(postsDir, 0755); err != nil {
		return fmt.Errorf("failed to create posts directory: %w", err)
	}

	// Generate filename with better sanitization
	sanitizedTitle := strings.ReplaceAll(title, " ", "_")
	// Remove any remaining problematic characters
	sanitizedTitle = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '_' || r == '-' {
			return r
		}
		return '_'
	}, sanitizedTitle)

	filename := fmt.Sprintf("posts/%s.md", sanitizedTitle)

	// Check if post already exists
	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("post already exists: %s", sanitizedTitle)
	}

	// Create the post file with enhanced error handling
	postFile, err := os.Create(filename)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("permission denied creating post file: %s", filename)
		} else if os.IsExist(err) {
			return fmt.Errorf("post file already exists: %s", filename)
		} else {
			return fmt.Errorf("failed to create post file: %w", err)
		}
	}
	defer postFile.Close()

	// Get current date and time
	now := time.Now()
	dateStr := now.Format("January 2, 2006")
	timeStr := now.Format("15:04")

	// Create post content with proper escaping
	content := fmt.Sprintf("---\ntitle: %s\ndate: %s\ntime: %s\n---\n\nStart writing here...\n", title, dateStr, timeStr)

	// Write content with error handling
	_, err = postFile.WriteString(content)
	if err != nil {
		// Clean up the file if writing fails
		postFile.Close()
		os.Remove(filename)
		return fmt.Errorf("failed to write post content: %w", err)
	}

	// Ensure content is written to disk
	if err := postFile.Sync(); err != nil {
		return fmt.Errorf("failed to save post content: %w", err)
	}

	// Close file before opening in editor
	postFile.Close()

	// Open in editor with enhanced error handling
	err = openInEditor(filename)
	if err != nil {
		// Don't delete the file if editor fails - user can still access it
		return fmt.Errorf("post created successfully but failed to open editor: %w", err)
	}

	return nil
}

func NewPage(title string) error {
	// Replace spaces with underscores for filename
	filename := fmt.Sprintf("pages/%s.md", strings.ReplaceAll(title, " ", "_"))
	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("page already exists")
	}

	pageFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create page: %w", err)
	}
	defer pageFile.Close()

	// Get current date
	now := time.Now()
	dateStr := now.Format("January 2, 2006")

	content := fmt.Sprintf("---\ntitle: %s\ndate: %s\n---\n\n# %s\n\nStart writing here...\n", title, dateStr, title)
	_, err = pageFile.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write page content: %w", err)
	}

	return openInEditor(filename)
}

func ListPosts() ([]string, error) {
	postsDir := "posts"
	if _, err := os.Stat(postsDir); os.IsNotExist(err) {
		return []string{}, nil // No posts directory
	}

	files, err := os.ReadDir(postsDir)
	if err != nil {
		return nil, err
	}

	var posts []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			postName := strings.TrimSuffix(file.Name(), ".md")
			posts = append(posts, postName)
		}
	}

	return posts, nil
}

func ListPages() ([]string, error) {
	pagesDir := "pages"
	if _, err := os.Stat(pagesDir); os.IsNotExist(err) {
		return []string{}, nil // No pages directory
	}

	files, err := os.ReadDir(pagesDir)
	if err != nil {
		return nil, err
	}

	var pages []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			pageName := strings.TrimSuffix(file.Name(), ".md")
			pages = append(pages, pageName)
		} else if strings.HasSuffix(file.Name(), ".html") {
			pageName := strings.TrimSuffix(file.Name(), ".html")
			pages = append(pages, pageName)
		}
	}

	return pages, nil
}

func EditPost(title string) error {
	// Check for .md file first
	filename := fmt.Sprintf("posts/%s.md", title)
	if _, err := os.Stat(filename); err == nil {
		return openInEditor(filename)
	}

	return fmt.Errorf("post not found: %s", title)
}

func EditPage(title string) error {
	// Check for .md file first
	filename := fmt.Sprintf("pages/%s.md", title)
	if _, err := os.Stat(filename); err == nil {
		return openInEditor(filename)
	}

	// Check for .html file
	filename = fmt.Sprintf("pages/%s.html", title)
	if _, err := os.Stat(filename); err == nil {
		return openInEditor(filename)
	}

	return fmt.Errorf("page not found: %s", title)
}

func openInEditor(filename string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		// Fallback to environment variable if config loading fails
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vi"
		}
		cmd := exec.Command(editor, filename)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	editor := cfg.GetEditor()
	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func DeletePost(title string) error {
	// Check for .md file first
	filename := fmt.Sprintf("posts/%s.md", title)
	if _, err := os.Stat(filename); err == nil {
		return os.Remove(filename)
	}

	return fmt.Errorf("post not found: %s", title)
}

func DeletePage(title string) error {
	// Check for .md file first
	filename := fmt.Sprintf("pages/%s.md", title)
	if _, err := os.Stat(filename); err == nil {
		return os.Remove(filename)
	}

	// Check for .html file
	filename = fmt.Sprintf("pages/%s.html", title)
	if _, err := os.Stat(filename); err == nil {
		return os.Remove(filename)
	}

	return fmt.Errorf("page not found: %s", title)
}
