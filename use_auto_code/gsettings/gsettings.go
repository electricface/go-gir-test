package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	gs := g.NewSettings("ca.desrt.dconf-editor.Demo")
	gs.Connect(gtk.SigChanged, func(args []interface{}) {
		spew.Dump(args)
		key := args[1].(string)
		log.Println("key changed", key)

		gs0 := g.WrapSettings(args[0].(g.Object).P)
		val := gs0.GetValue(key)
		str := val.Print(true)
		log.Println(str)
	})
	val, err := g.NewValueT(g.TYPE_STRING)
	if err != nil {
		log.Fatal(err)
	}
	gs.GetProperty("path", val)
	path := val.GetString()

	gs.GetProperty("schema", val)
	schema := val.GetString()

	gs.GetProperty("schema-id", val)
	schemaId := val.GetString()

	val.Free()
	log.Println("path:", path)
	log.Println("schema:", schema)
	log.Println("schema-id:", schemaId)

	valObj, err := g.NewValueT(g.SettingsSchemaGetType())
	if err != nil {
		log.Fatal(err)
	}
	gs.GetProperty("settings-schema", valObj)
	ss := g.SettingsSchema{P: valObj.GetBoxed()}
	keysArr := ss.ListKeys()
	keys := keysArr.Copy()
	log.Println(keys)
	keysArr.FreeAll()

	hasKey := ss.HasKey("hello-world")
	log.Println("hasKey:", hasKey)

	//variant := gs.GetDefaultValue("boolean")
	userVal := gs.GetUserValue("boolean")
	if userVal.P == nil {
		log.Println("do not set user value")
	} else {
		b := userVal.GetBoolean()
		log.Println("already set user value", b)
		userVal.Unref()
	}

	gtk.Main()
}
