package main

import (
  "flag"
  "github.com/mirzaakhena/danarisan/application"
  "github.com/mirzaakhena/danarisan/application/registry"
  "github.com/mirzaakhena/danarisan/application/registry2"
)

func main() {

  flag.Parse()
  switch flag.Arg(0) {
  case "aplikasi1":
    application.Run(registry.NewArisanSystemRegistry())

  case "cli":
    application.RunApp(registry2.NewMyApplicationOne())

  default:
    application.Run(registry.NewArisanSystemRegistry())
  }

}
