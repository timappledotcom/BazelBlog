# Changelog

All notable changes to Bazel Blog will be documented in this file.

## [1.3.0] - 2025-07-14

### Added
- **Automatic Footer**: All generated pages now include a footer with BazelBlog attribution and social links
- Footer displays "Made with **BazelBlog**" on every page
- Social media links from site configuration are automatically included in footer
- Footer styling is consistent with site themes and responsive design

### Changed
- Updated all page templates (index, posts, pages) to include the new footer
- Enhanced CSS with footer-specific styling
- Improved overall page layout with proper footer positioning

### Technical Details
- Footer is automatically generated in all three main templates
- Footer content is conditional - only shows social links if they exist in configuration
- Footer styling uses existing CSS variables for theme consistency
- Added `site-footer` CSS class for styling

## [1.2.0] - 2025-07-11

### Added
- Multi-site registry system for managing multiple sites
- Site selection interface when not in a site directory
- Underscore-based filenames for new posts and pages
- Enhanced markdown support with improved formatting
- Helix editor support

### Changed
- Improved command-line interface organization
- Better error handling and user feedback
- Enhanced site management capabilities

## [1.1.0] - 2025-07-09

### Added
- Interactive Bubbletea-powered UI
- Built-in themes: Pika Beach, Catppuccin variants, Dracula, Nord, Tokyo Night, 3li7e
- Live development server with auto-reload
- Font selection system
- Social media configuration

### Changed
- Redesigned configuration system
- Improved theme switching
- Better development workflow

## [1.0.0] - 2025-07-01

### Added
- Initial release of Bazel Blog
- Basic static site generation
- Markdown post support
- HTML page support
- Basic theming system
- Configuration management
