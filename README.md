# Bazel Blog - Static Site Generator

A simple, fast static site generator written in Go with interactive configuration using Bubble Tea.

## Features

- **Multi-Site Management**: Registry system for managing multiple sites from anywhere
- **Interactive Configuration**: Beautiful TUI for setting themes, fonts, and social media links
- **Vanilla Output**: Generates clean HTML, CSS, and JavaScript without frameworks
- **Multiple Themes**: Built-in color schemes including Pika Beach, Catppuccin variants, Dracula, Nord, Tokyo Night, and 3li7e
- **Font Selection**: Choose from system fonts or web-safe alternatives
- **Social Media Integration**: Configure links to your social profiles
- **Markdown Support**: Write posts in Markdown with frontmatter
- **Clean HTML Pages**: Create custom pages with raw HTML
- **Live Development Server**: Auto-reload server for development
- **Automatic Footer**: BazelBlog attribution footer with social links on all pages

## Installation

1. Clone the repository
2. Build the binary:
   ```bash
   go build -o bazel cmd/bazel/main.go
   ```

## Usage

### Creating a New Site

```bash
bazel new site my-site
cd my-site
```

This creates a new directory with the following structure:
```
my-site/
├── bazel.toml          # Configuration file
├── posts/              # Markdown posts
├── pages/              # HTML pages
└── themes/             # (Reserved for future use)
```

### Creating Content

#### Posts (Markdown)
```bash
bazel post
```

This opens an interactive menu for post management, allowing you to create, edit, or manage posts.

Posts support frontmatter:
```markdown
---
title: My Post Title
date: January 1, 2025
---

Your content here with **markdown** support!
```

#### Pages (HTML)
```bash
bazel page
```

This opens an interactive menu for page management, allowing you to create, edit, or organize pages.

### Interactive Configuration

```bash
bazel config
```

This opens an interactive menu where you can:
- **Site Settings**: Edit title, description, and domain
- **Theme**: Choose from 9 color schemes
- **Font**: Select from various font options
- **Social Links**: Configure social media profiles

### Building the Site

```bash
bazel build
```

This generates the static site in the `public/` directory.

### Configuration

The `bazel.toml` file stores your site configuration:

```json
{
  "site_name": "my-site",
  "title": "My Site",
  "description": "A static site generated with Bazel Blog",
  "base_url": "https://example.com",
  "theme": {
    "color_scheme": "catppuccin-latte",
    "font": "system"
  },
  "socials": {
    "github": "https://github.com/yourusername",
    "twitter": "https://twitter.com/yourusername"
  }
}
```

## Available Themes

### Default Theme
- **pika-beach**: Warm, beach-inspired theme (default)

### Catppuccin Variants
- **catppuccin-latte**: Light theme with warm, pastel colors
- **catppuccin-frappe**: Medium-dark theme with soft purple accents
- **catppuccin-macchiato**: Dark theme with vibrant purple highlights
- **catppuccin-mocha**: Darkest theme with beautiful purple tones

### Popular Developer Themes
- **dracula**: Dark theme with purple, pink, and cyan accents
- **nord**: Arctic, north-bluish color palette
- **tokyo-night**: Dark theme inspired by Tokyo's neon lights
- **3li7e**: Retro green-on-black CRT monitor theme

## Available Fonts

- **pika-serif**: Source Serif 4 (default)
- **system**: System default fonts
- **serif**: Georgia, serif
- **monospace**: Courier New, monospace
- **arial**: Arial, sans-serif
- **helvetica**: Helvetica Neue, sans-serif
- **georgia**: Georgia, serif
- **times**: Times New Roman, serif

## Social Platforms

Supported social media platforms:
- Twitter
- GitHub
- LinkedIn
- Facebook
- Instagram
- YouTube
- Mastodon
- Email

## Demo

Run the demo script to see Bazel in action:
```bash
./demo.sh
```

## Project Structure

```
bazel_blog/
├── cmd/bazel/          # Main application entry point
├── internal/
│   ├── config/         # Configuration management
│   ├── generator/      # Site generation logic
│   ├── registry/       # Site registry management
│   ├── ui/             # Bubble Tea interactive interface
│   └── upgrade/        # Site upgrade functionality
├── debian/             # Debian package structure
├── build-deb.sh        # Debian package build script
├── demo.sh             # Demo script
├── go.mod              # Go module file
├── go.sum              # Go module checksums
├── INSTALL.md          # Installation instructions
└── README.md           # This file
```

## License

MIT License - feel free to use and modify as needed!

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.
