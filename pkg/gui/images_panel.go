package gui

import (
	"fmt"
	"strings"

	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/danvergara/lazypodman/pkg/podman"
	"github.com/jesseduffield/gocui"
)

func (gui *Gui) refreshImagesView() error {
	v, err := gui.g.View("images")
	if err != nil {
		return err
	}

	ctx, err := podman.APIConn()
	if err != nil {
		return err
	}

	images, err := gui.PodmanBinding.Images(ctx, images.List)
	if err != nil {
		return err
	}

	gui.g.Update(func(g *gocui.Gui) error {
		v.Clear()
		fmt.Fprint(v, strings.Join(images, "\n"))
		return nil
	})

	return nil
}
