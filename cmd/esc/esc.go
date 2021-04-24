package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jx2lee/elastic-cli/pkg/config"
	"github.com/spf13/cobra"
	"io"
	"os"

	"github.com/mattn/go-colorable"
)

var cfgFile string

func getElasticConfig() (elasticConfig elasticsearch.Config) {
	cluster := currentCluster
	elasticConfig = elasticsearch.Config{
		Addresses: []string{"http://" + cluster.Nodes}}
	return elasticConfig
}

var (
	outWriter io.Writer = os.Stdout
	errWriter io.Writer = os.Stderr
	inReader  io.Reader = os.Stdin

	colorableOut io.Writer = colorable.NewColorableStdout()
)

var rootCmd = &cobra.Command{
	Use:   "esc",
	Short: "ElasticSearch Command Line utility for cluster management",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		outWriter = cmd.OutOrStdout()
		errWriter = cmd.ErrOrStderr()
		inReader = cmd.InOrStdin()

		if outWriter != os.Stdout {
			colorableOut = outWriter
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var currentCluster *config.Cluster
var cfg config.Config

var (
	nodeFlag          string
	clusterOverride   string
	noHeaderFlag	  bool
)

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.esc/config)")
	rootCmd.PersistentFlags().StringVarP(&nodeFlag, "nodes", "n", "", "set node ip:port")
	//rootCmd.PersistentFlags().StringVarP(&clusterOverride, "cluster", "c", "", "set a temporary current cluster")
	cobra.OnInitialize(onInit)
}

func onInit() {
	var err error
	cfg, err = config.ReadConfig(cfgFile)
	if err != nil {
		errorExit("Invalid config: %v", err)
	}

	cfg.ClusterOverride = clusterOverride

	cluster := cfg.ActiveCluster()
	if cluster != nil {
		// Use active cluster from config
		currentCluster = cluster
	} else {
		// Create sane default if not configured
		currentCluster = &config.Cluster{
			Nodes: "localhost:9092",
		}
	}
}

func getClient() *elasticsearch.Client {
	client, err := elasticsearch.NewClient(getElasticConfig())
	if err != nil {
		errorExit("Unable to get client: %v\n", err)
	}
	return client
}

func getClientFromConfig(config elasticsearch.Config) *elasticsearch.Client {
	client, err := elasticsearch.NewClient(getElasticConfig())
	if err != nil {
		errorExit("Unable to get client: %v\n", err)
	}
	return client
}

func errorExit(format string, a ...interface{}) {
	fmt.Fprintf(errWriter, format+"\n", a...)
	os.Exit(1)
}
