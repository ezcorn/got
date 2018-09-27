package cmd

import (
	"os"
	"strings"
)

const (
	HELP   = "help"
	NEW    = "new"
	MAKE   = "make"
	UPDATE = "update"
)

var (
	cmdRegistry = make(map[string]func(args []string))
)

func Exec() {
	words := strings.Split(os.Args[1], ":")
	if cmd, ok := cmdRegistry[words[0]]; ok {
		cmd(append(words[1:], os.Args[2:]...))
	}
}

func MakeCmdRegistry() {
	cmdRegistry[HELP] = help
	cmdRegistry[NEW] = neW
	cmdRegistry[MAKE] = mAke
	cmdRegistry[UPDATE] = update
}
