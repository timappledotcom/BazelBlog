package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/yourusername/bazel_blog/internal/config"
	"github.com/yourusername/bazel_blog/internal/generator"
)

// getFontStyle returns a stylized version of text for the given font
func getFontStyle(font, text string) string {
	switch font {
	case "serif":
		return "𝒮 " + text + " (serif)"
	case "monospace":
		return "🄼 " + text + " (monospace)"
	case "arial":
		return "🄰 " + text + " (arial)"
	case "helvetica":
		return "🄷 " + text + " (helvetica)"
	case "georgia":
		return "🄶 " + text + " (georgia)"
	case "times":
		return "🅃 " + text + " (times)"
	case "system":
		return "🅂 " + text + " (system)"
	default:
		return text
	}
}

// getFontPreview returns a preview string for the font
func getFontPreview(font string) string {
	switch font {
	case "serif":
		return "The quick brown fox jumps over the lazy dog (Serif)"
	case "monospace":
		return "The quick brown fox jumps over the lazy dog (Mono)"
	case "arial":
		return "The quick brown fox jumps over the lazy dog (Arial)"
	case "helvetica":
		return "The quick brown fox jumps over the lazy dog (Helvetica)"
	case "georgia":
		return "The quick brown fox jumps over the lazy dog (Georgia)"
	case "times":
		return "The quick brown fox jumps over the lazy dog (Times)"
	case "system":
		return "The quick brown fox jumps over the lazy dog (System)"
	default:
		return "The quick brown fox jumps over the lazy dog"
	}
}

// getThemeColors returns ANSI color codes for the given theme
func getThemeColors(theme string) (bgColor, textColor, accentColor string) {
	switch theme {
	case "catppuccin-latte":
		return "\033[48;2;239;241;245m", "\033[38;2;76;79;105m", "\033[38;2;136;57;239m"
	case "catppuccin-frappe":
		return "\033[48;2;48;52;70m", "\033[38;2;198;208;245m", "\033[38;2;202;158;230m"
	case "catppuccin-macchiato":
		return "\033[48;2;36;39;58m", "\033[38;2;202;211;245m", "\033[38;2;198;160;246m"
	case "catppuccin-mocha":
		return "\033[48;2;30;30;46m", "\033[38;2;205;214;244m", "\033[38;2;203;166;247m"
	case "dracula":
		return "\033[48;2;40;42;54m", "\033[38;2;248;248;242m", "\033[38;2;189;147;249m"
	case "nord":
		return "\033[48;2;46;52;64m", "\033[38;2;216;222;233m", "\033[38;2;136;192;208m"
	case "tokyo-night":
		return "\033[48;2;26;27;38m", "\033[38;2;169;177;214m", "\033[38;2;122;162;247m"
	default:
		return "\033[48;2;239;241;245m", "\033[38;2;76;79;105m", "\033[38;2;136;57;239m"
	}
}

// getThemeStyle applies theme colors to text
func getThemeStyle(theme, text string) string {
	bgColor, textColor, _ := getThemeColors(theme)
	reset := "\033[0m"
	return bgColor + textColor + text + reset
}

// getThemeAccentStyle applies theme accent color to text
func getThemeAccentStyle(theme, text string) string {
	_, _, accentColor := getThemeColors(theme)
	reset := "\033[0m"
	return accentColor + text + reset
}

// getThemePreview returns a colorized preview for the theme
func getThemePreview(theme string) string {
	bgColor, textColor, accentColor := getThemeColors(theme)
	reset := "\033[0m"

	switch theme {
	case "catppuccin-latte":
		return bgColor + textColor + "☀️ Light & Warm" + reset + " " + accentColor + "Purple accents" + reset
	case "catppuccin-frappe":
		return bgColor + textColor + "🌙 Medium Dark" + reset + " " + accentColor + "Soft purple" + reset
	case "catppuccin-macchiato":
		return bgColor + textColor + "🌃 Dark & Cozy" + reset + " " + accentColor + "Vibrant purple" + reset
	case "catppuccin-mocha":
		return bgColor + textColor + "🌌 Darkest" + reset + " " + accentColor + "Beautiful purple" + reset
	case "dracula":
		return bgColor + textColor + "🧛 Classic Dark" + reset + " " + accentColor + "Purple magic" + reset
	case "nord":
		return bgColor + textColor + "🏔️ Arctic Blue" + reset + " " + accentColor + "Cool blues" + reset
	case "tokyo-night":
		return bgColor + textColor + "🏙️ Neon Night" + reset + " " + accentColor + "Electric blue" + reset
	default:
		return "Preview not available"
	}
}

// Color constants for general UI
const (
	ColorReset         = "\033[0m"
	ColorBold          = "\033[1m"
	ColorDim           = "\033[2m"
	ColorRed           = "\033[31m"
	ColorGreen         = "\033[32m"
	ColorYellow        = "\033[33m"
	ColorBlue          = "\033[34m"
	ColorMagenta       = "\033[35m"
	ColorCyan          = "\033[36m"
	ColorWhite         = "\033[37m"
	ColorBrightRed     = "\033[91m"
	ColorBrightGreen   = "\033[92m"
	ColorBrightYellow  = "\033[93m"
	ColorBrightBlue    = "\033[94m"
	ColorBrightMagenta = "\033[95m"
	ColorBrightCyan    = "\033[96m"
	ColorBrightWhite   = "\033[97m"
)

// Style helper functions
func colorize(color, text string) string {
	return color + text + ColorReset
}

func bold(text string) string {
	return ColorBold + text + ColorReset
}

func dim(text string) string {
	return ColorDim + text + ColorReset
}

func formatTitle(text string) string {
	return colorize(ColorBrightCyan, bold("🏗️  "+text))
}

func formatSubtitle(text string) string {
	return colorize(ColorBrightBlue, text)
}

func formatSuccess(text string) string {
	return colorize(ColorBrightGreen, "✅ "+text)
}

func formatError(text string) string {
	return colorize(ColorBrightRed, "❌ "+text)
}

func formatCursor(selected bool) string {
	if selected {
		return colorize(ColorBrightYellow, "▶ ")
	}
	return "  "
}

func formatMenuItem(text string, selected bool) string {
	if selected {
		return colorize(ColorBrightWhite, bold(text))
	}
	return colorize(ColorWhite, text)
}

func formatInputField(label, value string) string {
	return colorize(ColorBrightBlue, label) + colorize(ColorBrightWhite, value) + colorize(ColorBrightYellow, "|")
}

func formatInstruction(text string) string {
	return colorize(ColorDim, text)
}

// clearScreen clears the terminal screen
func clearScreen() string {
	return "\033[2J\033[H"
}

type MenuState int

const (
	MainMenu MenuState = iota
	ConfigMenu
	ThemeMenu
	FontMenu
	EditorMenu
	SocialMenu
	SocialEditMenu
	SiteSettingsMenu
	TitleEditMenu
	DomainEditMenu
	DescriptionEditMenu
	PostTitleInputMenu
	PageTitleInputMenu
	PostEditMenu
	PageEditMenu
)

type model struct {
	config             *config.Config
	state              MenuState
	cursor             int
	choices            []string
	socialPlatforms    []string
	editingSocial      string
	editingURL         string
	editingTitle       string
	editingDomain      string
	editingDescription string
	message            string
	previewFont        string // Font being previewed in font menu
	previewTheme       string // Theme being previewed in theme menu
}

func RunPostMenu() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	m := model{
		config: cfg,
		state:  MainMenu,
		choices: []string{
			"New Post",
			"Edit Post",
			"Draft Posts",
			"Done",
		},
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running Post Menu: %v\n", err)
	}
}

func RunPageMenu() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	m := model{
		config: cfg,
		state:  MainMenu,
		choices: []string{
			"New Page",
			"Edit Page",
			"Draft Pages",
			"Organize Pages",
			"Done",
		},
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running Page Menu: %v\n", err)
	}
}

func RunThemeMenu() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	m := model{
		config:  cfg,
		state:   ThemeMenu,
		choices: config.ColorSchemes,
	}

	// Set cursor to current theme
	for i, scheme := range config.ColorSchemes {
		if scheme == cfg.Theme.ColorScheme {
			m.cursor = i
			m.previewTheme = scheme
			break
		}
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running Theme Menu: %v\n", err)
	}
}

func RunFontMenu() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	m := model{
		config:  cfg,
		state:   FontMenu,
		choices: config.Fonts,
	}

	// Set cursor to current font
	for i, font := range config.Fonts {
		if font == cfg.Theme.Font {
			m.cursor = i
			m.previewFont = font
			break
		}
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running Font Menu: %v\n", err)
	}
}

func RunConfigMenu() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	m := model{
		config:          cfg,
		state:           ConfigMenu,
		cursor:          0,
		socialPlatforms: config.SocialPlatforms,
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running Config Menu: %v\n", err)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case MainMenu:
			return m.updateMainMenu(msg)
		case ConfigMenu:
			return m.updateConfigMenu(msg)
		case ThemeMenu:
			return m.updateThemeMenu(msg)
		case FontMenu:
			return m.updateFontMenu(msg)
		case EditorMenu:
			return m.updateEditorMenu(msg)
		case SocialMenu:
			return m.updateSocialMenu(msg)
		case SocialEditMenu:
			return m.updateSocialEditMenu(msg)
		case SiteSettingsMenu:
			return m.updateSiteSettingsMenu(msg)
		case TitleEditMenu:
			return m.updateTitleEditMenu(msg)
		case DomainEditMenu:
			return m.updateDomainEditMenu(msg)
		case DescriptionEditMenu:
			return m.updateDescriptionEditMenu(msg)
		case PostTitleInputMenu:
			return m.updatePostTitleInputMenu(msg)
		case PageTitleInputMenu:
			return m.updatePageTitleInputMenu(msg)
		case PostEditMenu:
			return m.updatePostEditMenu(msg)
		case PageEditMenu:
			return m.updatePageEditMenu(msg)
		}
	}

	return m, nil
}

func (m model) updateMainMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		}
	case "enter":
		// Check what menu we're in based on the choices
		if len(m.choices) == 4 && m.choices[0] == "New Post" {
			// Post menu
			switch m.cursor {
			case 0: // New Post
				m.editingTitle = ""
				m.state = PostTitleInputMenu
			case 1: // Edit Post
				// Load available posts and switch to edit menu
				posts, err := generator.ListPosts()
				if err != nil {
					m.message = fmt.Sprintf("Error loading posts: %v", err)
				} else if len(posts) == 0 {
					m.message = "No posts found to edit"
				} else {
					m.choices = posts
					m.cursor = 0
					m.state = PostEditMenu
				}
			case 2: // Draft Posts
				m.message = "Draft Posts functionality coming soon"
			case 3: // Done
				return m, tea.Quit
			}
		} else if len(m.choices) == 5 && m.choices[0] == "New Page" {
			// Page menu
			switch m.cursor {
			case 0: // New Page
				m.editingTitle = ""
				m.state = PageTitleInputMenu
			case 1: // Edit Page
				// Load available pages and switch to edit menu
				pages, err := generator.ListPages()
				if err != nil {
					m.message = fmt.Sprintf("Error loading pages: %v", err)
				} else if len(pages) == 0 {
					m.message = "No pages found to edit"
				} else {
					m.choices = pages
					m.cursor = 0
					m.state = PageEditMenu
				}
			case 2: // Draft Pages
				m.message = "Draft Pages functionality coming soon"
			case 3: // Organize Pages
				m.message = "Organize Pages functionality coming soon"
			case 4: // Done
				return m, tea.Quit
			}
		} else {
			// Default config menu
			switch m.cursor {
			case 0: // Configuration
				m.state = ConfigMenu
				m.cursor = 0
			case 1: // Build Site
				err := generator.BuildSite()
				if err != nil {
					m.message = fmt.Sprintf("Error building site: %v", err)
				} else {
					m.message = "Site built successfully!"
				}
			case 2: // Start Dev Server
				err := generator.StartDevServer()
				if err != nil {
					m.message = fmt.Sprintf("Error starting dev server: %v", err)
				} else {
					m.message = "Dev server started on http://localhost:3000"
				}
			case 3: // Quit
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) updateThemeMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "esc":
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
			m.previewTheme = config.ColorSchemes[m.cursor] // Update preview theme
		}
	case "down", "j":
		if m.cursor < len(config.ColorSchemes)-1 {
			m.cursor++
			m.previewTheme = config.ColorSchemes[m.cursor] // Update preview theme
		}
	case "enter":
		selected := config.ColorSchemes[m.cursor]
		fmt.Printf("\nSetting theme to: %s\n", selected)
		m.config.SetColorScheme(selected)
		fmt.Printf("Config updated, now saving...\n")
		err := m.config.Save()
		if err != nil {
			fmt.Printf("Error saving config: %v\n", err)
		} else {
			fmt.Printf("Theme successfully saved to %s\n", selected)
			fmt.Printf("Rebuilding site with new theme...\n")
			// Auto-rebuild the site to apply theme changes immediately
			buildErr := generator.BuildSite()
			if buildErr != nil {
				fmt.Printf("Warning: Failed to rebuild site: %v\n", buildErr)
				fmt.Printf("Run 'bazel build' manually to apply theme changes.\n")
			} else {
				fmt.Printf("Site rebuilt successfully! Theme changes are now live.\n")
			}
		}
		return m, tea.Quit
	}
	return m, nil
}

func (m model) updateFontMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "esc":
		m.state = ConfigMenu
		m.cursor = 0
		m.previewFont = "" // Clear preview font
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
			m.previewFont = config.Fonts[m.cursor] // Update preview font
		}
	case "down", "j":
		if m.cursor < len(config.Fonts)-1 {
			m.cursor++
			m.previewFont = config.Fonts[m.cursor] // Update preview font
		}
	case "enter":
		selected := config.Fonts[m.cursor]
		m.config.SetFont(selected)
		err := m.config.Save()
		if err != nil {
			m.message = fmt.Sprintf("Error saving config: %v", err)
		} else {
			m.message = fmt.Sprintf("Font set to %s", selected)
		}
		m.state = ConfigMenu
		m.cursor = 0
		m.previewFont = "" // Clear preview font
	}
	return m, nil
}

func (m model) updateEditorMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "esc":
		m.state = ConfigMenu
		m.cursor = 0
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(config.Editors)-1 {
			m.cursor++
		}
	case "enter":
		selected := config.Editors[m.cursor]
		m.config.SetEditor(selected)
		err := m.config.Save()
		if err != nil {
			m.message = fmt.Sprintf("Error saving config: %v", err)
		} else {
			m.message = fmt.Sprintf("Editor set to %s", selected)
		}
		m.state = ConfigMenu
		m.cursor = 0
	}
	return m, nil
}

func (m model) updateSocialMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "esc":
		m.state = ConfigMenu
		m.cursor = 0
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.socialPlatforms)-1 {
			m.cursor++
		}
	case "enter":
		selected := m.socialPlatforms[m.cursor]
		m.editingSocial = selected
		if url, exists := m.config.Socials[selected]; exists {
			m.editingURL = url
		} else {
			m.editingURL = ""
		}
		m.state = SocialEditMenu
	case "d":
		if m.cursor < len(m.socialPlatforms) {
			selected := m.socialPlatforms[m.cursor]
			m.config.RemoveSocial(selected)
			err := m.config.Save()
			if err != nil {
				m.message = fmt.Sprintf("Error saving config: %v", err)
			} else {
				m.message = fmt.Sprintf("Removed %s", selected)
			}
		}
	}
	return m, nil
}

func (m model) updateSocialEditMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.state = SocialMenu
		m.cursor = 0
	case "enter":
		if m.editingURL != "" {
			m.config.SetSocial(m.editingSocial, m.editingURL)
			err := m.config.Save()
			if err != nil {
				m.message = fmt.Sprintf("Error saving config: %v", err)
			} else {
				m.message = fmt.Sprintf("Set %s to %s", m.editingSocial, m.editingURL)
			}
		}
		m.state = SocialMenu
		m.cursor = 0
	case "backspace":
		if len(m.editingURL) > 0 {
			m.editingURL = m.editingURL[:len(m.editingURL)-1]
		}
	default:
		// Add character to URL
		if len(msg.String()) == 1 {
			m.editingURL += msg.String()
		}
	}
	return m, nil
}

func (m model) updateConfigMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	configChoices := []string{"Site Settings", "Set Theme", "Set Font", "Set Editor", "Set Socials", "Back to Main Menu"}
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "esc":
		m.state = MainMenu
		m.cursor = 0
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(configChoices)-1 {
			m.cursor++
		}
	case "enter":
		switch m.cursor {
		case 0: // Site Settings
			m.state = SiteSettingsMenu
			m.cursor = 0
		case 1: // Set Theme
			m.state = ThemeMenu
			m.cursor = 0
			for i, scheme := range config.ColorSchemes {
				if scheme == m.config.Theme.ColorScheme {
					m.cursor = i
					break
				}
			}
			m.previewTheme = config.ColorSchemes[m.cursor]
		case 2: // Set Font
			m.state = FontMenu
			m.cursor = 0
			for i, font := range config.Fonts {
				if font == m.config.Theme.Font {
					m.cursor = i
					break
				}
			}
			m.previewFont = config.Fonts[m.cursor]
		case 3: // Set Editor
			m.state = EditorMenu
			m.cursor = 0
			for i, editor := range config.Editors {
				if editor == m.config.GetEditor() {
					m.cursor = i
					break
				}
			}
		case 4: // Set Socials
			m.state = SocialMenu
			m.cursor = 0
		case 5: // Back to Main Menu
			m.state = MainMenu
			m.cursor = 0
		}
	}
	return m, nil
}

func (m model) updateSiteSettingsMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	siteChoices := []string{"Edit Site Title", "Edit Site Description", "Edit Site Domain", "Back to Configuration"}
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "esc":
		m.state = ConfigMenu
		m.cursor = 0
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(siteChoices)-1 {
			m.cursor++
		}
	case "enter":
		switch m.cursor {
		case 0: // Edit Site Title
			m.editingTitle = m.config.Title
			m.state = TitleEditMenu
		case 1: // Edit Site Description
			m.editingDescription = m.config.Description
			m.state = DescriptionEditMenu
		case 2: // Edit Site Domain
			m.editingDomain = m.config.BaseURL
			m.state = DomainEditMenu
		case 3: // Back to Configuration
			m.state = ConfigMenu
			m.cursor = 0
		}
	}
	return m, nil
}

func (m model) updateTitleEditMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.state = SiteSettingsMenu
		m.cursor = 0
	case "enter":
		if m.editingTitle != "" {
			m.config.Title = m.editingTitle
			err := m.config.Save()
			if err != nil {
				m.message = fmt.Sprintf("Error saving config: %v", err)
			} else {
				m.message = fmt.Sprintf("Site title set to %s", m.editingTitle)
			}
		}
		m.state = SiteSettingsMenu
		m.cursor = 0
	case "backspace":
		if len(m.editingTitle) > 0 {
			m.editingTitle = m.editingTitle[:len(m.editingTitle)-1]
		}
	default:
		// Add character to title
		if len(msg.String()) == 1 {
			m.editingTitle += msg.String()
		}
	}
	return m, nil
}

func (m model) updateDomainEditMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.state = SiteSettingsMenu
		m.cursor = 0
	case "enter":
		if m.editingDomain != "" {
			m.config.BaseURL = m.editingDomain
			err := m.config.Save()
			if err != nil {
				m.message = fmt.Sprintf("Error saving config: %v", err)
			} else {
				m.message = fmt.Sprintf("Site domain set to %s", m.editingDomain)
			}
		}
		m.state = SiteSettingsMenu
		m.cursor = 0
	case "backspace":
		if len(m.editingDomain) > 0 {
			m.editingDomain = m.editingDomain[:len(m.editingDomain)-1]
		}
	default:
		// Add character to domain
		if len(msg.String()) == 1 {
			m.editingDomain += msg.String()
		}
	}
	return m, nil
}

func (m model) updateDescriptionEditMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.state = SiteSettingsMenu
		m.cursor = 0
	case "enter":
		if m.editingDescription != "" {
			m.config.Description = m.editingDescription
			err := m.config.Save()
			if err != nil {
				m.message = fmt.Sprintf("Error saving config: %v", err)
			} else {
				m.message = fmt.Sprintf("Site description updated")
			}
		}
		m.state = SiteSettingsMenu
		m.cursor = 0
	case "backspace":
		if len(m.editingDescription) > 0 {
			m.editingDescription = m.editingDescription[:len(m.editingDescription)-1]
		}
	default:
		// Add character to description
		if len(msg.String()) == 1 {
			m.editingDescription += msg.String()
		}
	}
	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}

	// Clear screen for clean display
	s.WriteString(clearScreen())

	// Apply preview styling to the entire interface
	title := "🏗️  Bazel Site Generator Configuration"
	if m.state == FontMenu && m.previewFont != "" {
		s.WriteString(getFontStyle(m.previewFont, title))
	} else if m.state == ThemeMenu && m.previewTheme != "" {
		s.WriteString(getThemeStyle(m.previewTheme, title))
	} else {
		s.WriteString(title)
	}
	s.WriteString("\n\n")

	if m.message != "" {
		msgText := fmt.Sprintf("✅ %s", m.message)
		if m.state == FontMenu && m.previewFont != "" {
			s.WriteString(getFontStyle(m.previewFont, msgText) + "\n\n")
		} else if m.state == ThemeMenu && m.previewTheme != "" {
			s.WriteString(getThemeStyle(m.previewTheme, msgText) + "\n\n")
		} else {
			s.WriteString(msgText + "\n\n")
		}
	}

	switch m.state {
	case MainMenu:
		s.WriteString("Current Settings:\n")
		s.WriteString(fmt.Sprintf("  Theme: %s\n", m.config.Theme.ColorScheme))
		s.WriteString(fmt.Sprintf("  Font: %s\n", m.config.Theme.Font))
		s.WriteString(fmt.Sprintf("  Socials: %d configured\n\n", len(m.config.Socials)))

		s.WriteString("Choose an option:\n\n")
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
		}

	case ThemeMenu:
		headerText := "Choose a color scheme:"
		if m.previewTheme != "" {
			s.WriteString(getThemeStyle(m.previewTheme, headerText))
		} else {
			s.WriteString(headerText)
		}
		s.WriteString("\n\n")

		for i, scheme := range config.ColorSchemes {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			current := ""
			if scheme == m.config.Theme.ColorScheme {
				current = " (current)"
			}

			// Apply theme styling to the theme options
			if m.previewTheme != "" {
				optionText := fmt.Sprintf("%s %s%s", cursor, scheme, current)
				s.WriteString(getThemeStyle(m.previewTheme, optionText) + "\n")
			} else {
				s.WriteString(fmt.Sprintf("%s %s%s\n", cursor, scheme, current))
			}
		}

		// Add theme preview
		if m.previewTheme != "" {
			s.WriteString("\n🎨 Preview: ")
			s.WriteString(getThemePreview(m.previewTheme))
		}

	case FontMenu:
		if m.previewFont != "" {
			s.WriteString(getFontStyle(m.previewFont, "Choose a font:"))
		} else {
			s.WriteString("Choose a font:")
		}
		s.WriteString("\n\n")

		for i, font := range config.Fonts {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			current := ""
			if font == m.config.Theme.Font {
				current = " (current)"
			}

			// Apply font styling to the font options
			if m.previewFont != "" {
				s.WriteString(fmt.Sprintf("%s %s%s\n", cursor, getFontStyle(m.previewFont, font), current))
			} else {
				s.WriteString(fmt.Sprintf("%s %s%s\n", cursor, font, current))
			}
		}

		// Add font preview text
		if m.previewFont != "" {
			s.WriteString("\n📝 Preview: ")
			s.WriteString(getFontPreview(m.previewFont))
		}

	case EditorMenu:
		s.WriteString("Choose an editor:\n\n")
		for i, editor := range config.Editors {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			current := ""
			if editor == m.config.GetEditor() {
				current = " (current)"
			}

			// Add description for editors
			description := ""
			switch editor {
			case "auto":
				description = " - Auto-detect from $EDITOR or fallback to vi"
			case "vim":
				description = " - Vi IMproved"
			case "nvim":
				description = " - Neovim"
			case "nano":
				description = " - Simple terminal editor"
			case "emacs":
				description = " - GNU Emacs"
			case "hx":
				description = " - Helix editor"
			case "vi":
				description = " - Classic vi editor"
			case "code":
				description = " - Visual Studio Code"
			case "subl":
				description = " - Sublime Text"
			case "atom":
				description = " - GitHub Atom"
			case "gedit":
				description = " - GNOME Text Editor"
			case "kate":
				description = " - KDE Advanced Text Editor"
			}

			s.WriteString(fmt.Sprintf("%s %s%s%s\n", cursor, editor, current, description))
		}

	case SocialMenu:
		s.WriteString("Configure social profiles:\n")
		s.WriteString("(Press Enter to edit, 'd' to delete)\n\n")
		for i, platform := range m.socialPlatforms {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			url := ""
			if configuredURL, exists := m.config.Socials[platform]; exists {
				url = fmt.Sprintf(" -> %s", configuredURL)
			}
			s.WriteString(fmt.Sprintf("%s %s%s\n", cursor, platform, url))
		}

	case SocialEditMenu:
		s.WriteString(fmt.Sprintf("Edit %s URL:\n\n", m.editingSocial))
		s.WriteString(fmt.Sprintf("URL: %s|\n", m.editingURL))
		s.WriteString("\n(Press Enter to save, Esc to cancel, Ctrl+C to quit)")

	case ConfigMenu:
		configChoices := []string{"Site Settings", "Set Theme", "Set Font", "Set Editor", "Set Socials", "Back to Main Menu"}
		s.WriteString("Configuration Menu:\n\n")
		for i, choice := range configChoices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
		}

	case SiteSettingsMenu:
		siteChoices := []string{"Edit Site Title", "Edit Site Description", "Edit Site Domain", "Back to Configuration"}
		s.WriteString("Site Settings:\n")
		s.WriteString(fmt.Sprintf("  Current Title: %s\n", m.config.Title))
		s.WriteString(fmt.Sprintf("  Current Description: %s\n", m.config.Description))
		s.WriteString(fmt.Sprintf("  Current Domain: %s\n\n", m.config.BaseURL))
		for i, choice := range siteChoices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
		}

	case TitleEditMenu:
		s.WriteString("Edit Site Title:\n\n")
		s.WriteString(fmt.Sprintf("Title: %s|\n", m.editingTitle))
		s.WriteString("\n(Press Enter to save, Esc to cancel, Ctrl+C to quit)")

	case DescriptionEditMenu:
		s.WriteString("Edit Site Description:\n\n")
		s.WriteString(fmt.Sprintf("Description: %s|\n", m.editingDescription))
		s.WriteString("\n(Press Enter to save, Esc to cancel, Ctrl+C to quit)")

	case DomainEditMenu:
		s.WriteString("Edit Site Domain:\n\n")
		s.WriteString(fmt.Sprintf("Domain: %s|\n", m.editingDomain))
		s.WriteString("\n(Press Enter to save, Esc to cancel, Ctrl+C to quit)")

	case PostTitleInputMenu:
		s.WriteString("Create New Post\n\n")
		s.WriteString(fmt.Sprintf("Post Title: %s|\n", m.editingTitle))
		s.WriteString("\n(Press Enter to create, Esc to cancel, Ctrl+C to quit)")

	case PageTitleInputMenu:
		s.WriteString("Create New Page\n\n")
		s.WriteString(fmt.Sprintf("Page Title: %s|\n", m.editingTitle))
		s.WriteString("\n(Press Enter to create, Esc to cancel, Ctrl+C to quit)")

	case PostEditMenu:
		s.WriteString("Select a post to edit:\n\n")
		for i, post := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, post))
		}
		s.WriteString("\n(Press Enter to edit, Esc to go back, Ctrl+C to quit)")

	case PageEditMenu:
		s.WriteString("Select a page to edit:\n\n")
		for i, page := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, page))
		}
		s.WriteString("\n(Press Enter to edit, Esc to go back, Ctrl+C to quit)")
	}

	// Apply preview styling to the quit instruction
	quitText := "Press 'q' to quit"
	if m.state == FontMenu && m.previewFont != "" {
		s.WriteString("\n\n")
		s.WriteString(getFontStyle(m.previewFont, quitText))
	} else if m.state == ThemeMenu && m.previewTheme != "" {
		s.WriteString("\n\n")
		s.WriteString(getThemeStyle(m.previewTheme, quitText))
	} else {
		s.WriteString("\n\n" + quitText)
	}

	return s.String()
}

// updatePostTitleInputMenu handles post title input
func (m model) updatePostTitleInputMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.state = MainMenu
		m.cursor = 0
	case "enter":
		if m.editingTitle != "" {
			err := generator.NewPost(m.editingTitle)
			if err != nil {
				m.message = fmt.Sprintf("Error creating post: %v", err)
			} else {
				m.message = fmt.Sprintf("Created post: %s", m.editingTitle)
			}
		}
		m.state = MainMenu
		m.cursor = 0
	case "backspace":
		if len(m.editingTitle) > 0 {
			m.editingTitle = m.editingTitle[:len(m.editingTitle)-1]
		}
	default:
		// Add character to title
		if len(msg.String()) == 1 {
			m.editingTitle += msg.String()
		}
	}
	return m, nil
}

// updatePageTitleInputMenu handles page title input
func (m model) updatePageTitleInputMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.state = MainMenu
		m.cursor = 0
	case "enter":
		if m.editingTitle != "" {
			err := generator.NewPage(m.editingTitle)
			if err != nil {
				m.message = fmt.Sprintf("Error creating page: %v", err)
			} else {
				m.message = fmt.Sprintf("Created page: %s", m.editingTitle)
			}
		}
		m.state = MainMenu
		m.cursor = 0
	case "backspace":
		if len(m.editingTitle) > 0 {
			m.editingTitle = m.editingTitle[:len(m.editingTitle)-1]
		}
	default:
		// Add character to title
		if len(msg.String()) == 1 {
			m.editingTitle += msg.String()
		}
	}
	return m, nil
}

// updatePostEditMenu handles post selection for editing
func (m model) updatePostEditMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		// Go back to main post menu
		m.state = MainMenu
		m.choices = []string{"New Post", "Edit Post", "Draft Posts", "Done"}
		m.cursor = 0
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		}
	case "enter":
		if m.cursor < len(m.choices) {
			selectedPost := m.choices[m.cursor]
			err := generator.EditPost(selectedPost)
			if err != nil {
				m.message = fmt.Sprintf("Error opening post for editing: %v", err)
			} else {
				m.message = fmt.Sprintf("Opened %s in editor", selectedPost)
			}
			// Return to main post menu after editing
			m.state = MainMenu
			m.choices = []string{"New Post", "Edit Post", "Draft Posts", "Done"}
			m.cursor = 0
		}
	}
	return m, nil
}

// updatePageEditMenu handles page selection for editing
func (m model) updatePageEditMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		// Go back to main page menu
		m.state = MainMenu
		m.choices = []string{"New Page", "Edit Page", "Draft Pages", "Organize Pages", "Done"}
		m.cursor = 0
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		}
	case "enter":
		if m.cursor < len(m.choices) {
			selectedPage := m.choices[m.cursor]
			err := generator.EditPage(selectedPage)
			if err != nil {
				m.message = fmt.Sprintf("Error opening page for editing: %v", err)
			} else {
				m.message = fmt.Sprintf("Opened %s in editor", selectedPage)
			}
			// Return to main page menu after editing
			m.state = MainMenu
			m.choices = []string{"New Page", "Edit Page", "Draft Pages", "Organize Pages", "Done"}
			m.cursor = 0
		}
	}
	return m, nil
}
