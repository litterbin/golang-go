package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type clusterCmd struct {
	clusterName string
}

func (*clusterCmd) Name() string { return "cluster" }
func (*clusterCmd) Synopsis() string {
	return "Manage cluster"
}

func (*clusterCmd) Usage() string {
	return "Run 'container COMMAND --help' for more information on a command."
}

func (c *clusterCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.clusterName, "name", "", "cluster name")
}

func (c *clusterCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	commander := subcommands.NewCommander(f, c.Name())
	commander.Register(commander.HelpCommand(), "")
	commander.Register(commander.FlagsCommand(), "")
	commander.Register(commander.CommandsCommand(), "")

	pcmd := &printCmd{
		clusterName: c.clusterName,
	}
	commander.Register(pcmd, "")

	return commander.Execute(ctx, args...)
}
