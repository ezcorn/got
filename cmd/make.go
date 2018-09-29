package cmd

func mAke(args []string) {
	if len(args) < 2 {
		printInterrupt("Please enter the type and name to be make")
	}
	typeList := map[string]func(actionName string, routeName string) string{
		"action": makeAction,
		"listen": makeListen,
	}
	makeType := args[0]
	makeName := cleanAllSymbol(args[1])
	if makeName[0] > 47 && makeName[0] < 58 {
		printInterrupt("The first letter of the name cannot be a number")
	}
	for t, f := range typeList {
		if t == makeType {
			printInterrupt(f(upFirst(makeName), makeName))
		}
	}
}

func makeAction(actionName string, routeName string) string {
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

func makeListen(actionName string, routeName string) string {
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
