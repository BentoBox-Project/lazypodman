package config

// Config struct
type Config struct {
	Name       string
	ProjectDir string
	UserConfig *UserConfig
}

// UserConfig struct
type UserConfig struct {
	ComposeFile string
}

// NewConfig boostrap a new config for the application
func NewConfig(name, composeFile, projectDir string) (*Config, error) {
	appConfig := &Config{
		Name:       name,
		ProjectDir: projectDir,
		UserConfig: &UserConfig{
			ComposeFile: composeFile,
		},
	}

	return appConfig, nil
}
