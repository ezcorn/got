package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	HELP   = "help"
	NEW    = "new"
	MAKE   = "make"
	UPDATE = "update"
)

type (
	cmd struct {
		note string
		exec func(args []string)
	}
)

var (
	cmdRegistry = make(map[string]*cmd)
)

func Exec() {
	words := strings.Split(os.Args[1], ":")
	if cmd, ok := cmdRegistry[words[0]]; ok {
		cmd.exec(append(words[1:], os.Args[2:]...))
	}
}

func MakeCmdRegistry() {
	cmdRegistry[HELP] = &cmd{exec: help, note: HELP}
	cmdRegistry[NEW] = &cmd{exec: neW, note: NEW}
	cmdRegistry[MAKE] = &cmd{exec: mAke, note: MAKE}
	cmdRegistry[UPDATE] = &cmd{exec: update, note: UPDATE}
}

func printInterrupt(content string) {
	fmt.Println(content)
	os.Exit(0)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func gitClone(repo string, name string) {
	exec.Command("git", []string{"clone", repo, name}...).Run()
}
