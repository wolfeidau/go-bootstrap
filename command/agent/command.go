package agent

import (
	"flag"
	"strings"

	"github.com/juju/loggo"
	"github.com/mitchellh/cli"
)

type Command struct {
	Ui         cli.Ui
	ShutdownCh <-chan struct{}
	Log        loggo.Logger
	debug      bool
}

func (c *Command) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("agent", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.BoolVar(&c.debug, "debug", false, "Enable debug level logging for the agent.")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if c.debug {
		// set the root logger to debug
		loggo.GetLogger("").SetLogLevel(loggo.DEBUG)
	}

	// configure the agent logger
	c.Log = loggo.GetLogger("agent")

	// main logic here

	return 0
}

func (c *Command) Synopsis() string {
	return "Runs a appname agent"
}

func (c *Command) Help() string {
	helpText := `
Usage: appname agent [options]

  Starts the appname agent and runs until an interrupt is received.

Options:

  -debug              Enable debug level logging for the agent.
`
	return strings.TrimSpace(helpText)
}
