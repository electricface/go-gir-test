package main

import (
	"log"

	"pkg.deepin.io/lib/gettext"

	"github.com/davecgh/go-spew/spew"
	"github.com/electricface/go-gir/g-2.0"
)

func main() {
	gettext.SetLocale(gettext.LcAll, "")
	str := "こんにちわ世界\n"
	arr, bytesRead, err := g.LocaleFromUtf8(str, -1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("bytes read:", bytesRead)
	spew.Dump(arr)
	spew.Dump(arr.AsSlice())
	log.Println(string(arr.AsSlice()))
}
