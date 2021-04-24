package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"

	"github.com/jx2lee/elastic-cli/pkg/config"
	"github.com/manifoldco/promptui"
)

func init() {
	configCmd.AddCommand(configUseCmd)
	configCmd.AddCommand(configLsCmd)
	configCmd.AddCommand(configAddClusterCmd)
	configCmd.AddCommand(configRemoveClusterCmd)
	configCmd.AddCommand(configSelectCluster)
	configCmd.AddCommand(configCurrentContext)

	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Handle elasticsearch configuration",
}

var configCurrentContext = &cobra.Command{
	Use:   "current-context",
	Short: "Displays the current context",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfg.CurrentCluster)
	},
}

var configUseCmd = &cobra.Command{
	Use:               "use-cluster [NAME]",
	Short:             "Sets the current cluster in the configuration",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: validConfigArgs,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := cfg.SetCurrentCluster(name); err != nil {
			fmt.Printf("Cluster with name %v not found\n", name)
		} else {
			fmt.Printf("Switched to cluster \"%v\".\n", name)
		}
	},
}

var configLsCmd = &cobra.Command{
	Use:   "get-clusters",
	Short: "Display clusters in the configuration file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NAME")
		for _, cluster := range cfg.Clusters {
			fmt.Println(cluster.Name)
		}
	},
}

var configSelectCluster = &cobra.Command{
	Use:   "select-cluster",
	Short: "Interactively select a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		var clusterNames []string
		for _, cluster := range cfg.Clusters {
			clusterNames = append(clusterNames, cluster.Name)
		}

		searcher := func(input string, index int) bool {
			cluster := clusterNames[index]
			name := strings.Replace(strings.ToLower(cluster), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		}

		p := promptui.Select{
			Label:    "Select cluster",
			Items:    clusterNames,
			Searcher: searcher,
			Size:     12,
		}

		_, selected, err := p.Run()
		if err != nil {
			os.Exit(0)
		}

		// How to have selection on currently selected cluster?

		// TODO copy pasta
		if err := cfg.SetCurrentCluster(selected); err != nil {
			fmt.Printf("Cluster with selected %v not found\n", selected)
		}
	},
}

var configAddClusterCmd = &cobra.Command{
	Use:   "add-cluster [NAME]",
	Short: "Add cluster",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		for _, cluster := range cfg.Clusters {
			if cluster.Name == name {
				errorExit("Could not add cluster: cluster with name '%v' exists already.", name)
			}
		}

		cfg.Clusters = append(cfg.Clusters, &config.Cluster{
			Name:              name,
			Nodes:             nodeFlag,
		})
		err := cfg.Write()
		if err != nil {
			errorExit("Unable to write config: %v\n", err)
		}
		fmt.Println("Added cluster.")
	},
}

var configRemoveClusterCmd = &cobra.Command{
	Use:               "remove-cluster [NAME]",
	Short:             "remove cluster",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: validConfigArgs,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		var pos = -1
		for i, cluster := range cfg.Clusters {
			if cluster.Name == name {
				pos = i
				break
			}
		}

		if pos == -1 {
			errorExit("Could not delete cluster: cluster with name '%v' not exists.", name)
		}

		cfg.Clusters = append(cfg.Clusters[:pos], cfg.Clusters[pos+1:]...)

		err := cfg.Write()
		if err != nil {
			errorExit("Unable to write config: %v\n", err)
		}
		fmt.Println("Removed cluster.")
	},
}
