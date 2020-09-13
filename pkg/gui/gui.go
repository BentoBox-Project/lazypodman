package gui

import (
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Controller is the interface used to represent the Gui objects
type Controller interface {
	Render()
	Resize()
}

// Gui wrapes the gocui object which handles rendering and events
type Gui struct {
	Grid     *ui.Grid
	Pods     *widgets.List
	Services *widgets.List
	Images   *widgets.List
	Volumes  *widgets.List
}

// NewGui returns a new Gui object
func NewGui() *Gui {
	gui := &Gui{
		Grid:     ui.NewGrid(),
		Pods:     widgets.NewList(),
		Services: widgets.NewList(),
		Images:   widgets.NewList(),
		Volumes:  widgets.NewList(),
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

	termWidth, termHeight := ui.TerminalDimensions()
	g.Grid.SetRect(0, 0, termWidth, termHeight)
}

// render display the Grid on the terminal
func (g *Gui) render() {
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
