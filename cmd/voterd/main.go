package main

import (
	"os"

	"github.com/borbier/voter/cmd/voterd/cmd"
    svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/borbier/voter/app"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
    if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
