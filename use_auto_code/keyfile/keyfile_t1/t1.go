package main

import (
	"log"

	"github.com/electricface/go-gir/g-2.0"
	gi "github.com/electricface/go-gir/gi"
)

func main() {
	ok, charset := g.GetCharset()
	log.Println(ok, charset)

	ok, charset = g.GetCharset()
	log.Println(ok, charset)

	//kf := g.NewKeyFile()
	//kf.Unref()

	//dai := g.NewDesktopAppInfoFromKeyfile(kf)
	//dai.Unref()
	//kf.Unref()

	dai := g.NewDesktopAppInfo("dde-control-center.desktop")
	filename := dai.GetFilename()

	kf := g.NewKeyFile()
	_, err := kf.LoadFromFile(filename, 0)
	if err != nil {
		log.Fatal(err)
	}
	result, err := kf.GetLocaleString("Desktop Entry", "Name", gi.NilStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("name:", result)
	//dai.Unref()

	//gi.Free(nil)
	//arr := gi.NewCStrArrayZTWithStrings("hello","world")
	//arrCp := arr.Copy()
	//log.Println(arrCp)
	//arr.FreeAll()
	//
	//
	//arr = gi.NewCStrArrayWithStrings("hello","world", gi.NilStr)
	//arrCp = arr.Copy()
	//log.Printf("%#v\n", arrCp)
	//arr.FreeAll()
	//dai.GetLocaleString()

	kf = g.NewKeyFile()
	searchDirs := gi.NewCStrArrayZTWithStrings("/usr/local/share/applications", "/usr/share/applications")
	_, fullPath, err := kf.LoadFromDirs("dde-control-center.desktop", searchDirs, 0)
	defer searchDirs.FreeAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("fullPath:", fullPath)
}
