package cmd

func neW(args []string) {
	if len(args) == 0 {
		printInterrupt("Please enter the project name")
	}
	projectName := args[0]
	if pathExists(projectName) {
		printInterrupt(`Folder "` + projectName + `" already exists`)
	}
	gitClone("https://github.com/ezcorn/goe-example.git", projectName)
	printInterrupt(`New project [ ` + projectName + ` ] build complete`)
}
