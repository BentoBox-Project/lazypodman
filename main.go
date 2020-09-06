package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/libpod/pkg/bindings"
	"github.com/containers/libpod/pkg/bindings/containers"
)

func main() {
	fmt.Println("Welcome to the Podman Go bindings tutorial")

	// Get Podman socket location
	sockDir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sockDir + "/podman/podman.sock"

	// Connect to Podman socket
	connText, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Container list
	var latestContainers = 3
	containerLatestList, err := containers.List(connText, nil, nil, &latestContainers, nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, container := range containerLatestList {
		fmt.Printf("container: %s\n", container.Names[0])
	}

	// fmt.Printf("Latest container is %s\n", containerLatestList[0].Names[0])
}
