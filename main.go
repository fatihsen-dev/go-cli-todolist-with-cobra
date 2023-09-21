package main

import (
	"github.com/fatihsen-dev/go-cli-todolist-with-cobra/cmd"
	"github.com/fatihsen-dev/go-cli-todolist-with-cobra/stores"
)

func main() {
	stores.FetchTodos()
	cmd.Execute()
}
