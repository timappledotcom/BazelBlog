# Contributing to Bazel Blog

Thank you for your interest in contributing to Bazel Blog! This guide will help you get started.

## Development Setup

### Prerequisites

- Go 1.21 or later
- Make (for build automation)
- For package building:
  - `dpkg` (for Debian packages)
  - `fpm` (for RPM packages)

### Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/timappledotcom/BazelBlog.git
   cd BazelBlog
   ```

2. **Build the project:**
   ```bash
   make build
   ```

3. **Run tests:**
   ```bash
   make test
   ```

4. **Install for local development:**
   ```bash
   make install
   ```

## Development Workflow

### Building

- `make build` - Build the binary
- `make dev` - Build development version with race detection
- `make clean` - Clean all build artifacts

### Testing

- `make test` - Run all tests
- `make demo` - Run the demo script

### Code Quality

- `make fmt` - Format Go code
- `make lint` - Run linter (requires golangci-lint)

### Packaging

- `make deb` - Build Debian package
- `make rpm` - Build RPM package (requires fpm)

## Project Structure

```
bazel_blog/
├── cmd/bazel/          # Main application entry point
├── internal/           # Internal packages
│   ├── config/         # Configuration management
│   ├── generator/      # Site generation logic
│   ├── registry/       # Site registry management
│   ├── ui/             # Bubble Tea interactive interface
│   └── upgrade/        # Site upgrade functionality
├── docs/               # Documentation
├── debian/             # Debian package structure
├── rpmbuild/           # RPM package structure
└── Makefile           # Build automation
```

## Coding Standards

### Go Code Style

- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions focused and small
- Handle errors appropriately

### Commit Messages

Use conventional commit format:
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `refactor:` for code refactoring
- `test:` for test additions/changes

Example:
```
feat: add new theme selection interface

- Implement interactive theme picker
- Add preview functionality
- Update configuration system
```

## Adding New Features

### Adding a New Theme

1. Add the theme name to `ColorSchemes` in `internal/config/config.go`
2. Add CSS variables in `GetCSSVariables()` method
3. Test the theme with existing sites
4. Update documentation

### Adding New Commands

1. Add command handling in `cmd/bazel/main.go`
2. Implement the functionality in appropriate internal package
3. Add tests for the new functionality
4. Update help text and documentation

### Adding UI Components

1. Create new components in `internal/ui/`
2. Follow Bubble Tea patterns
3. Ensure consistent styling with existing components
4. Test interactive functionality

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/config
```

### Writing Tests

- Write unit tests for all new functionality
- Use table-driven tests where appropriate
- Mock external dependencies
- Test error conditions

## Documentation

### Updating Documentation

- Update README.md for user-facing changes
- Update docs/CHANGELOG.md for all changes
- Add examples for new features
- Keep documentation current with code changes

### Documentation Structure

- `README.md` - Main project documentation
- `docs/INSTALL.md` - Installation guide
- `docs/CHANGELOG.md` - Version history
- `docs/MARKDOWN.md` - Markdown syntax guide
- `CONTRIBUTING.md` - This file

## Release Process

### Version Numbering

We use semantic versioning (MAJOR.MINOR.PATCH):
- MAJOR: Breaking changes
- MINOR: New features, backward compatible
- PATCH: Bug fixes, backward compatible

### Creating a Release

1. Update version in relevant files:
   - `cmd/bazel/main.go` (Version constant)
   - `Makefile` (VERSION variable)
   - Build scripts
   - Package files

2. Update `docs/CHANGELOG.md`

3. Build and test packages:
   ```bash
   make clean
   make deb
   make rpm
   ```

4. Create git tag:
   ```bash
   git tag v1.x.x
   git push origin v1.x.x
   ```

## Getting Help

- Check existing issues on GitHub
- Read the documentation in the `docs/` directory
- Look at the code examples in the project
- Ask questions in GitHub discussions

## Code of Conduct

- Be respectful and inclusive
- Focus on constructive feedback
- Help others learn and grow
- Follow the project's coding standards

Thank you for contributing to Bazel Blog! 🚀