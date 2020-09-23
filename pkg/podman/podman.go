package podman

import (
	"context"
	"os"

	"github.com/containers/libpod/v2/pkg/bindings"
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
	ctx, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}

// Pods returns a slice of strings with the name of the active pods
func (p *Podman) Pods(ctx context.Context, pods Pods) ([]string, error) {
	podList, err := pods(ctx, nil)
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
func (p *Podman) Containers(ctx context.Context, crs Containers) ([]string, error) {
	var latestContainers = 10

	containerList, err := crs(ctx, nil, nil, &latestContainers, nil, nil, nil)
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
func (p *Podman) Images(ctx context.Context, imgs Images) ([]string, error) {
	// List images
	imageSummary, err := imgs(ctx, nil, nil)
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
func (p *Podman) Volumes(ctx context.Context, v Volumes) ([]string, error) {
	volumeList, err := v(ctx, nil)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, v := range volumeList {
		names = append(names, v.Name)
	}

	return names, nil
}
