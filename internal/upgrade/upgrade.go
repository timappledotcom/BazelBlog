package upgrade

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/yourusername/bazel_blog/internal/config"
	"github.com/yourusername/bazel_blog/internal/generator"
)

// SiteVersion tracks the version of templates and structure used
type SiteVersion struct {
	Version      string            `json:"version"`
	TemplateHash string            `json:"template_hash"`
	LastUpgrade  time.Time         `json:"last_upgrade"`
	Features     map[string]string `json:"features,omitempty"`
}

const CurrentVersion = "1.4.2"
const VersionFile = ".bazel-version"

// Upgrade represents a version upgrade step
type Upgrade struct {
	FromVersion string
	ToVersion   string
	Description string
	Apply       func() error
}

// RunUpgrade checks for and applies any necessary upgrades to the current site
func RunUpgrade() error {
	fmt.Println("🔄 Bazel Site Upgrade")
	fmt.Println("")

	// Check if we're in a bazel site
	if !isInBazelSite() {
		return fmt.Errorf("not in a bazel site directory (bazel.toml not found)")
	}

	// Load current site version
	siteVersion, err := loadSiteVersion()
	if err != nil {
		fmt.Printf("⚠️  No version file found, creating new version tracking\n")
		siteVersion = &SiteVersion{
			Version:     "0.0.0",
			LastUpgrade: time.Time{},
			Features:    make(map[string]string),
		}
	}

	fmt.Printf("� Currenit site version: %s\n", siteVersion.Version)
	fmt.Printf("📦 Bazel version: %s\n", CurrentVersion)

	// Check if already up to date
	if compareVersions(siteVersion.Version, CurrentVersion) >= 0 {
		fmt.Println("✅ Site is already up to date!")
		return nil
	}

	fmt.Println("")
	fmt.Println("🔍 Checking for upgrades...")

	// Define all available upgrades
	upgrades := []Upgrade{
		{
			FromVersion: "0.0.0",
			ToVersion:   "1.1.0",
			Description: "Add version tracking and theme improvements",
			Apply:       upgradeToV1_1_0,
		},
		{
			FromVersion: "1.1.0",
			ToVersion:   "1.1.5",
			Description: "Remove dark mode media query interference",
			Apply:       upgradeToV1_1_5,
		},
		{
			FromVersion: "1.1.5",
			ToVersion:   "1.1.7",
			Description: "Enhanced UI with colorful menus and improved navigation spacing",
			Apply:       upgradeToV1_1_7,
		},
		{
			FromVersion: "1.1.7",
			ToVersion:   "1.1.8",
			Description: "Added 3li7e retro CRT theme and enhanced post/page editing functionality",
			Apply:       upgradeToV1_1_8,
		},
		{
			FromVersion: "1.1.8",
			ToVersion:   "1.4.0",
			Description: "Added comprehensive markdown documentation and improved user experience",
			Apply:       upgradeToV1_4_0,
		},
		{
			FromVersion: "1.4.0",
			ToVersion:   "1.4.1",
			Description: "Project cleanup and build system improvements",
			Apply:       upgradeToV1_4_1,
		},
		{
			FromVersion: "1.4.1",
			ToVersion:   "1.4.2",
			Description: "Improved site structure with organized directories",
			Apply:       upgradeToV1_4_2,
		},
	}

	// Apply upgrades sequentially
	applied := false
	currentVersion := siteVersion.Version

	for _, upgrade := range upgrades {
		if shouldApplyUpgrade(currentVersion, upgrade) {
			fmt.Printf("🔧 Applying upgrade: %s\n", upgrade.Description)
			fmt.Printf("   %s → %s\n", upgrade.FromVersion, upgrade.ToVersion)

			err := upgrade.Apply()
			if err != nil {
				return fmt.Errorf("failed to apply upgrade %s: %w", upgrade.ToVersion, err)
			}

			currentVersion = upgrade.ToVersion
			applied = true
			fmt.Printf("✅ Upgrade to %s completed\n", upgrade.ToVersion)
			fmt.Println("")
		}
	}

	if applied {
		// Update version file with final version
		siteVersion.Version = currentVersion
		siteVersion.LastUpgrade = time.Now()
		err = saveSiteVersion(siteVersion)
		if err != nil {
			return fmt.Errorf("failed to save version info: %w", err)
		}

		// Rebuild site with new templates
		fmt.Println("🏗️  Rebuilding site with updated templates...")
		err = generator.BuildSite()
		if err != nil {
			return fmt.Errorf("failed to rebuild site: %w", err)
		}

		fmt.Println("🎉 Upgrade completed successfully!")
		fmt.Printf("   Site updated to version %s\n", currentVersion)
	} else {
		fmt.Println("ℹ️  No upgrades needed")
	}

	return nil
}

// shouldApplyUpgrade determines if an upgrade should be applied
func shouldApplyUpgrade(currentVersion string, upgrade Upgrade) bool {
	// Current version must be >= FromVersion and < ToVersion
	return compareVersions(currentVersion, upgrade.FromVersion) >= 0 &&
		compareVersions(currentVersion, upgrade.ToVersion) < 0
}

// compareVersions compares two semantic version strings
// Returns: -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2
func compareVersions(v1, v2 string) int {
	if v1 == v2 {
		return 0
	}

	// Handle special case for 0.0.0
	if v1 == "0.0.0" && v2 != "0.0.0" {
		return -1
	}
	if v2 == "0.0.0" && v1 != "0.0.0" {
		return 1
	}

	// Parse version strings
	parts1 := parseVersion(v1)
	parts2 := parseVersion(v2)

	// Compare each part (major, minor, patch)
	for i := 0; i < 3; i++ {
		if parts1[i] < parts2[i] {
			return -1
		}
		if parts1[i] > parts2[i] {
			return 1
		}
	}

	return 0
}

// parseVersion parses a semantic version string into [major, minor, patch]
func parseVersion(version string) [3]int {
	parts := strings.Split(version, ".")
	result := [3]int{0, 0, 0}

	for i := 0; i < len(parts) && i < 3; i++ {
		if num, err := strconv.Atoi(parts[i]); err == nil {
			result[i] = num
		}
	}

	return result
}

// Upgrade functions for each version

func upgradeToV1_1_0() error {
	fmt.Println("   • Adding version tracking")
	// Version tracking is handled by the upgrade system itself
	return nil
}

func upgradeToV1_1_5() error {
	fmt.Println("   • Updating CSS generation (removing dark mode conflicts)")
	fmt.Println("   • Theme selection improvements applied")
	// CSS will be regenerated on next build
	return nil
}

func upgradeToV1_1_7() error {
	fmt.Println("   • Enhanced colorful menu interface")
	fmt.Println("   • Added input field cursor indicators")
	fmt.Println("   • Improved screen clearing for clean navigation")
	fmt.Println("   • Updated navigation spacing")
	// UI improvements are in the binary itself
	return nil
}

func upgradeToV1_1_8() error {
	fmt.Println("   • Added 3li7e retro CRT monitor theme (green-on-black)")
	fmt.Println("   • Enhanced post and page editing functionality")
	fmt.Println("   • Improved markdown page support")
	// Theme and editing improvements are in the binary
	return nil
}

func upgradeToV1_4_0() error {
	fmt.Println("   • Added comprehensive markdown documentation")
	fmt.Println("   • Improved user experience with new docs and guides")
	// Documentation improvements
	return nil
}

func upgradeToV1_4_1() error {
	fmt.Println("   • Project cleanup and build system improvements")
	fmt.Println("   • Enhanced documentation structure")
	fmt.Println("   • Improved build system with Makefile")
	fmt.Println("   • Better developer experience and contribution workflow")

	// Convert JSON config to TOML if needed
	err := convertJSONConfigToTOML()
	if err != nil {
		return fmt.Errorf("failed to convert config format: %w", err)
	}

	return nil
}

func upgradeToV1_4_2() error {
	fmt.Println("   • Migrating to organized directory structure")
	fmt.Println("   • Moving posts and pages to subdirectories")

	// Migrate existing public directory structure
	err := migrateDirectoryStructure()
	if err != nil {
		return fmt.Errorf("failed to migrate directory structure: %w", err)
	}

	fmt.Println("   • Site structure updated for better organization")
	return nil
}

// migrateDirectoryStructure moves existing posts and pages to subdirectories
func migrateDirectoryStructure() error {
	publicDir := "public"

	// Check if public directory exists
	if _, err := os.Stat(publicDir); os.IsNotExist(err) {
		// No public directory to migrate
		return nil
	}

	// Create subdirectories if they don't exist
	postsDir := filepath.Join(publicDir, "posts")
	pagesDir := filepath.Join(publicDir, "pages")

	if err := os.MkdirAll(postsDir, 0755); err != nil {
		return fmt.Errorf("failed to create posts directory: %w", err)
	}
	if err := os.MkdirAll(pagesDir, 0755); err != nil {
		return fmt.Errorf("failed to create pages directory: %w", err)
	}

	// Read all files in public directory
	files, err := os.ReadDir(publicDir)
	if err != nil {
		return fmt.Errorf("failed to read public directory: %w", err)
	}

	movedCount := 0

	for _, file := range files {
		if file.IsDir() {
			continue // Skip subdirectories
		}

		filename := file.Name()

		// Skip core files that should stay in root
		if filename == "index.html" || filename == "style.css" || filename == "feed.xml" {
			continue
		}

		// Check if it's a post or page by looking at source directories
		oldPath := filepath.Join(publicDir, filename)
		var newPath string

		// Determine if it's a post or page by checking source files
		if isPostFile(filename) {
			newPath = filepath.Join(postsDir, filename)
			fmt.Printf("   • Moving post: %s → posts/%s\n", filename, filename)
		} else if isPageFile(filename) {
			newPath = filepath.Join(pagesDir, filename)
			fmt.Printf("   • Moving page: %s → pages/%s\n", filename, filename)
		} else {
			// Unknown file type, leave it in root
			continue
		}

		// Move the file
		err = os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("   ⚠️  Warning: Failed to move %s: %v\n", filename, err)
			continue
		}

		movedCount++
	}

	if movedCount > 0 {
		fmt.Printf("   • Successfully migrated %d files to organized structure\n", movedCount)
	} else {
		fmt.Println("   • No files needed migration")
	}

	return nil
}

// isPostFile checks if a filename corresponds to a post
func isPostFile(filename string) bool {
	if !strings.HasSuffix(filename, ".html") {
		return false
	}

	// Check if corresponding markdown file exists in posts directory
	mdName := strings.Replace(filename, ".html", ".md", 1)
	if _, err := os.Stat(filepath.Join("posts", mdName)); err == nil {
		return true
	}

	// Fallback: check common post patterns
	// This is a heuristic for files that might be posts
	return true // Default to treating HTML files as posts unless proven otherwise
}

// isPageFile checks if a filename corresponds to a page
func isPageFile(filename string) bool {
	if !strings.HasSuffix(filename, ".html") {
		return false
	}

	// Check if corresponding markdown file exists in pages directory
	mdName := strings.Replace(filename, ".html", ".md", 1)
	if _, err := os.Stat(filepath.Join("pages", mdName)); err == nil {
		return true
	}

	// Check if HTML file exists in pages directory
	if _, err := os.Stat(filepath.Join("pages", filename)); err == nil {
		return true
	}

	// Common page names
	pagenames := []string{"about.html", "contact.html", "projects.html", "resume.html"}
	for _, pagename := range pagenames {
		if filename == pagename {
			return true
		}
	}

	return false
}

// convertJSONConfigToTOML converts a JSON format bazel.toml to proper TOML format
func convertJSONConfigToTOML() error {
	configPath := "bazel.toml"

	// Read the current config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil // No config file to convert
	}

	// Check if it's JSON format (starts with '{')
	content := strings.TrimSpace(string(data))
	if !strings.HasPrefix(content, "{") {
		// Already in TOML format
		return nil
	}

	fmt.Println("   • Converting JSON config to TOML format")

	// Create backup first
	backupPath := fmt.Sprintf("bazel.toml.json-backup.%s", time.Now().Format("20060102-150405"))
	err = os.WriteFile(backupPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}
	fmt.Printf("   • JSON config backed up to %s\n", backupPath)

	// Parse JSON config
	var jsonConfig struct {
		SiteName    string `json:"site_name"`
		Title       string `json:"title"`
		Description string `json:"description"`
		BaseURL     string `json:"base_url"`
		Theme       struct {
			ColorScheme string `json:"color_scheme"`
			Font        string `json:"font"`
		} `json:"theme"`
		Socials map[string]string `json:"socials"`
		Editor  string            `json:"editor"`
	}

	err = json.Unmarshal(data, &jsonConfig)
	if err != nil {
		return fmt.Errorf("failed to parse JSON config: %w", err)
	}

	// Create new config struct
	cfg := &config.Config{
		SiteName:    jsonConfig.SiteName,
		Title:       jsonConfig.Title,
		Description: jsonConfig.Description,
		BaseURL:     jsonConfig.BaseURL,
		Theme: config.ThemeConfig{
			ColorScheme: jsonConfig.Theme.ColorScheme,
			Font:        jsonConfig.Theme.Font,
		},
		Socials: jsonConfig.Socials,
		Editor:  jsonConfig.Editor,
	}

	// Ensure socials map is initialized
	if cfg.Socials == nil {
		cfg.Socials = make(map[string]string)
	}

	// Save as TOML
	err = cfg.Save()
	if err != nil {
		return fmt.Errorf("failed to save TOML config: %w", err)
	}

	fmt.Println("   • Successfully converted config from JSON to TOML format")
	return nil
}

// Utility functions

func isInBazelSite() bool {
	_, err := os.Stat("bazel.toml")
	return err == nil
}

func loadSiteVersion() (*SiteVersion, error) {
	data, err := os.ReadFile(VersionFile)
	if err != nil {
		return nil, err
	}

	var version SiteVersion
	err = json.Unmarshal(data, &version)
	if err != nil {
		return nil, err
	}

	return &version, nil
}

func saveSiteVersion(version *SiteVersion) error {
	data, err := json.MarshalIndent(version, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(VersionFile, data, 0644)
}

// CheckSiteVersion returns the current site version
func CheckSiteVersion() (string, error) {
	version, err := loadSiteVersion()
	if err != nil {
		return "unknown", err
	}
	return version.Version, nil
}

// BackupConfig creates a backup of the current configuration
func BackupConfig() error {
	configPath := "bazel.toml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil // No config to backup
	}

	backupPath := fmt.Sprintf("bazel.toml.backup.%s", time.Now().Format("20060102-150405"))

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = os.WriteFile(backupPath, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("   • Configuration backed up to %s\n", backupPath)
	return nil
}

// UpgradeConfig updates configuration file with new features
func UpgradeConfig() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Add any new configuration fields or update defaults
	changed := false

	// Validate theme options
	validSchemes := map[string]bool{
		"pika-beach": true, "catppuccin-latte": true, "catppuccin-frappe": true,
		"catppuccin-macchiato": true, "catppuccin-mocha": true, "dracula": true,
		"nord": true, "tokyo-night": true, "3li7e": true,
	}

	if !validSchemes[cfg.Theme.ColorScheme] {
		fmt.Printf("   • Updating invalid theme '%s' to 'pika-beach'\n", cfg.Theme.ColorScheme)
		cfg.Theme.ColorScheme = "pika-beach"
		changed = true
	}

	if changed {
		err = cfg.Save()
		if err != nil {
			return err
		}
		fmt.Println("   • Configuration updated with new options")
	}

	return nil
}
