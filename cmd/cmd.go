package cmd

import (
	"os"
)

const HELP = "help"

var cmdRegistry = make(map[string]*Cmd)

type Cmd struct {
	word string
	exec func(args []string)
}

func MakeCmdRegistry() {
	regCmd(HELP, help)
}

func Exec(word string) {
	if cmd, ok := cmdRegistry[word]; ok {
		cmd.exec(os.Args[1:])
	}
}

func regCmd(word string, exec func(args []string)) {
	cmdRegistry[word] = &Cmd{
		word: word,
		exec: func(args []string) {
			if args == nil || len(args) == 0 || args[0] != word {
				return
			}
			exec(args[1:])
		},
	}
}
