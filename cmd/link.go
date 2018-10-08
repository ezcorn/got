package cmd

import (
	"io/ioutil"
	"strings"
)

const (
	linkError001 = ""
)

func link(args []string) {
	if len(args) < 2 {
		printInterrupt(linkError001)
	}
	checkIsGoeProject()

	ioutil.WriteFile(goeServer, []byte(strings.Replace(
		readFile(goeServer), goeBeacon,
		relateTemplate(args[1], args[0])+"\n\t"+goeBeacon, -1)), filePermission)
}

func relateTemplate(routeName string, listenName string) string {
	return `// Relate ` +
		routeName + " to " +
		listenName + "\n\t" + `goe.RelateActionToListen("` +
		routeName + `", "` +
		listenName + `")` + "\n"
}
