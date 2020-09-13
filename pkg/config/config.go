package config

import "github.com/danvergara/lazypodman/pkg/gui"

// Config struct
type Config struct {
	Debug      bool
	Gui        *gui.Gui
	ProjectDir string
	Name       string
	UserConfig *UserConfig
	Version    string
}

// UserConfig struct
type UserConfig struct {
	ComposeFile string
}

// NewConfig boostrap a new config for the application
func NewConfig(name, composeFile, projectDir, version string) (*Config, error) {
	appConfig := &Config{
		Name:       name,
		ProjectDir: projectDir,
		UserConfig: &UserConfig{
			ComposeFile: composeFile,
		},
		Gui:     gui.NewGui(),
		Version: version,
	}

	return appConfig, nil
}
