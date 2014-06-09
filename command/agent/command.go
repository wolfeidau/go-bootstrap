package agent

import (
	"strings"

	"github.com/juju/loggo"
	"github.com/mitchellh/cli"
)

type Command struct {
	Ui         cli.Ui
	ShutdownCh <-chan struct{}
	Log        loggo.Logger
}

func (c *Command) Run(_ []string) int {
	c.Log = loggo.GetLogger("agent")
	c.Log.Warningf("test")
	return 0
}

func (c *Command) Synopsis() string {
	return "Runs a myapp agent"
}

func (c *Command) Help() string {
	helpText := `
Usage: myapp agent [options]

  Starts the myapp agent and runs until an interrupt is received.

Options:

  -log-level=info          Log level of the agent.
`
	return strings.TrimSpace(helpText)
}
