package gui

import (
	"errors"
	"os"
	"os/exec"

	"github.com/danvergara/lazypodman/pkg/commands"
	"github.com/danvergara/lazypodman/pkg/podman"
	"github.com/golang-collections/collections/stack"
	"github.com/jesseduffield/gocui"
	"github.com/sirupsen/logrus"
)

// SentinelErrors are the errors that have special meaning and need to be checked
// by calling functions. The less of these, the better
type SentinelErrors struct {
	ErrSubProcess   error
	ErrNoContainers error
	ErrNoImages     error
	ErrNoVolumes    error
}

// GenerateSentinelErrors makes the sentinel errors for the gui. We're defining it here
// because we can't do package-scoped errors with localization, and also because
// it seems like package-scoped variables are bad in general
// https://dave.cheney.net/2017/06/11/go-without-package-scoped-variables
// In the future it would be good to implement some of the recommendations of
// that article. For now, if we don't need an error to be a sentinel, we will just
// define it inline. This has implications for error messages that pop up everywhere
// in that we'll be duplicating the default values. We may need to look at
// having a default localisation bundle defined, and just using keys-only when
// localising things in the code.
func (gui *Gui) GenerateSentinelErrors() {
	gui.Errors = SentinelErrors{
		ErrSubProcess:   errors.New("running subprocess"),
		ErrNoContainers: errors.New("No containers"),
		ErrNoImages:     errors.New("No Images"),
		ErrNoVolumes:    errors.New("No volumes"),
	}
}

// Gui wrapes the gocui object which handles rendering and events
type Gui struct {
	g             *gocui.Gui
	Errors        SentinelErrors
	Log           *logrus.Entry
	PodmanBinding *podman.Podman
	State         guiState
	CyclableViews []string
	SubProcess    *exec.Cmd
	OSCommand     *commands.OSCommand
}

type servicePanelState struct {
	SelectedLine int
	ContextIndex int // for specifying if you are looking at logs/stats/config/etc
}

type containerPanelState struct {
	SelectedLine int
	ContextIndex int // for specifying if you are looking at logs/stats/config/etc
}

type podPanelState struct {
	SelectedLine int
	ContextIndex int // for specifying if you are looking at logs/stats/config/etc
}

type projectState struct {
	ContextIndex int // for specifying if you are looking ate credits/logs
}

type menuPanelState struct {
	SelectedLine int
	OnPress      func(*gocui.Gui, *gocui.View) error
}

type mainPanelState struct {
	// ObjectKey tells us what context we are in. For example, if we are looking at the logs a particular services in the services panel this key might be 'services-<service id>-logs'. The key is made so that if something changes which might require us to re-run the logs command or run a different command, the key will be different, and we'll then know to do whatever is required. Object key probably isn't the best name for this but Context is already used to refer to tabs. Maybe I should just call them tabs.
	ObjectKey string
}

type imagePanelState struct {
	SelectedLine int
	ContextIndex int // for specifying if you are looking at logs/stats/config/etc
}

type volumePanelState struct {
	SelectedLine int
	ContextIndex int // for specifying if you are looking at logs/stats/config/etc
}

type panelState struct {
	Pods       *podPanelState
	Services   *servicePanelState
	Containers *containerPanelState
	Menu       *menuPanelState
	Main       *mainPanelState
	Images     *imagePanelState
	Volumes    *volumePanelState
	Project    *projectState
}

type guiState struct {
	MenuItemCount int
	Panels        *panelState
	PreviousViews *stack.Stack
	SessionIndex  int
}

// NewGui returns a new Gui object
func NewGui(composeFile string) *Gui {
	initalState := guiState{
		Panels: &panelState{
			Pods:       &podPanelState{SelectedLine: -1},
			Containers: &containerPanelState{SelectedLine: -1},
			Images:     &imagePanelState{SelectedLine: -1},
			Volumes:    &volumePanelState{SelectedLine: -1},
			Menu:       &menuPanelState{},
			Main: &mainPanelState{
				ObjectKey: "",
			},
			Project: &projectState{},
		},
		PreviousViews: stack.New(),
	}

	gui := &Gui{
		PodmanBinding: &podman.Podman{
			ComposeFile: composeFile,
		},
		State:         initalState,
		CyclableViews: []string{"project", "pod", "containers", "images", "volumes"},
	}

	gui.GenerateSentinelErrors()

	return gui
}

// Run run the cli graphic interface
func (gui *Gui) Run() error {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		return err
	}
	defer g.Close()

	g.Mouse = true

	gui.g = g

	g.SetManager(gocui.ManagerFunc(gui.layout), gocui.ManagerFunc(gui.getFocusLayout()))

	if err := g.SetKeybinding("", []string{}, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		// log.Panicln(err)
		os.Exit(0)
		return nil
	}

	err = g.MainLoop()
	return err
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
