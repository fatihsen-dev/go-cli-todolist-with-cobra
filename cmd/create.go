package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fatihsen-dev/go-cli-todolist-with-cobra/stores"
	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "Create a todo",
	// Long:  `long desc`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Print("Please enter text")
			return
		}

		var text string

		for i := 0; i < len(args); i++ {
			text += args[i] + " "
		}

		newTodo := stores.TodoT{
			Id:     rand.Intn(99999-10000) + 1,
			Text:   text,
			Status: "pending",
			Date:   time.Now().Format("02-01-2006 15:04:05"),
		}

		stores.Todos = append(stores.Todos, newTodo)

		tableData = append(tableData, []string{strconv.Itoa(newTodo.Id), newTodo.Text, newTodo.Status, newTodo.Date})

		RenderTable()
		stores.SyncJson()
	},
}
