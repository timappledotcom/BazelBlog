package ui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/yourusername/bazel_blog/internal/registry"
)

type SiteSelector struct {
	sites        []registry.Site
	cursor       int
	selectedSite *registry.Site
	currentDir   string
	message      string
}

func NewSiteSelector() (*SiteSelector, error) {
	reg, err := registry.LoadRegistry()
	if err != nil {
		return nil, fmt.Errorf("failed to load site registry: %w", err)
	}

	// Clean up invalid sites
	reg.ValidateSites()
	reg.Save()

	currentDir, _ := os.Getwd()

	return &SiteSelector{
		sites:      reg.GetSites(),
		cursor:     0,
		currentDir: currentDir,
	}, nil
}

func (s SiteSelector) Init() tea.Cmd {
	return nil
}

func (s SiteSelector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return s, tea.Quit
		case "esc":
			return s, tea.Quit
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
			}
		case "down", "j":
			if s.cursor < len(s.sites)-1 {
				s.cursor++
			}
		case "enter":
			if len(s.sites) > 0 && s.cursor < len(s.sites) {
				s.selectedSite = &s.sites[s.cursor]
				return s, tea.Quit
			}
		case "d":
			// Delete site from registry
			if len(s.sites) > 0 && s.cursor < len(s.sites) {
				siteToDelete := s.sites[s.cursor]
				reg, err := registry.LoadRegistry()
				if err == nil {
					reg.RemoveSite(siteToDelete.Path)
					reg.Save()
					// Reload sites
					s.sites = reg.GetSites()
					if s.cursor >= len(s.sites) && len(s.sites) > 0 {
						s.cursor = len(s.sites) - 1
					}
					s.message = fmt.Sprintf("Removed '%s' from registry", siteToDelete.Name)
				}
			}
		case "r":
			// Refresh sites list
			reg, err := registry.LoadRegistry()
			if err == nil {
				reg.ValidateSites()
				reg.Save()
				s.sites = reg.GetSites()
				if s.cursor >= len(s.sites) && len(s.sites) > 0 {
					s.cursor = len(s.sites) - 1
				}
				s.message = "Refreshed sites list"
			}
		}
	}
	return s, nil
}

func (s SiteSelector) View() string {
	var b strings.Builder

	b.WriteString("🏗️  Bazel Site Selector\n\n")

	if s.message != "" {
		b.WriteString(fmt.Sprintf("✅ %s\n\n", s.message))
	}

	if len(s.sites) == 0 {
		b.WriteString("No Bazel sites found in registry.\n")
		b.WriteString("Create a new site with: bazel new site <name>\n\n")
		b.WriteString("Press 'q' to quit")
		return b.String()
	}

	b.WriteString("Select a site to work with:\n\n")

	for i, site := range s.sites {
		cursor := " "
		if s.cursor == i {
			cursor = ">"
		}

		// Format the last used time
		lastUsed := "never"
		if !site.LastUsed.IsZero() {
			if time.Since(site.LastUsed) < 24*time.Hour {
				lastUsed = fmt.Sprintf("%d hours ago", int(time.Since(site.LastUsed).Hours()))
			} else {
				lastUsed = site.LastUsed.Format("Jan 2, 2006")
			}
		}

		// Check if site is the current directory
		currentIndicator := ""
		if site.Path == s.currentDir {
			currentIndicator = " (current directory)"
		}

		b.WriteString(fmt.Sprintf("%s %s%s\n", cursor, site.Name, currentIndicator))
		b.WriteString(fmt.Sprintf("   📁 %s\n", site.Path))
		if site.Description != "" {
			b.WriteString(fmt.Sprintf("   📝 %s\n", site.Description))
		}
		b.WriteString(fmt.Sprintf("   🕒 Last used: %s\n", lastUsed))
		b.WriteString("\n")
	}

	b.WriteString("Navigation:\n")
	b.WriteString("• ↑/↓ or k/j: Move cursor\n")
	b.WriteString("• Enter: Select site\n")
	b.WriteString("• d: Remove site from registry\n")
	b.WriteString("• r: Refresh sites list\n")
	b.WriteString("• q: Quit\n")

	return b.String()
}

func (s *SiteSelector) GetSelectedSite() *registry.Site {
	return s.selectedSite
}

func RunSiteSelector() (*registry.Site, error) {
	selector, err := NewSiteSelector()
	if err != nil {
		return nil, err
	}

	p := tea.NewProgram(selector)
	finalModel, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("error running site selector: %w", err)
	}

	finalSelector := finalModel.(SiteSelector)
	return finalSelector.GetSelectedSite(), nil
}
