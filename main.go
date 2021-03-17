package main

import (
	"flag"
	"fmt"
	"github.com/mirzaakhena/danarisan/application"
	"github.com/mirzaakhena/danarisan/application/registry"
)

func main() {

	flag.Parse()
	switch flag.Arg(0) {

	case "danarisan":
		application.Run(registry.NewArisanSystemRegistry())

	case "danamock":
	  application.Run(registry.NewDANAMockRegistry())

	default:
		fmt.Printf("Try\ngo run main.go danarisan\nor\ngo run main.go danamock")
	}

}
