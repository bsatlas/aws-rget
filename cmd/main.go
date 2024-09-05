package main

import (
	"fmt"
	"os"

	"github.com/bsatlas/aws-rget/cmd/version"

	"github.com/spf13/cobra"
)

func command() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "caliper",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	subCommands := []func() *cobra.Command{
		version.Command,
	}
	for _, subCommandFn := range subCommands {
		subCommand := subCommandFn()
		cmd.AddCommand(subCommand)
	}
	return cmd
}

func execute() {
	err := command().Execute()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func main() {
	execute()
}
