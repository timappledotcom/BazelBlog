# Bazel Static Site Generator - Installation Guide

## Quick Installation

1. **Install the package:**
   ```bash
   sudo dpkg -i bazel-ssg_1.0.0_amd64.deb
   ```

2. **If you get dependency errors, run:**
   ```bash
   sudo apt-get install -f
   ```

## Usage

### Create a new site:
```bash
bazel new my-website
cd my-website
```

### Interactive menu (run inside a site directory):
```bash
bazel
```

### Command line usage:
```bash
bazel new <site-name>    # Create a new site
bazel post <title>       # Create a new post
bazel page <title>       # Create a new page
bazel build              # Build the site
bazel serve              # Start dev server with live reload
```

## Features

- **Interactive Terminal Menu**: Built with Bubbletea for a great user experience
- **Live Reload Development Server**: Automatically rebuilds and refreshes on file changes
- **RSS Feed Generation**: Automatic RSS feed creation for your posts
- **Theme Configuration**: Multiple themes and font options
- **Social Media Integration**: Configure social media links
- **Markdown Support**: Write posts and pages in Markdown
- **Site Configuration**: Edit site title, domain, and other settings

## Configuration Menu

When you run `bazel` in a site directory, you'll get an interactive menu with:

- **Configuration**: Access theme, font, social media, and site settings
- **Build**: Build your site for production
- **Dev Server**: Start the development server with live reload
- **Quit**: Exit the application

## Uninstallation

To remove the package:
```bash
sudo dpkg -r bazel-ssg
```

## System Requirements

- Ubuntu/Debian-based Linux distribution
- libc6 (usually pre-installed)

## Package Details

- **Package Name**: bazel-ssg
- **Version**: 1.0.0
- **Architecture**: amd64
- **Installed Size**: ~13MB
- **Binary Location**: /usr/bin/bazel
- **Symlink Created**: /usr/local/bin/bazel (for easier access)
