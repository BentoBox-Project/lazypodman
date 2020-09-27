package gui

import (
	"github.com/jesseduffield/gocui"
)

// getFocusLayout returns a manager function for when view gain and lose focus
func (gui *Gui) getFocusLayout() func(g *gocui.Gui) error {
	var previousView *gocui.View
	return func(g *gocui.Gui) error {
		newView := gui.g.CurrentView()
		if err := gui.onFocusChange(); err != nil {
			return err
		}
		// for now we don't consider losing focus to a popup panel as actually losing focus
		if newView != previousView && !gui.isPopupPanel(newView.Name()) {
			if err := gui.onFocusLost(previousView, newView); err != nil {
				return err
			}
			if err := gui.onFocus(newView); err != nil {
				return err
			}
			previousView = newView
		}
		return nil
	}
}

func (gui *Gui) onFocusChange() error {
	currentView := gui.g.CurrentView()
	for _, view := range gui.g.Views() {
		view.Highlight = view == currentView && view.Name() != "main"
	}
	return nil
}

func (gui *Gui) onFocusLost(v *gocui.View, newView *gocui.View) error {
	if v == nil {
		return nil
	}

	if !gui.isPopupPanel(newView.Name()) {
		v.ParentView = nil
	}

	// refocusing because in responsive mode (when the window is very short) we want to ensure that after the view size changes we can still see the last selected item
	if err := gui.focusPointInView(v); err != nil {
		return err
	}

	gui.Log.Info(v.Name() + " focus lost")
	return nil
}

func (gui *Gui) onFocus(v *gocui.View) error {
	if v == nil {
		return nil
	}

	if err := gui.focusPointInView(v); err != nil {
		return err
	}

	gui.Log.Info(v.Name() + " focus gained")
	return nil
}

func (gui *Gui) layout(g *gocui.Gui) error {
	g.Highlight = true

	width, height := g.Size()

	minimumHeight := 9
	minimumWidth := 10

	if height < minimumHeight || width < minimumWidth {
		v, err := g.SetView("limit", 0, 0, width-1, height-1, 0)
		if err != nil {
			if err.Error() != "unknown view" {
				return err
			}
			v.Title = "Not enough space to render panels"
			v.Wrap = true
			_, _ = g.SetViewOnTop("limit")
		}
		return nil
	}

	currView := gui.g.CurrentView()
	currentCyclebleView := gui.peekPreviousView()

	if currView != nil {
		viewName := currView.Name()
		usePreviousView := true
		for _, view := range gui.CyclableViews {
			if view == viewName {
				currentCyclebleView = viewName
				usePreviousView = false
				break
			}
		}
		if usePreviousView {
			currentCyclebleView = gui.peekPreviousView()
		}
	}

	usableSpace := height - 4
	tallPanels := 3

	var vHeights map[string]int
	tallPanels++
	vHeights = map[string]int{
		"project":    3,
		"pods":       usableSpace / tallPanels,
		"services":   usableSpace/tallPanels + usableSpace%tallPanels,
		"containers": usableSpace / tallPanels,
		"images":     usableSpace / tallPanels,
		"volumes":    usableSpace / tallPanels,
	}

	if height < 28 {
		defaultHeight := 3
		if height < 21 {
			defaultHeight = 1
		}
		vHeights = map[string]int{
			"project":    defaultHeight,
			"pods":       defaultHeight,
			"containers": defaultHeight,
			"images":     defaultHeight,
			"voumes":     defaultHeight,
			"options":    defaultHeight,
		}
		vHeights[currentCyclebleView] = height - defaultHeight*tallPanels - 1
	}
	leftSideView := width / 3
	v, err := g.SetView("main", leftSideView+1, 0, width-1, height-2, gocui.LEFT)
	if err != nil {
		if err.Error() != "unknown view" {
			return err
		}
		v.Wrap = true
		v.FgColor = gocui.ColorDefault
		v.IgnoreCarriageReturns = true
	}

	if v, err := g.SetView("project", 0, 0, leftSideView, vHeights["project"]-1, gocui.BOTTOM|gocui.RIGHT); err != nil {
		if err.Error() != "unknown view" {
			return err
		}
		v.Title = "Pod prototype"
		v.FgColor = gocui.ColorDefault
	}

	aboveContainersView := "project"
	podsViews, err := g.SetViewBeneath("pods", aboveContainersView, vHeights["volumes"])
	if err != nil {
		if err.Error() != "unknown view" {
			return err
		}
		podsViews.Highlight = true
		podsViews.Title = "Pods"
		podsViews.FgColor = gocui.ColorDefault
	}

	containersView, err := g.SetViewBeneath("containers", "pods", vHeights["containers"])
	if err != nil {
		if err.Error() != "unknown view" {
			return err
		}
		containersView.Title = "Containers"
		containersView.Highlight = true
		containersView.FgColor = gocui.ColorDefault
	}

	imagesView, err := g.SetViewBeneath("images", "containers", vHeights["images"])
	if err != nil {
		if err.Error() != "unknown view" {
			return err
		}
		imagesView.Highlight = true
		imagesView.Title = "Images"
		imagesView.FgColor = gocui.ColorDefault
	}

	volumesView, err := g.SetViewBeneath("volumes", "images", vHeights["volumes"])
	if err != nil {
		if err.Error() != "unknown view" {
			return err
		}
		volumesView.Highlight = true
		volumesView.Title = "Volumes"
		volumesView.FgColor = gocui.ColorDefault
	}

	return nil
}

type listViewState struct {
	selectedLine int
	lineCount    int
}

func (gui *Gui) focusPointInView(view *gocui.View) error {
	if view == nil {
		return nil
	}

	listViews := map[string]listViewState{
		"containers": {selectedLine: gui.State.Panels.Containers.SelectedLine, lineCount: 4},
		"images":     {selectedLine: gui.State.Panels.Images.SelectedLine, lineCount: 4},
		"volumes":    {selectedLine: gui.State.Panels.Volumes.SelectedLine, lineCount: 4},
		"services":   {selectedLine: gui.State.Panels.Services.SelectedLine, lineCount: 4},
		"menu":       {selectedLine: gui.State.Panels.Menu.SelectedLine, lineCount: gui.State.MenuItemCount},
	}

	if state, ok := listViews[view.Name()]; ok {
		if err := gui.focusPoint(0, state.selectedLine, state.lineCount, view); err != nil {
			return err
		}
	}

	return nil
}
