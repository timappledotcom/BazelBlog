package registry

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Site struct {
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	CreatedAt   time.Time `json:"created_at"`
	LastUsed    time.Time `json:"last_used"`
	Description string    `json:"description"`
}

type Registry struct {
	Sites []Site `json:"sites"`
}

func getRegistryPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "bazel")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create config directory: %w", err)
	}

	return filepath.Join(configDir, "sites.json"), nil
}

func LoadRegistry() (*Registry, error) {
	registryPath, err := getRegistryPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(registryPath); os.IsNotExist(err) {
		// Create empty registry if it doesn't exist
		return &Registry{Sites: []Site{}}, nil
	}

	data, err := os.ReadFile(registryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read registry file: %w", err)
	}

	var registry Registry
	if err := json.Unmarshal(data, &registry); err != nil {
		return nil, fmt.Errorf("failed to parse registry file: %w", err)
	}

	return &registry, nil
}

func (r *Registry) Save() error {
	registryPath, err := getRegistryPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal registry: %w", err)
	}

	if err := os.WriteFile(registryPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write registry file: %w", err)
	}

	return nil
}

func (r *Registry) AddSite(name, path, description string) error {
	// Check if site already exists
	for i, site := range r.Sites {
		if site.Path == path {
			// Update existing site
			r.Sites[i].Name = name
			r.Sites[i].Description = description
			r.Sites[i].LastUsed = time.Now()
			return r.Save()
		}
	}

	// Add new site
	site := Site{
		Name:        name,
		Path:        path,
		CreatedAt:   time.Now(),
		LastUsed:    time.Now(),
		Description: description,
	}

	r.Sites = append(r.Sites, site)
	return r.Save()
}

func (r *Registry) RemoveSite(path string) error {
	for i, site := range r.Sites {
		if site.Path == path {
			r.Sites = append(r.Sites[:i], r.Sites[i+1:]...)
			return r.Save()
		}
	}
	return fmt.Errorf("site not found in registry")
}

func (r *Registry) UpdateLastUsed(path string) error {
	for i, site := range r.Sites {
		if site.Path == path {
			r.Sites[i].LastUsed = time.Now()
			return r.Save()
		}
	}
	return nil // Don't error if site not found, just ignore
}

func (r *Registry) GetSites() []Site {
	return r.Sites
}

func (r *Registry) FindSiteByName(name string) (*Site, error) {
	for _, site := range r.Sites {
		if site.Name == name {
			return &site, nil
		}
	}
	return nil, fmt.Errorf("site with name '%s' not found", name)
}

func (r *Registry) ValidateSites() {
	// Remove sites that no longer exist on disk
	validSites := []Site{}
	for _, site := range r.Sites {
		configPath := filepath.Join(site.Path, "bazel.toml")
		if _, err := os.Stat(configPath); err == nil {
			validSites = append(validSites, site)
		}
	}
	r.Sites = validSites
}
