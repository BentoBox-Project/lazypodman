package app

import (
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
	if err := app.Config.Gui.Run(); err != nil {
		return err
	}

	return nil
}
