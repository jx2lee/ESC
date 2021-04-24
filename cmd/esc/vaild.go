package main

import "github.com/spf13/cobra"

func validConfigArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	clusterList := make([]string, 0, len(cfg.Clusters))
	for _, cluster := range cfg.Clusters {
		clusterList = append(clusterList, cluster.Name)
	}
	return clusterList, cobra.ShellCompDirectiveNoFileComp
}

const (
	tabwriterMinWidth       = 6
	tabwriterWidth          = 4
	tabwriterPadding        = 3
	tabwriterPadChar        = ' '
	tabwriterFlags          = 0
)
