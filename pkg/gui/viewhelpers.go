package gui

import (
	"github.com/danvergara/lazypodman/pkg/utils"
	"github.com/jesseduffield/gocui"
)

func (gui *Gui) currentViewName() string {
	currentView := gui.g.CurrentView()
	return currentView.Name()
}

// if the cursor down past the last item, move it to the last line
func (gui *Gui) focusPoint(selectedX int, selectedY int, lineCount int, v *gocui.View) error {
	if selectedY < 0 || selectedY > lineCount {
		return nil
	}
	ox, oy := v.Origin()
	originalOy := oy
	cx, cy := v.Cursor()
	originalCy := cy
	_, height := v.Size()

	ly := utils.Max(height-1, 0)

	windowStart := oy
	windowEnd := oy + ly

	if selectedY < windowStart {
		oy = utils.Max(oy-(windowStart-selectedY), 0)
	} else if selectedY > windowEnd {
		oy += (selectedY - windowEnd)
	}

	if windowEnd > lineCount-1 {
		shiftAmount := (windowEnd - (lineCount - 1))
		oy = utils.Max(oy-shiftAmount, 0)
	}

	if originalOy != oy {
		_ = v.SetOrigin(ox, oy)
	}

	cy = selectedY - oy
	if originalCy != cy {
		_ = v.SetCursor(cx, selectedY-oy)
	}
	return nil
}

func (gui *Gui) isPopupPanel(viewName string) bool {
	return viewName == "confirmation" || viewName == "menu"
}

func (gui *Gui) peekPreviousView() string {
	if gui.State.PreviousViews.Len() > 0 {
		return gui.State.PreviousViews.Peek().(string)
	}

	return ""
}

func (gui *Gui) popPreviousView() string {
	if gui.State.PreviousViews.Len() > 0 {
		return gui.State.PreviousViews.Pop().(string)
	}

	return ""
}

func (gui *Gui) pushPreviousView(name string) {
	gui.State.PreviousViews.Push(name)
}
