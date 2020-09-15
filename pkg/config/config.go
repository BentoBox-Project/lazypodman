package config

import "github.com/danvergara/lazypodman/pkg/gui"

// Config struct
type Config struct {
	Debug      bool
	Gui        *gui.Gui
	ProjectDir string
	Name       string
	Version    string
}

// NewConfig boostrap a new config for the application
func NewConfig(name, composeFile, projectDir, version string) (*Config, error) {
	appConfig := &Config{
		Name:       name,
		ProjectDir: projectDir,
		Gui:        gui.NewGui(composeFile),
		Version:    version,
	}

	return appConfig, nil
}
