package app

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/libpod/v2/pkg/bindings"
	"github.com/containers/libpod/v2/pkg/bindings/containers"
	"github.com/containers/libpod/v2/pkg/bindings/pods"
	"github.com/danvergara/lazypodman/pkg/config"
)

// App struct
type App struct {
	Config *config.Config
}

// NewApp boostrap a new application
func NewApp(config *config.Config) (*App, error) {
	return &App{Config: config}, nil
}

// Run the application
func (app *App) Run() error {
	// Get Podman socket location
	sockDir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sockDir + "/podman/podman.sock"

	// Connect to Podman socket
	connText, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	podsList, err := pods.List(connText, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, pod := range podsList {
		fmt.Printf("Pod ID: %s Name: %s\n", pod.Id, pod.Name)
		fmt.Printf("Containers of %s: \n", pod.Name)
		for _, ctr := range pod.Containers {
			ctrData, err := containers.Inspect(connText, ctr.Id, nil)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("Specs: %v\n", ctrData)
		}
	}

	return nil
}
