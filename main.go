package main

import (
	"flag"
	"github.com/mirzaakhena/danarisan/application"
	"github.com/mirzaakhena/danarisan/application/registry"
)

func main() {

	flag.Parse()
	switch flag.Arg(0) {

	default:
		application.Run(registry.NewArisanSystemRegistry())
	}

}
