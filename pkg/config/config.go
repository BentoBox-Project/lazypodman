package config

// Config struct
type Config struct {
	Debug      bool
	Name       string
	ProjectDir string
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
		Version: version,
	}

	return appConfig, nil
}
