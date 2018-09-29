package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"unicode"
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

func gitClone(repo string, name string) {
	exec.Command("git", []string{"clone", repo, name}...).Run()
}

func upFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func cleanAllSymbol(str string) string {
	result := ""
	arr := regexp.MustCompile("[A-Za-z0-9]*").FindAllString(str, len(str))
	for _, val := range arr {
		result += val
	}
	return result
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func readFile(fileName string) string {
	buf, err := ioutil.ReadFile(fileName)
	if os.IsNotExist(err) {
		return ""
	}
	return string(buf)
}

func fileContainsString(fileName string, content string) bool {
	lines := strings.Split(readFile(fileName), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == content {
			return true
		}
	}
	return false
}
