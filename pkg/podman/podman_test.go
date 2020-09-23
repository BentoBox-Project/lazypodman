package podman

import (
	"context"
	"testing"

	"github.com/containers/libpod/v2/pkg/domain/entities"
	"github.com/stretchr/testify/assert"
)

func mockPodList(ctx context.Context, filters map[string][]string) ([]*entities.ListPodsReport, error) {
	reports := []*entities.ListPodsReport{
		{
			Name: "application_web",
		},
		{
			Name: "api_startup",
		},
	}

	return reports, nil
}

func mockEmptyPodList(ctx context.Context, filter map[string][]string) ([]*entities.ListPodsReport, error) {
	reports := make([]*entities.ListPodsReport, 0)
	return reports, nil
}

func mockContainersList(ctx context.Context, filters map[string][]string, all *bool, last *int, pod, size, sync *bool) ([]entities.ListContainer, error) {
	containers := []entities.ListContainer{
		{
			Names: []string{"web"},
		},
		{
			Names: []string{"db"},
		},
	}
	return containers, nil
}

func mockEmptyContainersList(ctx context.Context, filters map[string][]string, all *bool, last *int, pod, size, sync *bool) ([]entities.ListContainer, error) {
	containers := make([]entities.ListContainer, 0)
	return containers, nil
}

func mockImagesList(ctx context.Context, all *bool, filters map[string][]string) ([]*entities.ImageSummary, error) {
	images := []*entities.ImageSummary{
		{
			RepoTags: []string{"docker.io/library/mariabdb:latest"},
		},
		{
			RepoTags: []string{"docker.io/library/wordpress:latest"},
		},
		{
			RepoTags: []string{"docker.io/library/busybox:latest"},
		},
	}
	return images, nil
}

func mockEmptyImagesList(ctx context.Context, all *bool, filters map[string][]string) ([]*entities.ImageSummary, error) {
	images := make([]*entities.ImageSummary, 0)
	return images, nil
}

func mockVolumeList(ctx context.Context, filters map[string][]string) ([]*entities.VolumeListReport, error) {
	volumes := []*entities.VolumeListReport{
		{
			VolumeConfigResponse: entities.VolumeConfigResponse{
				Name: "2398675946hj5gh435574gerrw36hg59",
			},
		},
		{
			VolumeConfigResponse: entities.VolumeConfigResponse{
				Name: "746592436592ret336592hg63v54k3659",
			},
		},
	}

	return volumes, nil
}

func mockEmptyVolumelist(ctx context.Context, filters map[string][]string) ([]*entities.VolumeListReport, error) {
	volumes := make([]*entities.VolumeListReport, 0)

	return volumes, nil
}

func TestPodList(t *testing.T) {
	ctx := context.Background()

	podmanObj := Podman{}

	pods, err := podmanObj.Pods(ctx, mockPodList)
	assert.NoError(t, err)
	assert.Equal(t, "application_web", pods[0])
}

func TestEmptyPodList(t *testing.T) {
	ctx := context.Background()

	podmanObj := Podman{}

	pods, err := podmanObj.Pods(ctx, mockEmptyPodList)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(pods))
}

func TestContainerList(t *testing.T) {
	ctx := context.Background()

	podmanObj := Podman{}

	containers, err := podmanObj.Containers(ctx, mockContainersList)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(containers))
	assert.Equal(t, "web", containers[0])
}

func TestEmptyContainersList(t *testing.T) {
	ctx := context.Background()

	podmanObj := Podman{}

	containers, err := podmanObj.Containers(ctx, mockEmptyContainersList)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(containers))
}

func TestImageList(t *testing.T) {
	ctx := context.Background()

	podmanObj := Podman{}

	images, err := podmanObj.Images(ctx, mockImagesList)

	assert.NoError(t, err)
	assert.Equal(t, 3, len(images))
	assert.Equal(t, "docker.io/library/mariabdb:latest", images[0])
}

func TestEmptyImageList(t *testing.T) {
	ctx := context.Background()

	podmanObj := Podman{}

	images, err := podmanObj.Images(ctx, mockEmptyImagesList)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(images))
}

func TestVolumeList(t *testing.T) {
	ctx := context.Background()

	podmanObj := Podman{}

	volumes, err := podmanObj.Volumes(ctx, mockVolumeList)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(volumes))
	assert.Equal(t, "2398675946hj5gh435574gerrw36hg59", volumes[0])
}

func TestEmptyVolumeList(t *testing.T) {
	ctx := context.Background()

	podmanObj := Podman{}

	volumes, err := podmanObj.Volumes(ctx, mockEmptyVolumelist)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(volumes))
}
