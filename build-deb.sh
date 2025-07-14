#!/bin/bash
set -e

# Build the Go binary
echo "Building bazel binary..."
go build -o bazel cmd/bazel/main.go

# Copy binary to debian structure
cp bazel debian/usr/bin/

# Set proper permissions
chmod +x debian/usr/bin/bazel

# Build the package
echo "Building Debian package..."
dpkg-deb --build debian bazel-blog_1.4.0_amd64.deb

echo "Package built successfully: bazel-blog_1.4.0_amd64.deb"
echo "Install with: sudo dpkg -i bazel-blog_1.4.0_amd64.deb"
