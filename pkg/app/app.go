package app

import (
	"fmt"
	"os"

	"github.com/danvergara/lazypodman/pkg/compose"
	"github.com/danvergara/lazypodman/pkg/config"
	"github.com/danvergara/lazypodman/pkg/logger"
	"github.com/sirupsen/logrus"
)

// App struct
type App struct {
	Config *config.Config
	Log    *logrus.Entry
}

// NewApp boostrap a new application
func NewApp(config *config.Config) (*App, error) {
	app := &App{Config: config}
	app.Log = logger.NewLogger(config)

	return app, nil

}

// Run the application
func (app *App) Run() error {
	// Get Podman socket location

	if compose.FileExists(app.Config.UserConfig.ComposeFile) {
		services, err := compose.Services(app.Config.UserConfig.ComposeFile)
		if err != nil {
			app.Log.Error(err)
			os.Exit(1)
		}
		fmt.Printf("%v\n", services)
	}
	return nil
}
