package cmd

import (
	"github.com/universal-operator/universal-operator/cmd/maestroctl/cmd/plan"
	"github.com/spf13/cobra"
)

func NewPlanCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "plan",
		Short: "-> View all available plans.",
		Long:  `The plan command has subcommands to view all available plans.`,
	}

	newCmd.AddCommand(plan.NewPlanHistoryCmd())
	newCmd.AddCommand(plan.NewPlanStatusCmd())

	return newCmd
}
