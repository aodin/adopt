package main

import (
	"flag"
	"fmt"
	"github.com/aodin/adopt/robot"
	"github.com/aodin/aspect"
)

var applications = map[string][]*aspect.TableElem{
	"pets": {
		robot.Pets,
	},
}

func main() {
	flag.Parse()
	for _, app := range flag.Args() {
		tables, ok := applications[app]
		if !ok {
			continue
		}
		for _, table := range tables {
			fmt.Println(table.Create())
		}
	}
}
