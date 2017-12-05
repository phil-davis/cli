package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.0.1"

func newVersionCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Args:  cobra.NoArgs,
		RunE:  cli.wrap(runVersion),
	}
	return cmd
}

func runVersion(cli *CLI, cmd *cobra.Command, args []string) error {
	fmt.Printf("hcloud %s\n", Version)
	return nil
}