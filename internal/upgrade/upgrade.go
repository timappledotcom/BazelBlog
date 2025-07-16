package upgrade

import (
	"encoding/json"
	"fmt"
	"os"
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

const CurrentVersion = "1.4.1"
const VersionFile = ".bazel-version"

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
		fmt.Printf("⚠️  No version file found, assuming new site structure needed\n")
		siteVersion = &SiteVersion{
			Version:     "0.0.0",
			LastUpgrade: time.Time{},
			Features:    make(map[string]string),
		}
	}

	fmt.Printf("📊 Current site version: %s\n", siteVersion.Version)
	fmt.Printf("📦 Bazel version: %s\n", CurrentVersion)

	if siteVersion.Version == CurrentVersion {
		fmt.Println("✅ Site is already up to date!")
		return nil
	}

	fmt.Println("")
	fmt.Println("🔍 Checking for upgrades...")

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
	}

	applied := false
	for _, upgrade := range upgrades {
		if shouldApplyUpgrade(siteVersion.Version, upgrade) {
			fmt.Printf("🔧 Applying upgrade: %s\n", upgrade.Description)
			fmt.Printf("   %s → %s\n", upgrade.FromVersion, upgrade.ToVersion)

			err := upgrade.Apply()
			if err != nil {
				return fmt.Errorf("failed to apply upgrade %s: %w", upgrade.ToVersion, err)
			}

			siteVersion.Version = upgrade.ToVersion
			applied = true
			fmt.Printf("✅ Upgrade to %s completed\n", upgrade.ToVersion)
			fmt.Println("")
		}
	}

	if applied {
		// Update version file
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
		fmt.Printf("   Site updated to version %s\n", CurrentVersion)
	} else {
		fmt.Println("ℹ️  No upgrades needed")
	}

	return nil
}

// Upgrade represents a version upgrade step
type Upgrade struct {
	FromVersion string
	ToVersion   string
	Description string
	Apply       func() error
}

// shouldApplyUpgrade determines if an upgrade should be applied
func shouldApplyUpgrade(currentVersion string, upgrade Upgrade) bool {
	return versionCompare(currentVersion, upgrade.FromVersion) >= 0 &&
		versionCompare(currentVersion, upgrade.ToVersion) < 0
}

// versionCompare compares two version strings (simple implementation)
func versionCompare(v1, v2 string) int {
	if v1 == v2 {
		return 0
	}
	if v1 == "0.0.0" {
		return -1
	}
	if v2 == "0.0.0" {
		return 1
	}
	if v1 < v2 {
		return -1
	}
	return 1
}

// upgradeToV1_1_0 adds version tracking to existing sites
func upgradeToV1_1_0() error {
	fmt.Println("   • Adding version tracking")

	// Create .bazel-version file if it doesn't exist
	if _, err := os.Stat(VersionFile); os.IsNotExist(err) {
		version := &SiteVersion{
			Version:     "1.1.0",
			LastUpgrade: time.Now(),
			Features:    make(map[string]string),
		}
		return saveSiteVersion(version)
	}

	return nil
}

// upgradeToV1_1_5 removes dark mode media query and updates CSS generation
func upgradeToV1_1_5() error {
	fmt.Println("   • Updating CSS generation (removing dark mode conflicts)")

	// Check if public/style.css exists and has the problematic media query
	cssPath := "public/style.css"
	if _, err := os.Stat(cssPath); err == nil {
		content, err := os.ReadFile(cssPath)
		if err == nil && strings.Contains(string(content), "@media (prefers-color-scheme: dark)") {
			fmt.Println("   • Found conflicting dark mode CSS - will be fixed on next build")
		}
	}

	fmt.Println("   • Theme selection improvements applied")
	return nil
}

// upgradeToV1_1_7 adds enhanced UI features and improved navigation spacing
func upgradeToV1_1_7() error {
	fmt.Println("   • Enhanced colorful menu interface")
	fmt.Println("   • Added input field cursor indicators")
	fmt.Println("   • Improved screen clearing for clean navigation")
	fmt.Println("   • Updated navigation spacing (reduced gap between menu and content)")

	// This upgrade primarily affects the UI/menu system, not site templates
	// The changes are in the binary itself, so no file modifications needed
	// Just ensure the site will be rebuilt with any updated CSS generation

	return nil
}

// upgradeToV1_1_8 adds 3li7e theme and enhanced editing functionality
func upgradeToV1_1_8() error {
	fmt.Println("   • Added 3li7e retro CRT monitor theme (green-on-black)")
	fmt.Println("   • Enhanced post and page editing functionality")
	fmt.Println("   • Improved markdown page support")
	fmt.Println("   • Updated about page template with new theme information")

	// This upgrade primarily adds new theme support and editing features
	// The changes are in the binary itself and templates, so no file modifications needed
	// The site will be rebuilt with the updated theme options

	return nil
}

// upgradeToV1_4_0 implements new markdown documentation upgrade
func upgradeToV1_4_0() error {
	fmt.Println("   • Added markdown documentation file for users")
	fmt.Println("   • Improved user experience with new docs and guides")

	// Add logic here related to version 1.4.0
	// Normally, it would modify site templates, config, or files

	fmt.Println("   • Ensure the site rebuild reflects the new version")

	return nil
}

// upgradeToV1_4_1 implements project cleanup and build system improvements
func upgradeToV1_4_1() error {
	fmt.Println("   • Project cleanup and build system improvements")
	fmt.Println("   • Enhanced documentation structure")
	fmt.Println("   • Improved build system with Makefile")
	fmt.Println("   • Better developer experience and contribution workflow")

	// This upgrade primarily affects the development workflow and project structure
	// No site template changes needed, just ensure rebuild with latest version

	return nil
}

// isInBazelSite checks if current directory is a bazel site
func isInBazelSite() bool {
	_, err := os.Stat("bazel.toml")
	return err == nil
}

// loadSiteVersion loads the site version information
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

// saveSiteVersion saves the site version information
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

	// Example: Add new theme options if they don't exist
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
