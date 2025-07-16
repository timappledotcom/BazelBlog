#!/bin/bash
set -e

echo "Building RPM package using Makefile..."
make rpm

echo "Package built successfully!"
echo "Install with: sudo rpm -i dist/bazel-blog-1.4.2-1.x86_64.rpm"
