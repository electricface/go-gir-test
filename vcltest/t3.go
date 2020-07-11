package main

import (
	"github.com/ying32/govcl/vcl"
	"log"
)

func main() {
	vcl.Application.Initialize()
	namePath := vcl.Application.GetNamePath()
	log.Println(namePath)

	exeName := vcl.Application.ExeName()
	log.Println(exeName)
}
