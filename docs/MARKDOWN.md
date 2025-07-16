# Markdown Guide for Bazel Blog

This guide covers the markdown syntax and commands supported by Bazel Blog for creating posts and content.

## Basic Markdown Syntax

### Headers
```markdown
# H1 - Main Title
## H2 - Section Title
### H3 - Subsection Title
#### H4 - Sub-subsection Title
##### H5 - Minor Heading
###### H6 - Smallest Heading
```

### Text Formatting
```markdown
**Bold text**
*Italic text*
***Bold and italic text***
~~Strikethrough text~~
`Inline code`
```

### Lists

#### Unordered Lists
```markdown
- Item 1
- Item 2
  - Nested item 2.1
  - Nested item 2.2
- Item 3

* Alternative bullet style
* Another item
```

#### Ordered Lists
```markdown
1. First item
2. Second item
   1. Nested ordered item
   2. Another nested item
3. Third item
```

### Links
```markdown
[Link text](https://example.com)
[Link with title](https://example.com "This is a title")
[Reference link][1]

[1]: https://example.com
```

### Images
```markdown
![Alt text](image.jpg)
![Alt text](image.jpg "Image title")
![Reference image][image-ref]

[image-ref]: image.jpg
```

### Code Blocks

#### Inline Code
```markdown
Use `code` in your text.
```

#### Fenced Code Blocks
````markdown
```
Plain code block
```

```javascript
// JavaScript code block
function hello() {
    console.log("Hello, world!");
}
```

```go
// Go code block
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

```python
# Python code block
def hello():
    print("Hello, world!")
```

```html
<!-- HTML code block -->
<div class="container">
    <h1>Hello, world!</h1>
</div>
```

```css
/* CSS code block */
.container {
    max-width: 800px;
    margin: 0 auto;
}
```
````

### Blockquotes
```markdown
> This is a blockquote
> spanning multiple lines

> ## Blockquote with heading
> 
> This is a longer blockquote with multiple paragraphs.
> 
> Second paragraph in the blockquote.
```

### Horizontal Rules
```markdown
---
***
___
```

### Tables
```markdown
| Header 1 | Header 2 | Header 3 |
|----------|----------|----------|
| Cell 1   | Cell 2   | Cell 3   |
| Cell 4   | Cell 5   | Cell 6   |

| Left | Center | Right |
|:-----|:------:|------:|
| Left | Center | Right |
| Text | Text   | Text  |
```

### Line Breaks
```markdown
Line 1  
Line 2 (two spaces at end of Line 1)

Line 1

Line 2 (blank line between)
```

## Frontmatter

All posts in Bazel Blog support YAML frontmatter at the beginning of the file:

```yaml
---
title: "Your Post Title"
date: "January 1, 2025"
description: "Optional post description"
tags: ["tag1", "tag2", "tag3"]
author: "Your Name"
---
```

### Frontmatter Fields

- **title**: The post title (required)
- **date**: Publication date in human-readable format (required)
- **description**: Brief description of the post (optional)
- **tags**: Array of tags for categorization (optional)
- **author**: Author name (optional)

## HTML Support

Bazel Blog supports inline HTML within markdown for advanced formatting:

```markdown
This is <mark>highlighted text</mark> using HTML.

<div class="custom-class">
    <p>Custom HTML content</p>
</div>

<details>
<summary>Click to expand</summary>
Hidden content here
</details>
```

## Extended Markdown Features

### Task Lists
```markdown
- [x] Completed task
- [ ] Incomplete task
- [ ] Another incomplete task
```

### Footnotes
```markdown
This is a sentence with a footnote[^1].

[^1]: This is the footnote content.
```

### Definition Lists
```markdown
Term 1
:   Definition 1

Term 2
:   Definition 2a
:   Definition 2b
```

## Best Practices

1. **Use descriptive headers** - Create a clear hierarchy with H1 for the main title and H2-H6 for sections
2. **Add alt text to images** - Always include descriptive alt text for accessibility
3. **Use code blocks for code** - Properly format code with language-specific syntax highlighting
4. **Keep line length reasonable** - Wrap text at around 80-100 characters for readability
5. **Use frontmatter consistently** - Include required fields (title, date) in all posts
6. **Link responsibly** - Use descriptive link text instead of "click here"
7. **Structure content logically** - Use lists, headers, and whitespace to organize content

## Example Post

Here's a complete example of a well-formatted post:

```markdown
---
title: "Getting Started with Bazel Blog"
date: "January 15, 2025"
description: "A comprehensive guide to creating your first blog post"
tags: ["tutorial", "getting-started", "markdown"]
author: "Bazel Team"
---

# Getting Started with Bazel Blog

Welcome to **Bazel Blog**! This guide will help you create your first post.

## Creating Your First Post

1. Run the post creation command:
   ```bash
   bazel post
   ```

2. Choose "Create new post" from the menu
3. Enter your post title and content

### Writing Content

You can use all standard markdown features:

- **Bold** and *italic* text
- [Links](https://example.com)
- Code blocks with syntax highlighting
- Images and more!

> Remember to use frontmatter to set your post metadata.

## Next Steps

After creating your post:

1. Build your site: `bazel build`
2. Preview locally: `bazel serve`
3. Deploy to your hosting platform

Happy blogging! 🎉
```

This example demonstrates proper use of frontmatter, headers, formatting, code blocks, lists, and other markdown features supported by Bazel Blog.
