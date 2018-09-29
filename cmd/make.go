package cmd

import (
	"io/ioutil"
	"os"
)

const goeServer = "server.go"
const goeBeacon = "goe.InitServer()"

const makeError001 = "Please enter the type and name to be make"
const makeError002 = "The current dir is not a goe project"
const makeError003 = "The first letter of the name cannot be a number"
const makeError004 = " is exist"

func mAke(args []string) {
	if len(args) < 2 {
		printInterrupt(makeError001)
	}
	if !fileExists(goeServer) || !fileContainsString(goeServer, goeBeacon) {
		printInterrupt(makeError002)
	}
	typeList := map[string]func(actionName string, routeName string) string{
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
				os.Mkdir(makeType, 0755)
			}
			var writePath = makeType + "/" + makeName + ".go"
			if !fileExists(writePath) {
				ioutil.WriteFile(writePath, []byte(f(upFirst(makeName), makeName)), 0755)
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
