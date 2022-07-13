package nodecmd

import (
	json_rpc "github.com/forta-network/forta-node/cmd/json-rpc"
	"github.com/forta-network/forta-node/cmd/publisher"
	"github.com/forta-network/forta-node/cmd/scanner"
	sla_tracker "github.com/forta-network/forta-node/cmd/sla-tracker"
	"github.com/forta-network/forta-node/cmd/supervisor"
	"github.com/forta-network/forta-node/cmd/updater"
	"github.com/spf13/cobra"
)

var (
	cmdFortaNode = &cobra.Command{
		Use: "forta-node",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceUsage: true,
	}

	cmdUpdater = &cobra.Command{
		Use: "updater",
		RunE: func(cmd *cobra.Command, args []string) error {
			updater.Run()
			return nil
		},
	}

	cmdSupervisor = &cobra.Command{
		Use: "supervisor",
		RunE: func(cmd *cobra.Command, args []string) error {
			supervisor.Run()
			return nil
		},
	}

	cmdScanner = &cobra.Command{
		Use: "scanner",
		RunE: func(cmd *cobra.Command, args []string) error {
			scanner.Run()
			return nil
		},
	}

	cmdPublisher = &cobra.Command{
		Use: "publisher",
		RunE: func(cmd *cobra.Command, args []string) error {
			publisher.Run()
			return nil
		},
	}

	cmdSLATracker = &cobra.Command{
		Use: "sla-tracker",
		RunE: func(cmd *cobra.Command, args []string) error {
			sla_tracker.Run()
			return nil
		},
	}
	cmdJsonRpc = &cobra.Command{
		Use: "json-rpc",
		RunE: func(cmd *cobra.Command, args []string) error {
			json_rpc.Run()
			return nil
		},
	}
)

func init() {
	cmdFortaNode.AddCommand(cmdUpdater)
	cmdFortaNode.AddCommand(cmdSupervisor)
	cmdFortaNode.AddCommand(cmdScanner)
	cmdFortaNode.AddCommand(cmdPublisher)
	cmdFortaNode.AddCommand(cmdSLATracker)
	cmdFortaNode.AddCommand(cmdJsonRpc)
}

func Run() error {
	return cmdFortaNode.Execute()
}
