package main

import (
	"log"

	"github.com/electricface/go-gir/g-2.0"
)

func main() {
	dai := g.NewDesktopAppInfo("deepin-editor.desktop")
	//dai.RefSink()
	if dai.P == nil {
		log.Fatal("ptr is nil")
	}
	name := dai.GetName()
	log.Println("name:", name)

	gName := dai.GetGenericName()
	log.Println("gName:", gName)

	mimeTypes := dai.GetSupportedTypes()
	types := mimeTypes.Copy()
	log.Printf("types: %#v\n", types)

	dai.Unref()
}
