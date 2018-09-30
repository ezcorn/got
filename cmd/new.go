package cmd

import "os"

func neW(args []string) {
	if len(args) == 0 {
		printInterrupt("Please enter the project name")
	}
	projectName := args[0]
	if fileExists(projectName) {
		printInterrupt(`Folder "` + projectName + `" already exists`)
	}
	gitClone("https://github.com/ezcorn/goe-base.git", projectName)
	os.RemoveAll(projectName + "/.git")
	printInterrupt(`New project [ ` + projectName + ` ] building complete`)
}
