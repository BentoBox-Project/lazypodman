package podman

import (
	"context"
	"os"

	"github.com/containers/libpod/v2/pkg/bindings"
	"github.com/containers/libpod/v2/pkg/bindings/containers"
	"github.com/containers/libpod/v2/pkg/bindings/images"
	"github.com/containers/libpod/v2/pkg/bindings/pods"
	"github.com/containers/libpod/v2/pkg/bindings/volumes"
)

// Podman struct
type Podman struct {
	// the name of the docker-compose file, if any
	ComposeFile string
}

// APIConn returns an Podman V2 API connection as a context.Context
func APIConn() (context.Context, error) {
	// Get Podman socket location
	sockDir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sockDir + "/podman/podman.sock"

	// Connect to Podman socket
	connText, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		return nil, err
	}

	return connText, nil
}

// Pods returns a slice of strings with the name of the active pods
func (p *Podman) Pods(conn context.Context) ([]string, error) {
	podList, err := pods.List(conn, nil)
	if err != nil {
		return nil, err
	}

	var podNames []string

	for _, pod := range podList {
		podNames = append(podNames, pod.Name)
	}

	return podNames, nil
}

// Containers retuns a slice of strings with the names of the active containers or those listted on a docker-compose file
func (p *Podman) Containers(conn context.Context) ([]string, error) {
	var latestContainers = 10

	containerList, err := containers.List(conn, nil, nil, &latestContainers, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var containerNames []string
	for _, c := range containerList {
		containerNames = append(containerNames, c.Names[0])
	}

	return containerNames, nil
}

// Images return the list of the current podman images in the system
func (p *Podman) Images(conn context.Context) ([]string, error) {
	// List images
	imageSummary, err := images.List(conn, nil, nil)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, i := range imageSummary {
		names = append(names, i.RepoTags...)
	}

	return names, nil
}

// Volumes return the list of the current volumnes in the system
func (p *Podman) Volumes(conn context.Context) ([]string, error) {
	volumeList, err := volumes.List(conn, nil)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, v := range volumeList {
		names = append(names, v.Name)
	}

	return names, nil
}
