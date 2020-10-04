package gui

import (
	"fmt"
	"strings"

	"github.com/containers/podman/v2/pkg/bindings/volumes"
	"github.com/danvergara/lazypodman/pkg/podman"
	"github.com/jesseduffield/gocui"
)

func (gui *Gui) refreshVolumesView() error {
	v, err := gui.g.View("volumes")

	if err != nil {
		return err
	}

	ctx, err := podman.APIConn()
	if err != nil {
		return err
	}

	volumes, err := gui.PodmanBinding.Volumes(ctx, volumes.List)
	if err != nil {
		return err
	}

	gui.g.Update(func(g *gocui.Gui) error {
		v.Clear()
		fmt.Fprint(v, strings.Join(volumes, "\n"))
		return nil
	})

	return nil
}
