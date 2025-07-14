#!/bin/bash
set -e

# Build the Go binary
echo "Building bazel binary..."
go build -o bazel cmd/bazel/main.go

# Build the RPM package using fpm
echo "Building RPM package..."
fpm -s dir -t rpm \
  --name bazel-blog \
  --version 1.4.0 \
  --iteration 1 \
  --architecture x86_64 \
  --description "A fast and simple static site generator with multi-site support, interactive UI, and beautiful themes. Features include multi-site registry and selection, interactive Bubbletea-powered UI, built-in themes (Pika Beach, Catppuccin, Dracula, Nord, Tokyo Night, 3li7e), auto-rebuild on theme changes, live development server, Markdown post support, and automatic footer with BazelBlog attribution and social links." \
  --url "https://github.com/timappledotcom/BazelBlog" \
  --license "MIT" \
  --maintainer "Tim Apple <tim@example.com>" \
  --vendor "Tim Apple" \
  --package bazel-blog-1.4.0-1.x86_64.rpm \
  bazel=/usr/bin/bazel

echo "Package built successfully: bazel-blog-1.4.0-1.x86_64.rpm"
echo "Install with: sudo rpm -i bazel-blog-1.4.0-1.x86_64.rpm"
