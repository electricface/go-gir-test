package main

import (
	"log"
	"unsafe"

	"github.com/davecgh/go-spew/spew"

	"github.com/linuxdeepin/go-gir/g-2.0"
)

func main() {
	fileInfo := g.NewFileInfo()
	objClass := fileInfo.GetClass()
	propsArr := objClass.ListProperties()
	spew.Dump(propsArr)
	propsArr.Free()
}

func main1() {
	gType := g.FileInfoGetType()

	typeClass := g.TypeClassRef(gType)
	objClass := g.ObjectClass{P: typeClass.P}

	propsArr := objClass.ListProperties()
	propsSlice := propsArr.AsSlice()
	spew.Dump(propsArr)
	arr := *(*[]g.ParamSpec)(unsafe.Pointer(&propsSlice))
	for i, spec := range arr {
		name := spec.GetName()
		log.Println(i, name)
	}
	propsArr.Free()

	isMaxPs := objClass.FindProperty("is-maximized")
	if isMaxPs.P != nil {
		name := isMaxPs.GetName()
		log.Println("isMaxPs name:", name)
	} else {
		log.Println("not found prop is-maximized")
	}
}
