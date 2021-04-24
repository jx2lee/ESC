package main

import (
	"fmt"
	Form "github.com/jx2lee/elastic-cli/pkg/schema"
	"github.com/jx2lee/elastic-cli/pkg/util"
	"github.com/spf13/cobra"
	"reflect"
	"text/tabwriter"
)

func init() {
	rootCmd.AddCommand(catCommand)

	catCommand.AddCommand(catHealthCommand)
	catCommand.AddCommand(catIndicesCommand)
	catCommand.AddCommand(catIndexCommand)
	catCommand.AddCommand(catNodesCommand)
	catCommand.AddCommand(catMasterCommand)

	catCommand.PersistentFlags().BoolVarP(&noHeaderFlag, "no-headers", "", false, "Hide table headers")
}

var catCommand = &cobra.Command{
	Use:   "cat",
	Short: "Cat APIs for elasticsearch cluster",
}

var catHealthCommand = &cobra.Command{
	Use:   "health",
	Short: "Health for elasticsearch",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()

		health, err := client.Cat.Health(client.Cat.Health.WithFormat("json"))
		if err != nil {
			fmt.Println(err)
		}

		var healthFormData []Form.CatHealthForm
		util.ConvertJSONtoFormData(health.Body, &healthFormData)

		e := reflect.ValueOf(&healthFormData[0]).Elem()
		filedNum := e.NumField()

		w := tabwriter.NewWriter(outWriter, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
		if !noHeaderFlag {
			fmt.Fprintf(w, "%v\t%v\n", "Health API Component",  "Value")
		}

		for i := 0; i < filedNum; i++ {
			v := e.Field(i)
			t := e.Type().Field(i)

			fmt.Fprintf(w, "%v\t: %v\n", t.Name, fmt.Sprintf("%s", v.Interface()))
		}

		w.Flush()
	},
}

var catIndicesCommand = &cobra.Command{
	Use:   "indices",
	Short: "All Indices for elasticsearch",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()

		indices, err := client.Cat.Indices(client.Cat.Indices.WithFormat("json"))
		if err != nil {
			fmt.Println(err)
		}

		var indicesFormData []Form.CatIndexForm
		util.ConvertJSONtoFormData(indices.Body, &indicesFormData)

		w := tabwriter.NewWriter(outWriter, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
		if !noHeaderFlag {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n",
				"Index","health", "status", "pri", "rep", "store.size")
		}

		for _, index := range indicesFormData {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n",
				index.Index,
				index.Health,
				index.StoreSize,
				index.PrimaryShards,
				index.ReplicaShards,
				index.PriStoreSize)
		}

		w.Flush()
	},
}

var catIndexCommand = &cobra.Command{
	Use:   "index [INDEX_NAME]",
	Short: "Index for elasticsearch",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		indexName := args[0]
		client := getClient()

		indices, err := client.Cat.Indices(client.Cat.Indices.WithFormat("json"))
		if err != nil {
			fmt.Println(err)
		}

		var indicesFormData []Form.CatIndexForm
		util.ConvertJSONtoFormData(indices.Body, &indicesFormData)

		w := tabwriter.NewWriter(outWriter, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
		if !noHeaderFlag {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n",
				"Index","health", "status", "pri", "rep", "store.size")
		}

		for _, index := range indicesFormData {
			if index.Index ==  indexName {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n",
					index.Index,
					index.Health,
					index.StoreSize,
					index.PrimaryShards,
					index.ReplicaShards,
					index.PriStoreSize)
			}
		}

		w.Flush()
	},
}

var catNodesCommand = &cobra.Command{
	Use:   "nodes",
	Short: "Nodes for elasticsearch",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()

		//nodes, err := client.Cat.Nodes(client.Cat.Nodes.WithFormat("json"))
		nodes, err := client.Cat.Nodes(client.Cat.Nodes.WithH(
			"name","ip","role","uptime","disk.used_percent","load_1m","load_5m","load_15m"),
			client.Cat.Nodes.WithFormat("json"))

		if err != nil {
			fmt.Println(err)
		}

		var nodeFormData []Form.Node
		util.ConvertJSONtoFormData(nodes.Body, &nodeFormData)
		fmt.Println(nodeFormData)

		w := tabwriter.NewWriter(outWriter, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
		if !noHeaderFlag {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
				"Name", "IP","NodeRole", "Uptime(m)", "DiskUsedPercent", "Load1M", "Load5M", "Load15M")
		}

		for _, node := range nodeFormData {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
				node.Name,
				node.IP,
				node.NodeRole,
				node.DiskUsedPercent,
				node.Load1M,
				node.Load5M,
				node.Load15M,
				node.Uptime)
		}
		w.Flush()
	},
}

var catMasterCommand = &cobra.Command{
	Use:   "master",
	Short: "Master Node for elasticsearch",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()

		indices, err := client.Cat.Master(client.Cat.Master.WithFormat("json"))
		if err != nil {
			fmt.Println(err)
		}

		var masterFormData []Form.CatMasterForm
		util.ConvertJSONtoFormData(indices.Body, &masterFormData)

		w := tabwriter.NewWriter(outWriter, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
		if !noHeaderFlag {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\n",
				"Id","Host", "Ip", "Node")
			//fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\n",
			//	"--", "----", "--", "----")
		}

		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\n",
					masterFormData[0].ID,
					masterFormData[0].Host,
					masterFormData[0].Node,
					masterFormData[0].Host)

		w.Flush()
	},
}
