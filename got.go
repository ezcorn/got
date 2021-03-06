package main

import (
	"github.com/ezcorn/got/cmd"
	"os"
)

func main() {
	cmd.MakeCmdRegistry()
	if len(os.Args) == 1 {
		os.Args = append(os.Args, cmd.CommandHelp)
	}
	cmd.Exec()
}
