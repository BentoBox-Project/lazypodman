package podman

import (
	"context"

	"github.com/containers/libpod/v2/pkg/domain/entities"
)

// Pods is the function type with the same signature of pods.List
type Pods func(ctx context.Context, filters map[string][]string) ([]*entities.ListPodsReport, error)

// Containers is the function type with the same signature of containers.List
type Containers func(ctx context.Context, filters map[string][]string, all *bool, last *int, pod, size, sync *bool) ([]entities.ListContainer, error)

// Images is the function type with the same signature of images.List
type Images func(ctx context.Context, all *bool, filters map[string][]string) ([]*entities.ImageSummary, error)

// Volumes is the function type with the same signature of volumes.List
type Volumes func(ctx context.Context, filters map[string][]string) ([]*entities.VolumeListReport, error)
