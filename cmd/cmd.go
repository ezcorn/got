package cmd

import "os"

const (
	HELP = "help"
	NEW  = "new"
	MAKE = "make"
)

type (
	Cmd struct {
		word string
		exec func(args []string)
	}
)

var (
	cmdRegistry = make(map[string]*Cmd)
)

func reg(word string, exec func(args []string)) {
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

func MakeCmdRegistry() {
	reg(HELP, help)
	reg(NEW, nw)
	reg(MAKE, mk)
}

func Exec() {
	if cmd, ok := cmdRegistry[os.Args[1]]; ok {
		cmd.exec(os.Args[1:])
	}
}
