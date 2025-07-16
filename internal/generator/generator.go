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

	// Create default about page with Bazel information
	aboutPath := filepath.Join(name, "pages", "about.html")
	aboutFile, err := os.Create(aboutPath)
	if err != nil {
		return fmt.Errorf("failed to create about page: %w", err)
	}
	defer aboutFile.Close()

	aboutContent := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>About</title>
</head>
<body>
    <h1>About Bazel</h1>
    <p>Welcome to your new static site generated with <strong>Bazel</strong>, a fast and simple static site generator.</p>

    <h2>Getting Started</h2>
    <p>Your site is now set up and ready to use. Here's how to get started:</p>

    <h3>Creating Posts</h3>
    <p>To create a new blog post, run:</p>
    <pre><code>bazel new post "Your Post Title"</code></pre>
    <p>This will create a new Markdown file in the <code>posts/</code> directory and open it in your default editor.</p>

    <h3>Creating Pages</h3>
    <p>To create a new static page, run:</p>
    <pre><code>bazel new page "Your Page Title"</code></pre>
    <p>This will create a new HTML file in the <code>pages/</code> directory.</p>

    <h3>Building Your Site</h3>
    <p>To build your site for production, run:</p>
    <pre><code>bazel build</code></pre>
    <p>This will generate all HTML files and assets in the <code>public/</code> directory.</p>

    <h3>Development Server</h3>
    <p>To preview your site locally while developing, run:</p>
    <pre><code>bazel serve</code></pre>
    <p>This will start a local server at <code>http://localhost:3000</code> where you can preview your site.</p>

    <h3>Configuration</h3>
    <p>You can customize your site by editing the <code>bazel.toml</code> file in your site's root directory. This includes:</p>
    <ul>
        <li><strong>Site title and description</strong></li>
        <li><strong>Theme colors and fonts</strong></li>
        <li><strong>Base URL for deployment</strong></li>
        <li><strong>Social media links</strong></li>
    </ul>

    <h3>Themes</h3>
    <p>Bazel comes with built-in themes including:</p>
    <ul>
        <li><code>pika-beach</code> - A warm, beach-inspired theme (default)</li>
        <li><code>catppuccin-latte</code> - A light, elegant theme</li>
        <li><code>catppuccin-mocha</code> - A dark, modern theme</li>
        <li><code>3li7e</code> - A retro green-on-black CRT monitor theme</li>
    </ul>

    <p>Change the theme in your <code>bazel.toml</code> configuration file.</p>

    <h2>Next Steps</h2>
    <p>Start by creating your first post or customizing this about page. Happy blogging!</p>
</body>
</html>`

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
	// Replace spaces with underscores for filename
	filename := fmt.Sprintf("posts/%s.md", strings.ReplaceAll(title, " ", "_"))
	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("post already exists")
	}

	postFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create post: %w", err)
	}
	defer postFile.Close()

	// Get current date and time
	now := time.Now()
	dateStr := now.Format("January 2, 2006")
	timeStr := now.Format("15:04")

	content := fmt.Sprintf("---\ntitle: %s\ndate: %s\ntime: %s\n---\n\nStart writing here...\n", title, dateStr, timeStr)
	_, err = postFile.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write post content: %w", err)
	}

	return openInEditor(filename)
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
