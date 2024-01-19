package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/atomicals-go/lib/cli"
)

var rootCmd = &cobra.Command{
	Use:   "atomicals-cli",
	Short: "atomical CLI for golang",
	Long:  `atomical CLI for golang`,
}

func init() {
	rootCmd.SetVersionTemplate("v0.0.1")
	rootCmd.AddCommand(cli.MintDft)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
