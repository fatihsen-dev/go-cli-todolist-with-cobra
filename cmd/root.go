package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:   "[command]",
	Short: "Todo List with go cobra cli fw",
	Long:  `Github: www.github.com/fatihsen-dev`,
}

func Execute() {
	err := Root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	Root.AddCommand(create)
	Root.AddCommand(update)
	Root.AddCommand(delete)
	Root.AddCommand(lists)
}
