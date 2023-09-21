package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatihsen-dev/go-cli-todolist-with-cobra/stores"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var tableData = [][]string{}

var lists = &cobra.Command{
	Use:   "list",
	Short: "List all todo",
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetString("status")

		switch status {
		case "pending":
			ListWithStatus(status)
		case "complated":
			ListWithStatus(status)
		case "deleted":
			ListWithStatus(status)
		default:
			ListAll()
		}
	},
}

func init() {
	lists.Flags().StringP("status", "s", "", "-s complated / -s pending / -s deleted")
}

func ListAll() {
	for _, todo := range stores.Todos {
		tableData = append(tableData, []string{strconv.Itoa(todo.Id), todo.Text, todo.Status, todo.Date})
	}

	if len(tableData) == 0 {
		fmt.Print("Todo not found create one")
		return
	}
	RenderTable()
}

func ListWithStatus(status string) {
	for _, todo := range stores.Todos {
		if todo.Status == status {
			tableData = append(tableData, []string{strconv.Itoa(todo.Id), todo.Text, todo.Status, todo.Date})
		}
	}

	if len(tableData) == 0 {
		fmt.Print("Todo not found create one")
		return
	}
	RenderTable()
}

func RenderTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "Todo", "Status", "Date"})

	for _, todo := range tableData {
		table.Append(todo)
	}

	table.Render()
}
