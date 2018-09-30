package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const (
	goeServer = "server.go"
	goeBeacon = "goe.InitServer()"

	makeError001 = "Please enter the type and name to be make"
	makeError002 = "The current dir is not a goe project"
	makeError003 = "The first letter of the name cannot be a number"
	makeError004 = " is exist"
)

func mAke(args []string) {
	if len(args) < 2 {
		printInterrupt(makeError001)
	}
	if !fileExists(goeServer) || !fileContainsString(goeServer, goeBeacon) {
		printInterrupt(makeError002)
	}
	typeList := map[string]func(string, string) string{
		"action": actionTemplate,
		"listen": listenTemplate,
	}
	makeType := args[0]
	makeName := cleanAllSymbol(args[1])
	if makeName[0] > 47 && makeName[0] < 58 {
		printInterrupt(makeError003)
	}
	for t, f := range typeList {
		if t == makeType {
			if !fileExists(makeType) {
				os.Mkdir(makeType, filePermission)
			}
			var writePath = makeType + "/" + makeName + ".go"
			if !fileExists(writePath) {
				// Make items
				ioutil.WriteFile(writePath, []byte(
					f(upFirst(makeName),
						makeName)), filePermission)

				// Register make items
				content := strings.Replace(
					readFile(goeServer),
					goeBeacon,
					registerTemplate(
						makeType,
						makeName)+"\n\t"+goeBeacon, -1)

				// Check import
				importStr := regexp.MustCompile(`import ([^)]+)`).FindString(content)
				importTypeStr := "\t" + `"./` + makeType + `"`

				if !strContainsString(importStr, importTypeStr) {
					content = strings.Replace(content, importStr, importStr+importTypeStr+"\n", -1)
				}
				ioutil.WriteFile(goeServer, []byte(content), filePermission)
			} else {
				printInterrupt(writePath + makeError004)
			}
		}
	}
}

func actionTemplate(actionName string, routeName string) string {
	return `package action

import (
	"github.com/ezcorn/goe"
	"net/http"
)

func ` + actionName + `Action() *goe.Action {
	return goe.NewAction("/` + routeName + `", "` + routeName + ` Action", []string{
		http.MethodPost, http.MethodGet,
	}, func(in goe.In, out goe.Out) {
	})
}
`
}

func listenTemplate(actionName string, routeName string) string {
	return `package listen

import (
	"github.com/ezcorn/goe"
)

func ` + actionName + `Listen() *goe.Listen {
	return goe.NewListen("` + routeName + `", "` + routeName + ` Listen", func(in goe.In) goe.Program {
		return nil
	})
}
`
}

func registerTemplate(regType string, routeName string) string {
	upFirstType := upFirst(regType)
	upFirstRoute := upFirst(routeName)
	result := `// Register ` + upFirstType + " " + upFirstRoute + "\n\tgoe.Reg" +
		upFirstType + "(" + regType + "." + upFirstRoute + upFirstType + ")\n"
	return result
}
