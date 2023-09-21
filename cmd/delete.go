package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatihsen-dev/go-cli-todolist-with-cobra/stores"
	"github.com/spf13/cobra"
)

var delete = &cobra.Command{
	Use:   "delete",
	Short: "delete todo",
	Long:  `delete -id 12`,
	Run: func(cmd *cobra.Command, args []string) {
		stringId, _ := cmd.Flags().GetString("id")
		id, _ := strconv.Atoi(stringId)

		var todo stores.TodoT
		var todos []stores.TodoT
		status := false

		for _, td := range stores.Todos {
			if td.Id != id {
				todos = append(todos, td)
			} else {
				todo = td
				status = true
			}
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
	delete.Flags().StringP("id", "i", "", "delete -id 12")
	delete.MarkFlagRequired("id")
}
