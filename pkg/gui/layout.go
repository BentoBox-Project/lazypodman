package gui

import (
	ui "github.com/gizak/termui/v3"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

func (g *Gui) layout() error {
	g.Pods.Title = "Pods"
	g.Containers.Title = "Containers"
	g.Images.Title = "Images"
	g.Volumes.Title = "Volumes"

	g.Grid.Set(
		ui.NewCol(.3,
			ui.NewRow(1.0/4, g.Pods),
			ui.NewRow(1.0/4, g.Containers),
			ui.NewRow(1.0/4, g.Images),
			ui.NewRow(1.0/4, g.Volumes),
		),
	)
	width, err := terminal.Width()
	if err != nil {
		return err
	}

	height, err := terminal.Height()
	if err != nil {
		return err
	}

	g.Grid.SetRect(0, 0, int(width), int(height))
	return nil
}
