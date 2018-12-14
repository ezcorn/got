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
	CommandHelp   = "help"   // 帮助文档	got help
	CommandNew    = "new"    // 新建服务 got new < projectName >
	CommandMake   = "make"   // 新建模块 got make:< moduleName > < paramString >
	CommandDel    = "del"    // 删除模块
	CommandJoin   = "join"   // 把一个server加入生态
	CommandUnion  = "union"  // 联合两个生态
	CommandUpdate = "update" // 更新本体

	// join:
	//		申请一个port
	//		通过这个port生成一个server
	//		如果这个server的name已经在生态中存在,则直接派生
	//		如果这个server的name不存在生态中,则初始化一个新种类的生命形式
	//		获取/apis/supplier来构建需要join的服务
	//		构建完毕后,执行/apis/register

	filePermission = 0755
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
	cmdRegistry[CommandHelp] = &cmd{exec: help, note: CommandHelp}
	cmdRegistry[CommandNew] = &cmd{exec: neW, note: CommandNew}
	cmdRegistry[CommandDel] = &cmd{exec: del, note: CommandDel}
	cmdRegistry[CommandMake] = &cmd{exec: mAke, note: CommandMake}
	cmdRegistry[CommandUpdate] = &cmd{exec: update, note: CommandUpdate}
}

func printInterrupt(content string) {
	fmt.Println(content)
	os.Exit(0)
}

func gitClone(repo string, name string) {
	err := exec.Command("git", []string{"clone", repo, name}...).Run()
	if err != nil {
		printInterrupt(err.Error())
	}
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

func strContainsString(str string, c string) bool {
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == strings.TrimSpace(c) {
			return true
		}
	}
	return false
}

func fileContainsString(fileName string, c string) bool {
	return strContainsString(readFile(fileName), c)
}
