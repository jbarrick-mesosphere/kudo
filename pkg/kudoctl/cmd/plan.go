package cmd

import (
	"github.com/kudobuilder/kudo/pkg/kudoctl/cmd/plan"
	"github.com/spf13/cobra"
)

// newPlanCmd creates a new command that shows the plans available for an instance
func newPlanCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "plan",
		Short: "View all available plans.",
		Long:  `The plan command has subcommands to view all available plans.`,
	}

	newCmd.AddCommand(NewPlanHistoryCmd())
	newCmd.AddCommand(NewPlanStatusCmd())

	return newCmd
}

// NewPlanHistoryCmd creates a command that shows the plan history of an instance.
func NewPlanHistoryCmd() *cobra.Command {
	options := plan.DefaultHistoryOptions
	listCmd := &cobra.Command{
		Use:   "history",
		Short: "Lists history to a specific operator-version of an instance.",
		Long: `
	# View plan status
	kudoctl plan history <operatorVersion> --instance=<instanceName>`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return plan.RunHistory(cmd, args, options, &Settings)
		},
	}

	listCmd.Flags().StringVar(&options.Instance, "instance", "", "The instance name.")
	listCmd.Flags().StringVar(&options.Namespace, "namespace", "default", "The namespace where the operator watches for changes.")

	return listCmd
}

//NewPlanStatusCmd creates a new command that shows the status of an instance by looking at its current plan
func NewPlanStatusCmd() *cobra.Command {
	options := plan.DefaultStatusOptions
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Shows the status of all plans to an particular instance.",
		Long: `
	# View plan status
	kudoctl plan status --instance=<instanceName>`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return plan.RunStatus(cmd, args, options, &Settings)
		},
	}

	statusCmd.Flags().StringVar(&options.Instance, "instance", "", "The instance name available from 'kubectl get instances'")
	statusCmd.Flags().StringVar(&options.Namespace, "namespace", "default", "The namespace where the instance is running.")

	return statusCmd
}
