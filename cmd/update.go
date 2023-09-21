package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatihsen-dev/go-cli-todolist-with-cobra/stores"
	"github.com/spf13/cobra"
)

var update = &cobra.Command{
	Use:   "update",
	Short: "update todo",
	Long:  `update -id 12 -new complated`,
	Run: func(cmd *cobra.Command, args []string) {
		stringId, _ := cmd.Flags().GetString("id")
		id, _ := strconv.Atoi(stringId)
		new, _ := cmd.Flags().GetString("new")

		switch new {
		case "pending":

		case "complated":

		case "deleted":

		default:
			fmt.Print("Invalid new status name")
			return
		}

		var todo stores.TodoT
		var todos []stores.TodoT
		status := false

		for _, td := range stores.Todos {
			if td.Id == id {
				td.Status = new
				todo = td
				status = true
			}
			todos = append(todos, td)
		}

		stores.Todos = todos

		if status {
			tableData = append(tableData, []string{strconv.Itoa(todo.Id), todo.Text, todo.Status, todo.Date})
			stores.SyncJson()
			RenderTable()
			return
		}

		fmt.Print("Todo not found")
		return
	},
}

func init() {
	update.Flags().StringP("id", "i", "", "update -id 12 -status pending -new complated")
	update.Flags().StringP("new", "n", "", "update -id 12 -status pending -new complated")

	update.MarkFlagRequired("id")
	update.MarkFlagRequired("new")
}
