#!/bin/bash

echo "🏗️  Bazel Static Site Generator Demo"
echo "======================================"
echo ""

# Clean up previous demo
rm -rf demo-site

echo "1. Creating a new site..."
./bazel new demo-site

echo ""
echo "2. Navigating to the site directory..."
cd demo-site

echo ""
echo "3. Creating some sample content..."

# Create a sample post
cat > posts/welcome.md << 'EOF'
---
title: Welcome to My Site
---

This is my first post using the **Bazel** static site generator!

It supports:
- Markdown content
- Multiple themes
- Font selection
- Social media links
- Clean vanilla HTML/CSS/JS output

Pretty cool, right?
EOF

# Create another post
cat > posts/about-bazel.md << 'EOF'
---
title: About Bazel Generator
---

Bazel is a static site generator written in Go that focuses on simplicity and customization.

Key features:
- Interactive theme and font selection with Bubble Tea
- Vanilla HTML/CSS/JS output (no frameworks)
- Simple configuration
- Easy content management
EOF

# Create a sample page
cat > pages/contact.html << 'EOF'
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Contact</title>
</head>
<body>
    <h1>Contact</h1>
    <p>Get in touch with me!</p>
    <p>Email: hello@example.com</p>
    <p>GitHub: github.com/yourname</p>
</body>
</html>
EOF

echo ""
echo "4. Building the site..."
../bazel build

echo ""
echo "5. Site generated! Check the public/ directory:"
ls -la public/

echo ""
echo "6. To configure the site interactively, run:"
echo "   ./bazel"
echo ""
echo "   This will open a menu where you can:"
echo "   - Set the color theme"
echo "   - Choose fonts"
echo "   - Configure social media links"
echo "   - Build the site"
echo ""
echo "Demo complete! 🎉"
