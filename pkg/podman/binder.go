package podman

import (
	"context"

	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/containers/podman/v2/pkg/bindings/pods"
	"github.com/containers/podman/v2/pkg/bindings/volumes"
	"github.com/containers/podman/v2/pkg/domain/entities"
)

// Binder interface define all the required functions from the Podman's bindings
type Binder interface {
	Pods(context.Context, map[string][]string) ([]*entities.ListPodsReport, error)
	Containers(context.Context, map[string][]string, *bool, *int, *bool, *bool) ([]entities.ListContainer, error)
	Images(context.Context, *bool, map[string][]string) ([]*entities.ImageSummary, error)
	Volumes(context.Context, map[string][]string) ([]*entities.VolumeListReport, error)
	Pod(context.Context, string) (*entities.PodInspectReport, error)
}

// Bindings implements the Binder interface
// intended to make the package easier to test
type Bindings struct{}

// Pods returns a list of ListPodsReports and error if any
// Calls the List functions from the pods pkg
func (b *Bindings) Pods(ctx context.Context, filters map[string][]string) ([]*entities.ListPodsReport, error) {
	return pods.List(ctx, filters)
}

// Containers returns a slice of ListContainer and error if any
// Calls the List function from containers pkg
func (b *Bindings) Containers(ctx context.Context, filters map[string][]string, all *bool, last *int, size, sync *bool) ([]entities.ListContainer, error) {
	return containers.List(ctx, filters, all, last, size, sync)
}

// Images returns a slice of ImageSummary and error if any
// Calls the List function from the images pkg
func (b *Bindings) Images(ctx context.Context, all *bool, filters map[string][]string) ([]*entities.ImageSummary, error) {
	return images.List(ctx, all, filters)
}

// Volumes returns a slice of Volum and error if any
// Calls the List function from the volumes pkg
func (b *Bindings) Volumes(ctx context.Context, filters map[string][]string) ([]*entities.VolumeListReport, error) {
	return volumes.List(ctx, filters)
}

// Pod returns a PodInspectReport object and error if any
// Calls the Inspect function from pods pkg
func (b *Bindings) Pod(ctx context.Context, nameOrID string) (*entities.PodInspectReport, error) {
	return pods.Inspect(ctx, nameOrID)
}
