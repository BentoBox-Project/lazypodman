package gui

import (
	"time"

	"github.com/danvergara/lazypodman/pkg/podman"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/sirupsen/logrus"
)

// Gui wrapes the gocui object which handles rendering and events
type Gui struct {
	Grid          *ui.Grid
	PodmanBinding *podman.Podman
	Pods          *widgets.List
	Containers    *widgets.List
	Images        *widgets.List
	Volumes       *widgets.List
}

// NewGui returns a new Gui object
func NewGui() *Gui {
	gui := &Gui{
		Grid:          ui.NewGrid(),
		Pods:          widgets.NewList(),
		Containers:    widgets.NewList(),
		Images:        widgets.NewList(),
		Volumes:       widgets.NewList(),
		PodmanBinding: &podman.Podman{},
	}
	return gui
}

// render display the Grid on the terminal
func (g *Gui) render() {
	ctx, err := podman.APIConn()
	if err != nil {
		logrus.Error(err)
	}

	b := podman.Bindings{}

	podNames, err := g.PodmanBinding.Pods(ctx, &b)

	if err != nil {
		logrus.Error(err)
	}

	containerNames, err := g.PodmanBinding.Containers(ctx, &b)

	if err != nil {
		logrus.Error(err)
	}

	imageNames, err := g.PodmanBinding.Images(ctx, &b)
	if err != nil {
		logrus.Error(err)
	}

	vNames, err := g.PodmanBinding.Volumes(ctx, &b)
	if err != nil {
		logrus.Error(err)
	}

	g.Pods.Rows = podNames
	g.Containers.Rows = containerNames
	g.Images.Rows = imageNames
	g.Volumes.Rows = vNames

	ui.Render(g.Grid)
}

// Run run the cli graphic interface
func (g *Gui) Run() error {
	if err := g.layout(); err != nil {
		return err
	}

	if err := ui.Init(); err != nil {
		return err
	}

	defer ui.Close()

	ev := ui.PollEvents()
	//lint:ignore SA1015 we want to hide the actual implementation from the main package
	tick := time.Tick(time.Second)

	for {
		select {
		case e := <-ev:
			switch e.Type {
			case ui.KeyboardEvent:
				// quit on any keyboard event
				return nil
			}
		case <-tick:
			g.render()
		}
	}
}
