package gui

import (
	"fmt"
	"strings"

	"github.com/containers/libpod/v2/pkg/bindings/containers"
	"github.com/danvergara/lazypodman/pkg/podman"
	"github.com/jesseduffield/gocui"
)

func (gui *Gui) refreshContainersView() error {
	v, err := gui.g.View("containers")
	if err != nil {
		return err
	}

	ctx, err := podman.APIConn()
	if err != nil {
		return err
	}

	services, err := gui.PodmanBinding.Containers(ctx, containers.List)
	if err != nil {
		return err
	}

	gui.g.Update(func(g *gocui.Gui) error {
		v.Clear()
		fmt.Fprint(v, strings.Join(services, "\n"))
		return nil
	})
	return nil
}
