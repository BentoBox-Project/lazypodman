package podman

import "context"

// Binder declare the methods used to get the list of the active resources
type Binder interface {
	Pods(ctx context.Context) ([]string, error)
	Images(ctx context.Context) ([]string, error)
	Containers(ctx context.Context) ([]string, error)
	Volumes(ctx context.Context) ([]string, error)
}
