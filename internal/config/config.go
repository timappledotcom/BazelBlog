package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	SiteName    string            `toml:"site_name"`
	Title       string            `toml:"title"`
	Description string            `toml:"description"`
	BaseURL     string            `toml:"base_url"`
	Theme       ThemeConfig       `toml:"theme"`
	Socials     map[string]string `toml:"socials"`
	Editor      string            `toml:"editor"`
}

type ThemeConfig struct {
	ColorScheme string `toml:"color_scheme"`
	Font        string `toml:"font"`
}

var DefaultConfig = Config{
	Title:       "My Bazel Site",
	Description: "A static site generated with Bazel",
	BaseURL:     "https://example.com",
	Theme: ThemeConfig{
		ColorScheme: "pika-beach",
		Font:        "pika-serif",
	},
	Socials: make(map[string]string),
	Editor:  "auto",
}

// Available color schemes
var ColorSchemes = []string{
	"pika-beach",
	"catppuccin-latte",
	"catppuccin-frappe",
	"catppuccin-macchiato",
	"catppuccin-mocha",
	"dracula",
	"nord",
	"tokyo-night",
	"3li7e",
}

// Available fonts
var Fonts = []string{
	"pika-serif",
	"system",
	"serif",
	"monospace",
	"arial",
	"helvetica",
	"georgia",
	"times",
}

// Available editors
var Editors = []string{
	"auto",
	"vim",
	"nvim",
	"nano",
	"emacs",
	"hx",
	"vi",
	"code",
	"subl",
	"atom",
	"gedit",
	"kate",
}

// Available social platforms
var SocialPlatforms = []string{
	"twitter",
	"github",
	"linkedin",
	"facebook",
	"instagram",
	"youtube",
	"mastodon",
	"email",
}

func LoadConfig() (*Config, error) {
	configPath := "bazel.toml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Return default config if no config file exists
		config := DefaultConfig
		return &config, nil
	}

	var config Config
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

func (c *Config) Save() error {
	file, err := os.Create("bazel.toml")
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}

	return nil
}

func (c *Config) SetColorScheme(scheme string) {
	c.Theme.ColorScheme = scheme
}

func (c *Config) SetFont(font string) {
	c.Theme.Font = font
}

func (c *Config) SetSocial(platform, url string) {
	if c.Socials == nil {
		c.Socials = make(map[string]string)
	}
	c.Socials[platform] = url
}

func (c *Config) RemoveSocial(platform string) {
	delete(c.Socials, platform)
}

func (c *Config) SetEditor(editor string) {
	c.Editor = editor
}

func (c *Config) GetEditor() string {
	if c.Editor == "" || c.Editor == "auto" {
		// Try to detect from environment variable first
		if envEditor := os.Getenv("EDITOR"); envEditor != "" {
			return envEditor
		}
		// Fallback to vi
		return "vi"
	}
	return c.Editor
}

func (c *Config) GetCSSVariables() string {
	variables := ""

	// Color scheme variables
	switch c.Theme.ColorScheme {
	case "pika-beach":
		variables += "--bg-color: 255, 252, 245; --text-color: 40, 59, 67; --accent-color: #048AA2; --secondary-color: #6c6f85;"
	case "catppuccin-latte":
		variables += "--bg-color: 239, 241, 245; --text-color: 76, 79, 105; --accent-color: #8839ef; --secondary-color: #6c6f85;"
	case "catppuccin-frappe":
		variables += "--bg-color: 48, 52, 70; --text-color: 198, 208, 245; --accent-color: #ca9ee6; --secondary-color: #838ba7;"
	case "catppuccin-macchiato":
		variables += "--bg-color: 36, 39, 58; --text-color: 202, 211, 245; --accent-color: #c6a0f6; --secondary-color: #8087a2;"
	case "catppuccin-mocha":
		variables += "--bg-color: 30, 30, 46; --text-color: 205, 214, 244; --accent-color: #cba6f7; --secondary-color: #7f849c;"
	case "dracula":
		variables += "--bg-color: 40, 42, 54; --text-color: 248, 248, 242; --accent-color: #bd93f9; --secondary-color: #6272a4;"
	case "nord":
		variables += "--bg-color: 46, 52, 64; --text-color: 216, 222, 233; --accent-color: #88c0d0; --secondary-color: #4c566a;"
	case "tokyo-night":
		variables += "--bg-color: 26, 27, 38; --text-color: 169, 177, 214; --accent-color: #7aa2f7; --secondary-color: #565f89;"
	case "3li7e":
		variables += "--bg-color: 0, 0, 0; --text-color: 0, 255, 0; --accent-color: #00ff41; --secondary-color: #008f11;"
	default:
		variables += "--bg-color: 255, 252, 245; --text-color: 40, 59, 67; --accent-color: #048AA2; --secondary-color: #6c6f85;"
	}

	// Font variables
	switch c.Theme.Font {
	case "pika-serif":
		variables += " --font-family: 'Source Serif 4', Georgia, serif;"
	case "serif":
		variables += " --font-family: Georgia, serif;"
	case "monospace":
		variables += " --font-family: 'Courier New', monospace;"
	case "arial":
		variables += " --font-family: Arial, sans-serif;"
	case "helvetica":
		variables += " --font-family: 'Helvetica Neue', Helvetica, sans-serif;"
	case "georgia":
		variables += " --font-family: Georgia, serif;"
	case "times":
		variables += " --font-family: 'Times New Roman', Times, serif;"
	default:
		variables += " --font-family: system-ui, -apple-system, sans-serif;"
	}

	return variables
}
