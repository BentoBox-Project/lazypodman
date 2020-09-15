package gui

import (
	"time"

	"github.com/danvergara/lazypodman/pkg/podman"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/sirupsen/logrus"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

// Controller is the interface used to represent the Gui objects
type Controller interface {
	Render()
	Resize()
}

// Gui wrapes the gocui object which handles rendering and events
type Gui struct {
	Grid          *ui.Grid
	PodmanBinding *podman.Podman
	Pods          *widgets.List
	Services      *widgets.List
	Images        *widgets.List
	Volumes       *widgets.List
}

// NewGui returns a new Gui object
func NewGui(composeFile string) *Gui {
	gui := &Gui{
		Grid:     ui.NewGrid(),
		Pods:     widgets.NewList(),
		Services: widgets.NewList(),
		Images:   widgets.NewList(),
		Volumes:  widgets.NewList(),
		PodmanBinding: &podman.Podman{
			ComposeFile: composeFile,
		},
	}

	gui.initUI()

	return gui
}

func (g *Gui) initUI() {
	g.Pods.Title = "Pods"
	g.Services.Title = "Services"
	g.Images.Title = "Images"
	g.Volumes.Title = "Volumes"

	g.Grid.Set(
		ui.NewCol(.3,
			ui.NewRow(1.0/4, g.Pods),
			ui.NewRow(1.0/4, g.Services),
			ui.NewRow(1.0/4, g.Images),
			ui.NewRow(1.0/4, g.Volumes),
		),
	)
	width, _ := terminal.Width()
	height, _ := terminal.Height()

	g.Grid.SetRect(0, 0, int(width), int(height))
}

// render display the Grid on the terminal
func (g *Gui) render() {
	conn, err := podman.APIConn()
	if err != nil {
		logrus.Error(err)
	}

	podNames, err := g.PodmanBinding.Pods(conn)

	if err != nil {
		logrus.Error(err)
	}

	cNames, err := g.PodmanBinding.Containers(conn)

	if err != nil {
		logrus.Error(err)
	}

	imageNames, err := g.PodmanBinding.Images(conn)
	if err != nil {
		logrus.Error(err)
	}

	vNames, err := g.PodmanBinding.Volumes(conn)
	if err != nil {
		logrus.Error(err)
	}

	g.Pods.Rows = podNames
	g.Services.Rows = cNames
	g.Images.Rows = imageNames
	g.Volumes.Rows = vNames

	ui.Render(g.Grid)
}

// Run run the cli graphic interface
func (g *Gui) Run() error {

	if err := ui.Init(); err != nil {
		return err
	}

	defer ui.Close()

	ev := ui.PollEvents()
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