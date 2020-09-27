package commands

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/danvergara/lazypodman/pkg/config"
	"github.com/sirupsen/logrus"
)

// OSCommand holds all the os commands
type OSCommand struct {
	Log     *logrus.Entry
	Config  *config.Config
	command func(string, ...string) *exec.Cmd
	getenv  func(string) string
}

// NewOSCommand os command runner
func NewOSCommand(log *logrus.Entry, config *config.Config) *OSCommand {
	return &OSCommand{
		Log:     log,
		Config:  config,
		command: exec.Command,
		getenv:  os.Getenv,
	}
}

// Kill kills a process. If the process has Setpgid == true, then we have anticipated that it might spawn its own child processes, so we've given it a process group ID (PGID) equal to its process id (PID) and given its child processes will inherit the PGID, we can kill that group, rather than killing the process itself.
func (c *OSCommand) Kill(cmd *exec.Cmd) error {
	if cmd.Process == nil {
		// somebody got to it before we were able to, poor bastard
		return nil
	}
	if cmd.SysProcAttr != nil && cmd.SysProcAttr.Setpgid {
		// minus sign means we're talking about a PGID as opposed to a PID
		return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}

	return cmd.Process.Kill()
}
