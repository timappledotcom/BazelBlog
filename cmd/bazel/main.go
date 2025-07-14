package main

import (
	"fmt"
	"os"
	"time"

	"github.com/yourusername/bazel_blog/internal/generator"
	"github.com/yourusername/bazel_blog/internal/registry"
	"github.com/yourusername/bazel_blog/internal/ui"
	"github.com/yourusername/bazel_blog/internal/upgrade"
)

const Version = "1.3.0"

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]
	switch command {
	case "-v", "--version", "version":
		printVersion()
		return

	case "new":
		if len(os.Args) < 4 || os.Args[2] != "site" {
			fmt.Println("Usage: bazel new site <site-name>")
			os.Exit(1)
		}
		siteName := os.Args[3]
		err := generator.NewSite(siteName)
		if err != nil {
			fmt.Printf("Error creating site: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Created new site: %s\n", siteName)

	case "sites":
		listSites()

	case "post":
		runWithSiteSelection(func() error {
			ui.RunPostMenu()
			return nil
		})

	case "page":
		runWithSiteSelection(func() error {
			ui.RunPageMenu()
			return nil
		})

	case "theme":
		runWithSiteSelection(func() error {
			ui.RunThemeMenu()
			return nil
		})

	case "font":
		runWithSiteSelection(func() error {
			ui.RunFontMenu()
			return nil
		})

	case "config":
		runWithSiteSelection(func() error {
			ui.RunConfigMenu()
			return nil
		})

	case "build":
		runWithSiteSelection(func() error {
			return generator.BuildSite()
		})

	case "serve":
		runWithSiteSelection(func() error {
			return generator.StartDevServer()
		})

	case "upgrade":
		runWithSiteSelection(func() error {
			return upgrade.RunUpgrade()
		})

	case "help":
		printDetailedHelp()

	default:
		printHelp()
	}
}

func isInBazelSite() bool {
	_, err := os.Stat("bazel.toml")
	return err == nil
}

func printVersion() {
	fmt.Printf("Bazel Static Site Generator v%s\n", Version)
	fmt.Println("A fast and simple static site generator with multi-site support")
	fmt.Println("")
	fmt.Println("Features:")
	fmt.Println("• Multi-site registry and selection")
	fmt.Println("• Interactive Bubbletea-powered UI")
	fmt.Println("• Built-in themes (Pika Beach, Catppuccin, Dracula, Nord, Tokyo Night, 3li7e)")
	fmt.Println("• Auto-rebuild on theme changes")
	fmt.Println("• Live development server")
	fmt.Println("• Markdown post support")
	fmt.Println("• Automatic footer with BazeBlog attribution and social links")
}

func printHelp() {
	fmt.Println("Bazel - Static Site Generator")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  new site <name>   Create a new site")
	fmt.Println("  post              Interactive post management")
	fmt.Println("  page              Interactive page management")
	fmt.Println("  theme             Select site theme")
	fmt.Println("  font              Select site font")
	fmt.Println("  config            Configure site settings")
	fmt.Println("  build             Build the site")
	fmt.Println("  serve             Start dev server")
	fmt.Println("  upgrade           Upgrade site to latest version")
	fmt.Println("  sites             List registered sites")
	fmt.Println("  version           Show version information")
	fmt.Println("  help              Show detailed help")
	fmt.Println("")
	fmt.Println("Run 'bazel help' for detailed documentation.")
}

func printDetailedHelp() {
	fmt.Println("🏗️  Bazel - Static Site Generator")
	fmt.Println("")
	fmt.Println("A fast and simple static site generator with Pika-inspired themes.")
	fmt.Println("")
	fmt.Println("📋 COMMANDS:")
	fmt.Println("")
	fmt.Println("🆕 bazel new site <name>")
	fmt.Println("   Create a new Bazel site with:")
	fmt.Println("   • Sample posts demonstrating features")
	fmt.Println("   • About page with comprehensive documentation")
	fmt.Println("   • Pika Beach theme and Source Serif 4 font")
	fmt.Println("   • Ready-to-use configuration")
	fmt.Println("")
	fmt.Println("📝 bazel post")
	fmt.Println("   Interactive post management menu:")
	fmt.Println("   • New: Create a new post in your default editor")
	fmt.Println("   • Edit: Select and edit existing posts")
	fmt.Println("   • Drafts: Manage draft posts")
	fmt.Println("")
	fmt.Println("📄 bazel page")
	fmt.Println("   Interactive page management menu:")
	fmt.Println("   • New: Create a new page in your default editor")
	fmt.Println("   • Edit: Select and edit existing pages")
	fmt.Println("   • Drafts: Manage draft pages")
	fmt.Println("   • Organize: Reorder page navigation")
	fmt.Println("")
	fmt.Println("🎨 bazel theme")
	fmt.Println("   Interactive theme selector with live preview:")
	fmt.Println("   • pika-beach (default) - Warm, beach-inspired")
	fmt.Println("   • catppuccin themes - Latte, Frappe, Macchiato, Mocha")
	fmt.Println("   • dracula - Classic dark theme")
	fmt.Println("   • nord - Arctic blue theme")
	fmt.Println("   • tokyo-night - Neon night theme")
	fmt.Println("   • 3li7e - Retro green-on-black CRT monitor theme")
	fmt.Println("")
	fmt.Println("🔤 bazel font")
	fmt.Println("   Interactive font selector with preview:")
	fmt.Println("   • pika-serif (default) - Source Serif 4")
	fmt.Println("   • system - System fonts")
	fmt.Println("   • serif, monospace, arial, helvetica, georgia, times")
	fmt.Println("")
	fmt.Println("⚙️  bazel config")
	fmt.Println("   Comprehensive configuration menu:")
	fmt.Println("   • Site Settings: Edit title, description, and domain")
	fmt.Println("   • Theme: Change color scheme with live preview")
	fmt.Println("   • Font: Select typography with preview")
	fmt.Println("   • Social Links: Configure social media profiles")
	fmt.Println("   • All site configuration in one place")
	fmt.Println("")
	fmt.Println("🔧 bazel build")
	fmt.Println("   Build your site for production:")
	fmt.Println("   • Processes all posts and pages")
	fmt.Println("   • Generates CSS with selected theme")
	fmt.Println("   • Creates RSS feed")
	fmt.Println("   • Outputs to public/ directory")
	fmt.Println("")
	fmt.Println("🚀 bazel serve")
	fmt.Println("   Start development server:")
	fmt.Println("   • Serves site at http://localhost:3000")
	fmt.Println("   • Live reload on file changes")
	fmt.Println("   • Perfect for development and preview")
	fmt.Println("")
	fmt.Println("🔄 bazel upgrade")
	fmt.Println("   Upgrade site to latest version:")
	fmt.Println("   • Checks for template changes and new features")
	fmt.Println("   • Updates CSS generation and theme system")
	fmt.Println("   • Maintains backward compatibility")
	fmt.Println("   • Automatically rebuilds site after upgrade")
	fmt.Println("")
	fmt.Println("📁 DIRECTORY STRUCTURE:")
	fmt.Println("")
	fmt.Println("   your-site/")
	fmt.Println("   ├── bazel.toml      # Configuration file")
	fmt.Println("   ├── posts/          # Markdown blog posts")
	fmt.Println("   ├── pages/          # Static HTML pages")
	fmt.Println("   ├── themes/         # Custom themes (optional)")
	fmt.Println("   └── public/         # Generated site output")
	fmt.Println("")
	fmt.Println("📝 WRITING POSTS:")
	fmt.Println("")
	fmt.Println("   Posts are Markdown files with frontmatter:")
	fmt.Println("   ---")
	fmt.Println("   title: Your Post Title")
	fmt.Println("   date: January 1, 2025")
	fmt.Println("   ---")
	fmt.Println("")
	fmt.Println("   ## Your Content Here")
	fmt.Println("   Write your post content in Markdown format.")
	fmt.Println("")
	fmt.Println("🎛️  CONFIGURATION:")
	fmt.Println("")
	fmt.Println("   Edit bazel.toml to customize:")
	fmt.Println("   • Site title and description")
	fmt.Println("   • Theme colors and fonts")
	fmt.Println("   • Base URL for deployment")
	fmt.Println("   • Social media links")
	fmt.Println("")
	fmt.Println("🌟 GETTING STARTED:")
	fmt.Println("")
	fmt.Println("   1. bazel new site my-blog")
	fmt.Println("   2. cd my-blog")
	fmt.Println("   3. bazel serve")
	fmt.Println("   4. Visit http://localhost:3000")
	fmt.Println("   5. bazel post (to create your first post)")
	fmt.Println("   6. bazel build (when ready to deploy)")
	fmt.Println("")
	fmt.Println("💡 TIPS:")
	fmt.Println("")
	fmt.Println("   • Use 'bazel theme' to experiment with different looks")
	fmt.Println("   • The about page is auto-generated with documentation")
	fmt.Println("   • Sample posts show you how to format content")
	fmt.Println("   • All changes are live-reloaded in serve mode")
	fmt.Println("")
	fmt.Println("Happy blogging! 🎉")
}

func runWithSiteSelection(fn func() error) {
	// Check if we're in a Bazel site directory
	if isInBazelSite() {
		// We're in a site directory, run the command directly
		err := fn()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// We're not in a site directory, show site selector
	selectedSite, err := ui.RunSiteSelector()
	if err != nil {
		fmt.Printf("Error selecting site: %v\n", err)
		os.Exit(1)
	}

	if selectedSite == nil {
		fmt.Println("No site selected")
		return
	}

	// Change to the selected site directory
	originalDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	err = os.Chdir(selectedSite.Path)
	if err != nil {
		fmt.Printf("Error changing to site directory: %v\n", err)
		os.Exit(1)
	}

	// Update the site's last used time
	reg, err := registry.LoadRegistry()
	if err == nil {
		reg.UpdateLastUsed(selectedSite.Path)
		reg.Save()
	}

	fmt.Printf("Working with site: %s\n", selectedSite.Name)

	// Run the command
	err = fn()

	// Change back to original directory
	os.Chdir(originalDir)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func listSites() {
	reg, err := registry.LoadRegistry()
	if err != nil {
		fmt.Printf("Error loading site registry: %v\n", err)
		os.Exit(1)
	}

	// Clean up invalid sites
	reg.ValidateSites()
	reg.Save()

	sites := reg.GetSites()
	if len(sites) == 0 {
		fmt.Println("No Bazel sites found in registry.")
		fmt.Println("Create a new site with: bazel new site <name>")
		return
	}

	fmt.Println("🏗️  Registered Bazel Sites:")
	fmt.Println("")

	for i, site := range sites {
		fmt.Printf("%d. %s\n", i+1, site.Name)
		fmt.Printf("   📁 %s\n", site.Path)
		if site.Description != "" {
			fmt.Printf("   📝 %s\n", site.Description)
		}
		fmt.Printf("   🕒 Last used: %s\n", formatLastUsed(site.LastUsed))
		fmt.Println("")
	}

	fmt.Printf("Total: %d site(s)\n", len(sites))
	fmt.Println("")
	fmt.Println("Use 'bazel post', 'bazel build', etc. to work with a site.")
	fmt.Println("If not in a site directory, you'll be prompted to select one.")
}

func formatLastUsed(lastUsed time.Time) string {
	if lastUsed.IsZero() {
		return "never"
	}

	if time.Since(lastUsed) < 24*time.Hour {
		return fmt.Sprintf("%d hours ago", int(time.Since(lastUsed).Hours()))
	}

	return lastUsed.Format("Jan 2, 2006")
}
