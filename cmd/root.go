package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/danvergara/lazypodman/pkg/app"
	"github.com/danvergara/lazypodman/pkg/config"
	"github.com/go-errors/errors"
	"github.com/spf13/cobra"
)

var (
	composeFile string
)

var rootCmd = &cobra.Command{
	Use:   "lazypodman",
	Short: "Lazypodman is a monitoring tool for podman pods",
	Long:  "A flexible monitoring tool for pods powered by the podman V2 bindings",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err.Error())
		}

		if composeFile == "" {
			composeFile = "docker-compose.yml"
		}

		appConfig, err := config.NewConfig("lazypodman", projectDir, "0.1.0")

		if err != nil {
			log.Fatal(fmt.Sprintf("%s\n", err.Error()))
		}

		app, err := app.NewApp(appConfig, composeFile)
		if err == nil {
			err = app.Run()
		}

		if err != nil {
			newErr := errors.Wrap(err, 0)
			stackTrace := newErr.ErrorStack()
			app.Log.Error(stackTrace)
			log.Fatal(fmt.Sprintf("%s\n", err.Error()))
		}

		return nil
	},
}

// Execute main function of rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
