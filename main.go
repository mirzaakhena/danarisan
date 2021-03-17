package main

import (
	"flag"
	"github.com/mirzaakhena/danarisan/application"
	"github.com/mirzaakhena/danarisan/application/registry"
)

func main() {

	flag.Parse()
	switch flag.Arg(0) {

	// case "firstService":
	//   application.Run(registry.NewOrderServiceRegistry())

	// case "secondService":
	//   application.Run(registry.NewOrderServiceRegistry())

	default:
		application.Run(registry.NewDefaultRegistry())
	}

}
