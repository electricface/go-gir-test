package main

import (
	"log"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/electricface/go-gir3/gi"
)

func main() {
	vari := g.NewVariantString("helloworld")
	type0 := vari.GetType()
	dupStr := type0.DupString()
	str := gi.GoString(type0.P)
	log.Println(str)
	log.Println(dupStr)
}
