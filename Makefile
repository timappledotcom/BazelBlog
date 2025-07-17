# Bazel Blog - Static Site Generator
# Makefile for building and packaging

VERSION := 1.4.3
BINARY_NAME := bazel
PACKAGE_NAME := bazel-blog

# Go build settings
GOOS := linux
GOARCH := amd64
LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

# Directories
BUILD_DIR := build
DIST_DIR := dist

.PHONY: all build clean test install uninstall deb rpm demo help

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME) v$(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) cmd/bazel/main.go
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)
	rm -rf rpmbuild/BUILD
	rm -rf rpmbuild/BUILDROOT
	rm -rf rpmbuild/RPMS
	rm -rf rpmbuild/SRPMS
	rm -f $(BINARY_NAME)
	rm -f *.deb
	rm -f *.rpm
	rm -f debian/usr/bin/$(BINARY_NAME)
	@echo "Clean complete"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Install locally (for development)
install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installation complete"

# Uninstall
uninstall:
	@echo "Removing $(BINARY_NAME) from /usr/local/bin..."
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "Uninstall complete"

# Build Debian package
deb: build
	@echo "Building Debian package..."
	@mkdir -p debian/usr/bin
	cp $(BUILD_DIR)/$(BINARY_NAME) debian/usr/bin/
	chmod +x debian/usr/bin/$(BINARY_NAME)
	dpkg-deb --build debian $(PACKAGE_NAME)_$(VERSION)_amd64.deb
	@echo "Debian package built: $(PACKAGE_NAME)_$(VERSION)_amd64.deb"

# Build RPM package
rpm: build
	@echo "Building RPM package..."
	@mkdir -p $(DIST_DIR)
	fpm -s dir -t rpm \
		--name $(PACKAGE_NAME) \
		--version $(VERSION) \
		--iteration 1 \
		--architecture x86_64 \
		--description "A fast and simple static site generator with multi-site support, interactive UI, and beautiful themes." \
		--url "https://github.com/timappledotcom/BazelBlog" \
		--license "MIT" \
		--maintainer "Tim Apple <tim@example.com>" \
		--vendor "Tim Apple" \
		--package $(DIST_DIR)/$(PACKAGE_NAME)-$(VERSION)-1.x86_64.rpm \
		$(BUILD_DIR)/$(BINARY_NAME)=/usr/bin/$(BINARY_NAME)
	@echo "RPM package built: $(DIST_DIR)/$(PACKAGE_NAME)-$(VERSION)-1.x86_64.rpm"

# Run demo
demo: build
	@echo "Running demo..."
	cp $(BUILD_DIR)/$(BINARY_NAME) ./$(BINARY_NAME)
	./demo.sh
	rm -f ./$(BINARY_NAME)

# Development build (with debug info)
dev:
	@echo "Building development version..."
	go build -race -o $(BINARY_NAME) cmd/bazel/main.go

# Format code
fmt:
	@echo "Formatting Go code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting Go code..."
	golangci-lint run

# Show help
help:
	@echo "Bazel Blog - Build System"
	@echo ""
	@echo "Available targets:"
	@echo "  build     - Build the binary"
	@echo "  clean     - Clean build artifacts"
	@echo "  test      - Run tests"
	@echo "  install   - Install binary locally"
	@echo "  uninstall - Remove installed binary"
	@echo "  deb       - Build Debian package"
	@echo "  rpm       - Build RPM package (requires fpm)"
	@echo "  demo      - Run the demo script"
	@echo "  dev       - Build development version"
	@echo "  fmt       - Format Go code"
	@echo "  lint      - Lint Go code (requires golangci-lint)"
	@echo "  help      - Show this help"
	@echo ""
	@echo "Version: $(VERSION)"