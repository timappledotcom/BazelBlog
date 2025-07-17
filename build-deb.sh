#!/bin/bash
set -e

echo "Building Debian package using Makefile..."
make deb

echo "Package built successfully!"
echo "Install with: sudo dpkg -i bazel-blog_1.4.3_amd64.deb"
