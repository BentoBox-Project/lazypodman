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
	// flags
	name string

	podCmd = &cobra.Command{
		Use:   "pod",
		Short: "Handle and display all the information related to a given pod",
		Long:  "Handler and display all the information related to a given pod: $ lzd pod -n awesome_pod",
		RunE: func(cmd *cobra.Command, args []string) error {
			projectDir, err := os.Getwd()
			if err != nil {
				log.Fatal(err.Error())
			}

			appConfig, err := config.NewConfig("lazypodman", projectDir, "0.1.0")

			if err != nil {
				log.Fatal(fmt.Sprintf("%s\n", err.Error()))
			}

			app, err := app.NewApp(appConfig)
			app.Gui.PodmanBinding.Pod = name
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
)

func init() {
	podCmd.Flags().StringVarP(&name, "name", "n", "", "the name of the pod of interest")
}
