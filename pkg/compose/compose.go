package compose

import (
	"os"

	"gopkg.in/yaml.v2"
)

// File struct holds the docker-compose file main information
type File struct {
	Services map[string]Service `yaml:"services"`
}

// Service struct store the main information from each service in the document
type Service struct {
	Image string `yaml:"image"`
}

// FileExists checks if a docker compose file exists
func FileExists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func parseComposeFile(name string) (*File, error) {
	// Create file structure
	composeFile := &File{}

	// Open compose file
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&composeFile); err != nil {
		return nil, err
	}

	return composeFile, nil
}

// Services retuns a slice of the listes services on the docker-compose file
func Services(name string) ([]string, error) {
	composeFile, err := parseComposeFile(name)
	if err != nil {
		return nil, err
	}

	services := make([]string, 0, len(composeFile.Services))

	for k := range composeFile.Services {
		services = append(services, k)
	}

	return services, nil
}
