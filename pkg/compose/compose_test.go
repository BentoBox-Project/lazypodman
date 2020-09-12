package compose

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposeFileExist(t *testing.T) {
	assert.True(t, FileExists("docker-compose.yml"))
}

func TestComposeFileNotExist(t *testing.T) {
	assert.False(t, FileExists("docker-compose-dev.yml"))
}

func TestComposeParsing(t *testing.T) {
	composeObject, _ := parseComposeFile("docker-compose.yml")
	assert.Equal(t, 2, len(composeObject.Services))
	assert.Equal(t, "wordpress:latest", composeObject.Services["web"].Image)
	assert.Equal(t, "mariadb:latest", composeObject.Services["db"].Image)
}

func TestComposeServices(t *testing.T) {
	services, _ := Services("docker-compose.yml")
	assert.Equal(t, 2, len(services))
}
