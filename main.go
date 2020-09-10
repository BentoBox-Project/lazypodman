package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danvergara/lazypodman/pkg/app"
	"github.com/danvergara/lazypodman/pkg/config"
)

const (
	appVersion   = "0.1.0"
	versionUsage = "Prints current version"
	fileUsage    = "Specify a alternate compose file"
)

var (
	version     bool
	composeFile string
)

func init() {
	flag.BoolVar(&version, "version", false, versionUsage)
	flag.BoolVar(&version, "v", false, versionUsage+" (shorthand)")
	flag.StringVar(&composeFile, "file", "", fileUsage)
	flag.StringVar(&composeFile, "f", "", fileUsage+" (shorthand)")
}

func main() {

	flag.Parse()

	if version {
		fmt.Println(appVersion)
		os.Exit(0)
	}

	projectDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	appConfig, err := config.NewConfig("lazypodman", composeFile, projectDir)

	if app, err := app.NewApp(appConfig); err == nil {
		err = app.Run()
	}

	if err != nil {
		log.Fatal(fmt.Sprintf("%s\n", "something happend"))
	}
}
