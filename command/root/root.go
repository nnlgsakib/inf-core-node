package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Infinity-Green/inf/command/backup"
	"github.com/Infinity-Green/inf/command/bridge"
	"github.com/Infinity-Green/inf/command/genesis"
	"github.com/Infinity-Green/inf/command/helper"
	"github.com/Infinity-Green/inf/command/ibft"
	"github.com/Infinity-Green/inf/command/license"
	"github.com/Infinity-Green/inf/command/monitor"
	"github.com/Infinity-Green/inf/command/peers"
	"github.com/Infinity-Green/inf/command/polybft"
	"github.com/Infinity-Green/inf/command/polybftsecrets"
	"github.com/Infinity-Green/inf/command/regenesis"
	"github.com/Infinity-Green/inf/command/rootchain"
	"github.com/Infinity-Green/inf/command/secrets"
	"github.com/Infinity-Green/inf/command/server"
	"github.com/Infinity-Green/inf/command/status"
	"github.com/Infinity-Green/inf/command/txpool"
	"github.com/Infinity-Green/inf/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "inf is the core client node of INFINITY GREEN CHAIN",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
		polybftsecrets.GetCommand(),
		polybft.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
