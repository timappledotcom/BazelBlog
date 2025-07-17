# Bazel Blog v1.4.3 Release Notes

## 🎉 What's New in v1.4.3

### ✨ Major Features
- **Delete Functionality**: Added ability to delete posts and pages with confirmation prompts to prevent accidental deletions
- **Enhanced UI Navigation**: Improved menu navigation with better error handling and user feedback
- **Code Quality Improvements**: Modernized Go patterns and cleaned up codebase for better maintainability

### 🔧 Improvements
- **Menu System**: Enhanced post and page management menus with delete options
- **Error Handling**: Better error messages and user feedback throughout the application
- **Code Structure**: Modernized Go code patterns for improved performance and readability

### 🐛 Bug Fixes
- **Typos**: Fixed various typos throughout the codebase
- **Code Issues**: Resolved Go linting issues and modernized loop patterns
- **UI Flow**: Improved menu state management and navigation flow
- **Emoji Rendering**: Fixed emoji display issues in upgrade messages

### 🏗️ Technical Details
- Added delete confirmation dialogs to prevent accidental deletions
- Modernized for loops using range over int pattern
- Enhanced error handling in menu operations
- Improved code readability and maintainability
- Cleaned up build artifacts from repository

## 📦 Installation

### Debian/Ubuntu (.deb)
```bash
wget https://github.com/timappledotcom/BazelBlog/releases/download/v1.4.3/bazel-blog_1.4.3_amd64.deb
sudo dpkg -i bazel-blog_1.4.3_amd64.deb
```

### Red Hat/CentOS/Fedora (.rpm)
```bash
wget https://github.com/timappledotcom/BazelBlog/releases/download/v1.4.3/bazel-blog-1.4.3-1.x86_64.rpm
sudo rpm -i bazel-blog-1.4.3-1.x86_64.rpm
```

### From Source
```bash
git clone https://github.com/timappledotcom/BazelBlog.git
cd BazelBlog
make build
sudo make install
```

## 🚀 Getting Started

1. Create a new site:
   ```bash
   bazel new site my-blog
   cd my-blog
   ```

2. Start the development server:
   ```bash
   bazel serve
   ```

3. Visit http://localhost:3000 to see your site

4. Create your first post:
   ```bash
   bazel post
   ```

5. Build for production:
   ```bash
   bazel build
   ```

## 🎨 Features

- **Multi-Site Management**: Registry system for managing multiple sites
- **Interactive Configuration**: Beautiful TUI for themes, fonts, and social media
- **Multiple Themes**: 9 built-in color schemes including Catppuccin variants
- **Font Selection**: Choose from various typography options
- **Live Development Server**: Auto-reload server for development
- **Markdown Support**: Write posts in Markdown with frontmatter
- **Social Media Integration**: Configure links to your social profiles
- **Delete Functionality**: Safely delete posts and pages with confirmation

## 🔄 Upgrading from Previous Versions

If you're upgrading from a previous version, run the upgrade command in your site directory:

```bash
bazel upgrade
```

This will:
- Update your site structure to the latest format
- Migrate any configuration changes
- Rebuild your site with the latest templates

## 📋 Full Changelog

See [CHANGELOG.md](docs/CHANGELOG.md) for complete version history.

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details on development setup, coding standards, and the release process.

## 📄 License

MIT License - feel free to use and modify as needed!

---

**Package Files:**
- `bazel-blog_1.4.3_amd64.deb` (7.9 MB) - Debian/Ubuntu package
- `bazel-blog-1.4.3-1.x86_64.rpm` (8.7 MB) - Red Hat/CentOS/Fedora package