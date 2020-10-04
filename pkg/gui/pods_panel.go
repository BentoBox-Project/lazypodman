package gui

import (
	"fmt"
	"strings"

	"github.com/containers/podman/v2/pkg/bindings/pods"
	"github.com/danvergara/lazypodman/pkg/podman"
	"github.com/jesseduffield/gocui"
)

func (gui *Gui) refreshPodsView() error {
	v, err := gui.g.View("pods")
	if err != nil {
		return err
	}

	ctx, err := podman.APIConn()
	if err != nil {
		return err
	}

	pods, err := gui.PodmanBinding.Pods(ctx, pods.List)
	if err != nil {
		return err
	}

	gui.g.Update(func(g *gocui.Gui) error {
		v.Clear()
		fmt.Fprintf(v, strings.Join(pods, "\n"))
		return nil
	})
	return nil
}
