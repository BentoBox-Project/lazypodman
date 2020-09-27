package config

// Config struct
type Config struct {
	Debug      bool
	ProjectDir string
	Name       string
	Version    string
}

// NewConfig boostrap a new config for the application
func NewConfig(name, projectDir, version string) (*Config, error) {
	appConfig := &Config{
		Name:       name,
		ProjectDir: projectDir,
		Version:    version,
	}

	return appConfig, nil
}
