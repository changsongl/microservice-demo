package cmd

import (
	"github.com/changsongl/microservice/vars"
	"github.com/spf13/cobra"
)

var kcliCmd = &cobra.Command{
	Use:   "kcli",
	Short: "Kcli is a very fast static microservice generator.",
	Long: `Kcli is a project generator which is used to create microservice golang project`,
	Version: vars.KcliVersion,
}

// Execute executes the root command.
func Execute() error {
	return kcliCmd.Execute()
}