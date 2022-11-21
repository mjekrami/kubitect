package cmd

import (
	"cli/actions"
	"cli/env"
	"cli/file"
	"cli/ui"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	purgeShort = "Purge the cluster directory"
	purgeLong  = LongDesc(`
		Purge the directory of a given cluster.
		Directories of active clusters cannot be purged.`)

	purgeExample = Example(`
		Purge the directory of cluster 'cls-name':
		> kubitect purge --cluster cls-name`)
)

type PurgeOptions struct {
	ClusterName string

	env.ContextOptions
}

func NewPurgeCmd() *cobra.Command {
	var opts PurgeOptions

	cmd := &cobra.Command{
		Use:     "purge",
		GroupID: "support",
		Short:   purgeShort,
		Long:    purgeLong,
		Example: purgeExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.PersistentFlags().StringVar(&opts.ClusterName, "cluster", "", "specify the cluster to be used")
	cmd.PersistentFlags().BoolVar(&env.AutoApprove, "auto-approve", false, "automatically approve any user permission requests")

	cmd.RegisterFlagCompletionFunc("cluster", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var names []string

		clusters, err := actions.Clusters(opts.Context())

		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		for _, c := range clusters {
			if !c.ContainsTfStateConfig() {
				names = append(names, c.Name)
			}
		}

		return names, cobra.ShellCompDirectiveNoFileComp
	})

	return cmd
}

func (o *PurgeOptions) Run() error {
	cs, err := actions.Clusters(o.Context())

	if err != nil {
		return err
	}

	c := cs.FindByName(o.ClusterName)

	if c == nil {
		return fmt.Errorf("cluster '%s' does not exist", c.Name)
	}

	count := cs.CountByName(c.Name)

	if count > 1 {
		return fmt.Errorf("multiple clusters (%d) have been found with the name '%s'", count, c.Name)
	}

	if c.ContainsTfStateConfig() {
		return fmt.Errorf("cluster '%s' cannot be purged: only destroyed clusters can be purged", c.Name)
	}

	fmt.Printf("Cluster '%s' will be purged. This will remove cluster's directory including all of its content.", c.Name)

	if err := ui.Ask(); err != nil {
		return err
	}

	fmt.Printf("Purging cluster '%s'...\n", c.Name)

	if err := file.Remove(c.Path); err != nil {
		return fmt.Errorf("failed to purge cluster '%s': %v", c.Name, err)
	}

	fmt.Printf("Cluster '%s' has been successfully purged.\n", c.Name)

	return nil
}
