package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/atomicals-go/lib/atomicals"
	"github.com/atomicals-go/lib/cli"
	"github.com/atomicals-go/lib/plugin"
)

var rootCmd = &cobra.Command{
	Use:   "atomicals-cli",
	Short: "atomical CLI for golang",
	Long:  `atomical CLI for golang`,
}

func init() {
	// log setting
	logrus.SetOutput(os.Stdout)

	// config setting
	atomicals.Config = plugin.InitViperConfig()

	// cmd setting
	rootCmd.SetVersionTemplate("v0.0.1")
	rootCmd.AddCommand(cli.MintDft)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
