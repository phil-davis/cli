package cli

import "github.com/spf13/cobra"

func newLoadBalancerCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "load-balancer",
		Short:                 "Manage Load Balancers",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runLoadBalancer),
	}
	cmd.AddCommand(
		newLoadBalancerCreateCommand(cli),
		newLoadBalancerListCommand(cli),
		newLoadBalancerDescribeCommand(cli),
		newLoadBalancerDeleteCommand(cli),
		newLoadBalancerUpdateCommand(cli),
		newLoadBalancerAddLabelCommand(cli),
		newLoadBalancerRemoveLabelCommand(cli),
		newLoadBalancerAddTargetCommand(cli),
		newLoadBalancerRemoveTargetCommand(cli),
		newLoadBalancerChangeAlgorithmCommand(cli),
		newLoadBalancerUpdateServiceCommand(cli),
		newLoadBalancerDeleteServiceCommand(cli),
		newLoadBalancerAddServiceCommand(cli),
		newLoadBalancerEnableProtectionCommand(cli),
		newLoadBalancerDisableProtectionCommand(cli),
		newLoadBalancerAttachToNetworkCommand(cli),
		newLoadBalancerDetachFromNetworkCommand(cli),
		newLoadBalancerEnablePublicInterface(cli),
		newLoadBalancerDisablePublicInterface(cli),
	)
	return cmd
}

func runLoadBalancer(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
