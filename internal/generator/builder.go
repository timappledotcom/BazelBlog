package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/adrg/frontmatter"
	"github.com/araddon/dateparse"
	"github.com/yourusername/bazel_blog/internal/config"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
)

type Post struct {
	Title    string
	Date     time.Time
	Content  string
	Filename string
	URL      string
}

type Page struct {
	Title    string
	Content  string
	Filename string
	URL      string
}

// PostMatter represents the frontmatter structure for posts
type PostMatter struct {
	Title string `yaml:"title"`
	Date  string `yaml:"date"`
}

// PageMatter represents the frontmatter structure for pages
type PageMatter struct {
	Title string `yaml:"title"`
}

type Site struct {
	Config *config.Config
	Posts  []Post
	Pages  []Page
}

func BuildSite() error {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Create output directory structure
	outputDir := "public"
	if err := os.RemoveAll(outputDir); err != nil {
		return fmt.Errorf("failed to remove output directory: %w", err)
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create subdirectories
	if err := os.MkdirAll(filepath.Join(outputDir, "posts"), 0755); err != nil {
		return fmt.Errorf("failed to create posts directory: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(outputDir, "pages"), 0755); err != nil {
		return fmt.Errorf("failed to create pages directory: %w", err)
	}

	// Build site structure
	site := &Site{
		Config: cfg,
		Posts:  []Post{},
		Pages:  []Page{},
	}

	// Load posts
	if err := site.loadPosts(); err != nil {
		return fmt.Errorf("failed to load posts: %w", err)
	}

	// Load pages
	if err := site.loadPages(); err != nil {
		return fmt.Errorf("failed to load pages: %w", err)
	}

	// Generate CSS
	if err := site.generateCSS(); err != nil {
		return fmt.Errorf("failed to generate CSS: %w", err)
	}

	// Generate index page
	if err := site.generateIndex(); err != nil {
		return fmt.Errorf("failed to generate index: %w", err)
	}

	// Generate post pages
	if err := site.generatePosts(); err != nil {
		return fmt.Errorf("failed to generate posts: %w", err)
	}

	// Generate regular pages
	if err := site.generatePages(); err != nil {
		return fmt.Errorf("failed to generate pages: %w", err)
	}

	// Generate RSS feed
	if err := site.generateRSS(); err != nil {
		return fmt.Errorf("failed to generate RSS feed: %w", err)
	}

	return nil
}

func (s *Site) loadPosts() error {
	postsDir := "posts"
	if _, err := os.Stat(postsDir); os.IsNotExist(err) {
		return nil // No posts directory
	}

	files, err := ioutil.ReadDir(postsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			content, err := ioutil.ReadFile(filepath.Join(postsDir, file.Name()))
			if err != nil {
				continue
			}

			// Parse frontmatter and content using the frontmatter library
			var matter PostMatter
			rest, err := frontmatter.Parse(strings.NewReader(string(content)), &matter)
			if err != nil {
				// If frontmatter parsing fails, use defaults
				matter.Title = strings.TrimSpace(strings.Replace(file.Name(), ".md", "", 1))
				matter.Date = ""
				rest = content
			}

			// Use title from frontmatter or fallback to filename
			title := matter.Title
			if title == "" {
				title = strings.TrimSpace(strings.Replace(file.Name(), ".md", "", 1))
			}

			// Parse date with flexible parsing
			var postDate time.Time
			if matter.Date != "" {
				if parsedDate, err := dateparse.ParseAny(matter.Date); err == nil {
					postDate = parsedDate
				} else {
					// Fallback to file modification time
					postDate = file.ModTime()
				}
			} else {
				postDate = file.ModTime()
			}

			// Convert markdown to HTML using enhanced Goldmark
			htmlContent := s.markdownToHTML(strings.TrimSpace(string(rest)))
			postURL := "posts/" + strings.Replace(file.Name(), ".md", ".html", 1)

			post := Post{
				Title:    title,
				Date:     postDate,
				Content:  htmlContent,
				Filename: file.Name(),
				URL:      postURL,
			}

			s.Posts = append(s.Posts, post)
		}
	}

	// Sort posts by date (newest first)
	sort.Slice(s.Posts, func(i, j int) bool {
		return s.Posts[i].Date.After(s.Posts[j].Date)
	})

	return nil
}

func (s *Site) loadPages() error {
	pagesDir := "pages"
	if _, err := os.Stat(pagesDir); os.IsNotExist(err) {
		return nil // No pages directory
	}

	files, err := ioutil.ReadDir(pagesDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		// Handle Markdown pages
		if strings.HasSuffix(file.Name(), ".md") {
			content, err := ioutil.ReadFile(filepath.Join(pagesDir, file.Name()))
			if err != nil {
				continue
			}

			// Parse frontmatter and content using the frontmatter library
			var matter PageMatter
			rest, err := frontmatter.Parse(strings.NewReader(string(content)), &matter)
			if err != nil {
				// If frontmatter parsing fails, use defaults
				matter.Title = strings.TrimSpace(strings.Replace(file.Name(), ".md", "", 1))
				rest = content
			}

			// Use title from frontmatter or fallback to filename
			title := matter.Title
			if title == "" {
				title = strings.TrimSpace(strings.Replace(file.Name(), ".md", "", 1))
			}

			// Convert markdown to HTML using enhanced Goldmark
			htmlContent := s.markdownToHTML(strings.TrimSpace(string(rest)))
			pageURL := "pages/" + strings.Replace(file.Name(), ".md", ".html", 1)

			page := Page{
				Title:    title,
				Content:  htmlContent,
				Filename: file.Name(),
				URL:      pageURL,
			}

			s.Pages = append(s.Pages, page)

		} else if strings.HasSuffix(file.Name(), ".html") {
			// Handle existing HTML pages (for backward compatibility)
			content, err := ioutil.ReadFile(filepath.Join(pagesDir, file.Name()))
			if err != nil {
				continue
			}

			title := strings.TrimSpace(strings.Replace(file.Name(), ".html", "", 1))
			pageURL := "pages/" + file.Name()

			page := Page{
				Title:    title,
				Content:  string(content),
				Filename: file.Name(),
				URL:      pageURL,
			}

			s.Pages = append(s.Pages, page)
		}
	}

	return nil
}

func (s *Site) generateCSS() error {
	cssContent := `
:root {
	` + s.Config.GetCSSVariables() + `
	--base-font-size: 18px;
	--base-line-height: 1.5;
	--width-M: 700px;
	--space-XS: 0.5rem;
	--space-S: 0.75rem;
	--space-M: 1rem;
	--space-L: 1.5rem;
	--space-XL: 2rem;
	--space-2XL: 2.5rem;
	--space-3XL: 3rem;
	--space-4XL: 4rem;
	--border-radius: 7px;
	--color-bg: rgba(var(--bg-color), 1);
	--color-txt: rgba(var(--text-color), 1);
	--color-txt-light: rgba(var(--text-color), 0.65);
	--color-primary: var(--accent-color);
	--color-border: rgba(var(--text-color), 0.2);
	--color-bg-light: rgba(var(--text-color), 0.05);
}

* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
	font-family: inherit;
	font-size: inherit;
	overflow-wrap: break-word;
	vertical-align: baseline;
}

html {
	font-size: var(--base-font-size);
	height: 100%;
	overflow: auto;
	scroll-behavior: smooth;
}

body {
	background-color: var(--color-bg);
	color: var(--color-txt);
	display: flex;
	flex-direction: column;
	font-family: var(--font-family);
	font-size: 1rem;
	line-height: var(--base-line-height);
	margin: auto;
	min-height: 100%;
	padding: 0;
	text-wrap: pretty;
}

.site-view {
	font-family: var(--font-family);
	font-size: var(--base-font-size);
	max-width: var(--width-M);
	padding-inline: var(--space-M);
	width: 100%;
	margin: 0 auto;
}

a {
	color: var(--color-primary);
	cursor: pointer;
	display: inline-block;
	text-decoration: underline;
	text-decoration-thickness: 1px;
	text-underline-offset: 2px;
}

h1, h2, h3, h4, h5, h6 {
	font-family: var(--font-family);
	font-weight: 700;
}

h1 {
	font-size: 2em;
	letter-spacing: -0.33px;
	line-height: calc(var(--base-line-height) * 0.8);
	margin-block: 0.8em;
}

h2 {
	font-size: 1.5em;
	line-height: calc(var(--base-line-height) * 0.9);
	margin-block: 1.5em 0.5em;
}

h3 {
	font-size: 1.25em;
	line-height: calc(var(--base-line-height) * 0.95);
	margin-block: 1.5em 0.5em;
}

p {
	margin-bottom: var(--space-M);
}

p:last-child {
	margin-bottom: 0;
}

.site-header {
	align-items: center;
	display: flex;
	gap: 1.5ch;
	margin-top: var(--space-2XL);
	margin-bottom: var(--space-L);
}

.site-header h2 {
	line-height: calc(var(--base-line-height) * 0.8);
	margin: 0;
}

.site-header h2 a {
	color: inherit;
	text-decoration: none;
}

.site-nav {
	margin-bottom: var(--space-S);
}

.site-nav a {
	color: var(--color-primary);
	margin-right: var(--space-XS);
	text-decoration: none;
}

.site-nav a:hover {
	text-decoration: underline;
}

.site-main {
	padding-bottom: var(--space-4XL);
	flex: 1;
}

.site-list-of-posts {
	list-style: none;
	padding: 0;
}

.site-list-of-posts li {
	align-items: baseline;
	display: grid;
	gap: 2ch;
	grid-template-columns: 6ch 1fr;
	line-height: calc(var(--base-line-height) * 0.9);
}

.site-list-of-posts li:not(:first-of-type) {
	margin-top: var(--space-S);
}

.site-list-of-posts time {
	color: var(--color-txt-light);
	white-space: nowrap;
	text-align: right;
}

.site-list-of-posts .post-link a {
	display: inline;
}

.post-content {
	margin-top: var(--space-L);
}

.post-date {
	color: var(--color-txt-light);
	font-size: 0.875em;
	margin-bottom: var(--space-M);
}

hr {
	background-color: var(--color-border);
	border: 0;
	height: 1px;
	margin: var(--space-XL) 0;
	width: 100%;
}

main {
	flex: 1;
	width: 100%;
}

.site-footer {
	text-align: center;
	padding: var(--space-L) 0;
	border-top: 1px solid var(--color-border);
	margin-top: var(--space-2XL);
	color: var(--color-txt-light);
	font-size: 0.875em;
}

.site-footer a {
	color: var(--color-primary);
	text-decoration: none;
	margin: 0 var(--space-XS);
}

.site-footer a:hover {
	text-decoration: underline;
}
`

	return ioutil.WriteFile(filepath.Join("public", "style.css"), []byte(cssContent), 0644)
}

func (s *Site) generateIndex() error {
	indexTemplate := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Config.Title}}</title>
    <link rel="stylesheet" href="style.css">
    <link rel="alternate" type="application/rss+xml" title="RSS Feed" href="feed.xml">
</head>
<body class="site-view">
    <header class="site-header">
        <div>
            <h2><a href="/">{{.Config.Title}}</a></h2>
            <nav class="site-nav">
                {{range .Pages}}
                <a href="{{.URL}}">{{.Title}}</a>
                {{end}}
            </nav>
        </div>
    </header>

    <main class="site-main">
        {{if .Config.Description}}
        <div style="margin-bottom: 2rem;">
            <p>{{.Config.Description}}</p>
        </div>
        {{end}}

        <hr>

        {{if .Posts}}
        {{$currentYear := 0}}
        {{range .Posts}}
        {{$postYear := .Date.Year}}
        {{if ne $postYear $currentYear}}
        {{if ne $currentYear 0}}
        </ul>
        {{end}}
        <h2>{{$postYear}}</h2>
        <ul class="site-list-of-posts">
        {{$currentYear = $postYear}}
        {{end}}
            <li>
                <time>{{.Date.Format "2 Jan"}}</time>
                <div class="post-link"><a href="{{.URL}}">{{.Title}}</a></div>
            </li>
        {{end}}
        {{if .Posts}}
        </ul>
        {{end}}
        {{end}}
    </main>

    <footer class="site-footer">
        <p>
            Made with <strong>BazelBlog</strong>
            {{if .Config.Socials}}
            | Connect with us:
            {{range $platform, $url := .Config.Socials}}
            <a href="{{$url}}">{{$platform}}</a>
            {{end}}
            {{end}}
        </p>
    </footer>
</body>
</html>`

	tmpl, err := template.New("index").Funcs(sprig.FuncMap()).Parse(indexTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join("public", "index.html"))
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, s)
}

func (s *Site) generatePosts() error {
	postTemplate := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}} - {{.Config.Title}}</title>
    <link rel="stylesheet" href="../style.css">
</head>
<body class="site-view">
    <header class="site-header">
        <div>
            <h2><a href="../">{{.Config.Title}}</a></h2>
            <nav class="site-nav">
                <a href="../">Home</a>
                {{range .Pages}}
                <a href="../{{.URL}}">{{.Title}}</a>
                {{end}}
            </nav>
        </div>
    </header>

    <main class="site-main">
        <h1>{{.Title}}</h1>
        <div class="post-date">{{.Date.Format "January 2, 2006"}}</div>
        <div class="post-content">
            {{.Content}}
        </div>
    </main>

    <footer class="site-footer">
        <p>
            Made with <strong>BazelBlog</strong>
            {{if .Config.Socials}}
            | Connect with us:
            {{range $platform, $url := .Config.Socials}}
            <a href="{{$url}}">{{$platform}}</a>
            {{end}}
            {{end}}
        </p>
    </footer>
</body>
</html>`

	tmpl, err := template.New("post").Parse(postTemplate)
	if err != nil {
		return err
	}

	for _, post := range s.Posts {
		file, err := os.Create(filepath.Join("public", post.URL))
		if err != nil {
			return err
		}

		data := struct {
			Title   string
			Date    time.Time
			Content template.HTML
			Config  *config.Config
			Pages   []Page
		}{
			Title:   post.Title,
			Date:    post.Date,
			Content: template.HTML(post.Content),
			Config:  s.Config,
			Pages:   s.Pages,
		}

		err = tmpl.Execute(file, data)
		file.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Site) generatePages() error {
	pageTemplate := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}} - {{.Config.Title}}</title>
    <link rel="stylesheet" href="../style.css">
</head>
<body class="site-view">
    <header class="site-header">
        <div>
            <h2><a href="../">{{.Config.Title}}</a></h2>
            <nav class="site-nav">
                <a href="../">Home</a>
                {{range .Pages}}
                <a href="../{{.URL}}">{{.Title}}</a>
                {{end}}
            </nav>
        </div>
    </header>

    <main class="site-main">
        {{.Content}}
    </main>

    <footer class="site-footer">
        <p>
            Made with <strong>BazelBlog</strong>
            {{if .Config.Socials}}
            | Connect with us:
            {{range $platform, $url := .Config.Socials}}
            <a href="{{$url}}">{{$platform}}</a>
            {{end}}
            {{end}}
        </p>
    </footer>
</body>
</html>`

	tmpl, err := template.New("page").Parse(pageTemplate)
	if err != nil {
		return err
	}

	for _, page := range s.Pages {
		// Create output file
		file, err := os.Create(filepath.Join("public", page.URL))
		if err != nil {
			return err
		}

		// For Markdown pages, Content is already processed HTML
		// For HTML pages, we need to extract body content
		var pageContent template.HTML
		if strings.HasSuffix(page.Filename, ".md") {
			// Markdown page - content is already processed HTML
			pageContent = template.HTML(page.Content)
		} else {
			// HTML page - extract body content
			src := filepath.Join("pages", page.Filename)
			content, err := ioutil.ReadFile(src)
			if err != nil {
				return err
			}

			contentStr := string(content)
			bodyStart := strings.Index(contentStr, "<body>")
			bodyEnd := strings.Index(contentStr, "</body>")

			if bodyStart != -1 && bodyEnd != -1 {
				pageContent = template.HTML(contentStr[bodyStart+6 : bodyEnd])
			} else {
				pageContent = template.HTML(contentStr)
			}
		}

		data := struct {
			Title   string
			Content template.HTML
			Config  *config.Config
			Pages   []Page
		}{
			Title:   page.Title,
			Content: pageContent,
			Config:  s.Config,
			Pages:   s.Pages,
		}

		err = tmpl.Execute(file, data)
		file.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Site) generateRSS() error {
	rssTemplate := `<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
	<channel>
		<title>{{.Config.Title}}</title>
		<description>{{.Config.Description}}</description>
		<link>{{.Config.BaseURL}}</link>
		<atom:link href="{{.Config.BaseURL}}/feed.xml" rel="self" type="application/rss+xml" />
		<language>en-us</language>
		<lastBuildDate>{{.BuildDate}}</lastBuildDate>
		<generator>Bazel Static Site Generator</generator>
		{{range .Posts}}
		<item>
			<title>{{.Title}}</title>
			<description><![CDATA[{{.Content}}]]></description>
			<link>{{$.Config.BaseURL}}/{{.URL}}</link>
			<guid>{{$.Config.BaseURL}}/{{.URL}}</guid>
			<pubDate>{{.Date.Format "Mon, 02 Jan 2006 15:04:05 -0700"}}</pubDate>
		</item>
		{{end}}
	</channel>
</rss>`

	tmpl, err := template.New("rss").Parse(rssTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join("public", "feed.xml"))
	if err != nil {
		return err
	}
	defer file.Close()

	data := struct {
		*Site
		BuildDate string
	}{
		Site:      s,
		BuildDate: time.Now().Format("Mon, 02 Jan 2006 15:04:05 -0700"),
	}

	return tmpl.Execute(file, data)
}

// markdownToHTML converts markdown to HTML using enhanced Goldmark with extensions
func (s *Site) markdownToHTML(markdown string) string {
	// Configure Goldmark with extensions
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,           // GitHub Flavored Markdown
			extension.Table,         // Tables
			extension.Strikethrough, // Strikethrough
			extension.Linkify,       // Auto-linkify URLs
			extension.TaskList,      // Task lists
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // Auto-generate heading IDs
		),
		goldmark.WithRendererOptions(
			goldmarkhtml.WithHardWraps(), // Hard line breaks
			goldmarkhtml.WithXHTML(),     // XHTML output
			goldmarkhtml.WithUnsafe(),    // Allow raw HTML
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		// Fallback to plain text if conversion fails
		return markdown
	}
	return buf.String()
}
