# Changelog

All notable changes to Bazel Blog will be documented in this file.

## [1.4.2] - 2025-07-16

### Added
- **Organized Directory Structure**: Posts and pages now generate in proper subdirectories
- **Automatic Migration**: Upgrade system migrates existing flat structure to organized directories
- **Social Media Optimization**: Comprehensive Open Graph and Twitter Card meta tags
- **SEO Enhancement**: Canonical URLs, proper meta descriptions, and article metadata
- **Better Site Organization**: Clear separation between posts/, pages/, and root files

### Changed
- **Site Structure**: Generated sites now use `public/posts/` and `public/pages/` directories
- **Template Paths**: Updated CSS and navigation links to work with new directory structure
- **URL Structure**: Posts now accessible at `/posts/filename.html`, pages at `/pages/filename.html`
- **Title Processing**: Underscores in filenames now convert to spaces in titles

### Fixed
- **Post Title Display**: Filenames with underscores now show as proper spaced titles
- **Frontmatter Stripping**: YAML frontmatter no longer appears in post content
- **Page Format**: Default about page now created as Markdown (.md) instead of HTML for easier editing
- **Directory Organization**: No more flat file structure in public directory
- **Navigation Links**: Proper relative paths for subdirectory navigation
- **CSS Loading**: Correct relative paths for stylesheets in subdirectories
- **Template Data**: Fixed URL field access in post and page templates

### Technical Details
- Enhanced builder.go to create proper directory structure
- Updated templates with relative paths for subdirectories
- Added migration logic in upgrade system for existing sites
- Improved file organization for better deployment and SEO
- Fixed frontmatter parsing to properly separate content from metadata
- Added comprehensive social media meta tags for better sharing

## [1.4.1] - 2025-07-15

### Added
- **Comprehensive Makefile**: Added standardized build system with targets for build, clean, test, install, package creation
- **CONTRIBUTING.md**: Complete development guide with setup instructions, coding standards, and release process
- **Improved Project Structure**: Organized documentation into `docs/` directory for better maintainability

### Changed
- **Enhanced Build Scripts**: Updated build-deb.sh and build-rpm.sh to use Makefile for consistency
- **Better Documentation Organization**: Moved INSTALL.md, CHANGELOG.md, and MARKDOWN.md to docs/ directory
- **Improved .gitignore**: Added comprehensive patterns to prevent build artifacts from being committed
- **Updated Demo Script**: Enhanced demo.sh with proper cleanup and build process

### Removed
- **Build Artifacts**: Cleaned up committed binaries, packages, and build directories from repository
- **Redundant Build Directories**: Removed rpmbuild build artifacts that shouldn't be in version control

### Technical Improvements
- Standardized version management across all build files
- Better separation of source code and build artifacts
- Improved developer experience with clear build targets
- Enhanced project maintainability and contribution workflow

## [1.4.0] - 2025-07-14

### Added
- **Comprehensive Markdown Documentation**: Added separate MARKDOWN.md file with complete guide to markdown syntax and commands
- **Enhanced README**: Improved documentation structure with clear links to markdown guide
- **Upgraded Version System**: Updated upgrade system to support version 1.4.0 with proper migration path
- **Improved User Experience**: Better documentation organization and user-friendly guides

### Changed
- Updated README to include link to dedicated markdown documentation
- Enhanced upgrade system with new version tracking for 1.4.0
- Improved project structure with better documentation organization

### Technical Details
- Added upgradeToV1_4_0 function for smooth version transitions
- Updated all package files (DEB, RPM) to version 1.4.0
- Enhanced build scripts with new version numbers
- Updated changelog system for better version tracking

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
