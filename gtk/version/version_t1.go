package main

import (
	"log"

	"github.com/electricface/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	v1 := gtk.GetMajorVersion()
	v2 := gtk.GetMinorVersion()
	v3 := gtk.GetMicroVersion()
	log.Println(v1, v2, v3)
}
