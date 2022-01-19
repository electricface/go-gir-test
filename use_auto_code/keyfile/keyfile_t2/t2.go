package main

import (
	"log"

	"github.com/electricface/go-gir/g-2.0"
)

func main() {
	kf := g.NewKeyFile()
	ok, err := kf.LoadFromFile("/usr/share/applications/dde-control-center.desktop", 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ok)
	name, err := kf.GetString("Desktop Entry", "Name")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(name)

	name, err = kf.GetLocaleString("Desktop Entry", "Name", "zh_CN.UTF-8")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(name)
}
