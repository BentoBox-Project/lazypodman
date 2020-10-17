package podman

import (
	"context"
	"os"

	"github.com/containers/podman/v2/pkg/bindings"
)

// Podman struct
type Podman struct {
	// the name of the docker-compose file, if any
	ComposeFile string
	// the pod of interest to analyze
	Pod string
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
func (p *Podman) Pods(ctx context.Context, b Binder) ([]string, error) {
	podList, err := b.Pods(ctx, nil)
	if err != nil {
		return nil, err
	}

	var podNames []string

	for _, pod := range podList {
		podNames = append(podNames, pod.Name)
	}

	if p.Pod != "" {
		for _, pod := range podNames {
			if p.Pod == pod {
				return []string{p.Pod}, nil
			}
		}
	}

	return podNames, nil
}

// Containers retuns a slice of strings with the names of the active containers or those listted on a docker-compose file
func (p *Podman) Containers(ctx context.Context, b Binder) ([]string, error) {
	var latestContainers = 10
	var containerNames []string

	if p.Pod != "" {
		pod, err := b.Pod(ctx, p.Pod)
		if err != nil {
			return nil, err
		}

		for _, c := range pod.Containers {
			containerNames = append(containerNames, c.Name)
		}

		return containerNames, nil
	}

	containerList, err := b.Containers(ctx, nil, nil, &latestContainers, nil, nil)
	if err != nil {
		return nil, err
	}

	for _, c := range containerList {
		containerNames = append(containerNames, c.Names[0])
	}

	return containerNames, nil
}

// Images return the list of the current podman images in the system
func (p *Podman) Images(ctx context.Context, b Binder) ([]string, error) {
	// List images
	imageSummary, err := b.Images(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	var images []string
	for _, i := range imageSummary {
		images = append(images, i.RepoTags...)
	}

	return images, nil
}

// Volumes return the list of the current volumnes in the system
func (p *Podman) Volumes(ctx context.Context, b Binder) ([]string, error) {
	volumeList, err := b.Volumes(ctx, nil)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, v := range volumeList {
		names = append(names, v.Name)
	}

	return names, nil
}
