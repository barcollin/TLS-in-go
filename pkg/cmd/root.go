package cmd

import (
	"fmt"
	"os"

	"github.com/barcollin/TLS-in-go/pkg/cert"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Config struct {
	CACert *cert.CACert          `yaml:"caCert"`
	Cert   map[string]*cert.Cert `yaml:"certs"`
}

var cfgFilePath string
var config Config

var rootCmd = &cobra.Command{
	Use:   "tls",
	Short: "tls is a command line tool for TLS",
	Long: `A Fast and Flexible Static Site Generator built 
				with love by spf13 and friends in Go.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFilePath, "config", "", "config file (default is tls.yaml")
}

func initConfig() {
	if cfgFilePath == "" {
		cfgFilePath = "tls.yaml"
	}

	cfgFilebytes, err := os.ReadFile(cfgFilePath)
	if err != nil {
		fmt.Printf("Error while reading config file: %s\n ", err)
		return
	}
	err = yaml.Unmarshal(cfgFilebytes, &config)
	if err != nil {
		fmt.Printf("Error while reading config file: %s\n ", err)
		return
	}

	fmt.Printf("config file parsed: %+v\n", config)
}
