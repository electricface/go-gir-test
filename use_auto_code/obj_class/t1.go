package main

import (
	"log"

	"github.com/electricface/go-gir3/gi"
)

func main() {
	main2()
}

func main2() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Gtk", "3.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		log.Fatal(err)
	}
	bi := repo.FindByName("Gtk", "WidgetClass")

	objType := bi.Type()
	switch objType {

	case gi.INFO_TYPE_OBJECT:
		objInfo := gi.ToObjectInfo(bi)
		methodInfo := objInfo.FindMethod("get_css_name")
		if methodInfo != nil {
			log.Println("b1 found it")
		} else {

			log.Println("b1 not found")
		}

	case gi.INFO_TYPE_STRUCT:
		structInfo := gi.ToStructInfo(bi)
		numMethods := structInfo.NumMethod()
		for i := 0; i < numMethods; i++ {
			method := structInfo.Method(i)
			log.Println(i, method.Name())
		}
		// struct GObject.ObjectClass panic
		// struct Gtk.WidgetClass not found
		methodInfo := structInfo.FindMethod("get_css_name")
		if methodInfo != nil {
			log.Println("b2 found it")
		} else {
			log.Println("b2 not found")
		}
		structInfo.IterateAttributes(func(name, value string) {
			log.Printf("name: %v, value: %v\n", name, value)
		})

	}

	return
}

func main1() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GObject", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		log.Fatal(err)
	}
	bi := repo.FindByName("GObject", "ObjectClass")

	objType := bi.Type()
	switch objType {

	case gi.INFO_TYPE_OBJECT:
		objInfo := gi.ToObjectInfo(bi)
		methodInfo := objInfo.FindMethod("list_properties")
		if methodInfo != nil {
			log.Println("b1 found it")
		} else {

			log.Println("b1 not found")
		}

	case gi.INFO_TYPE_STRUCT:
		structInfo := gi.ToStructInfo(bi)
		numMethods := structInfo.NumMethod()
		for i := 0; i < numMethods; i++ {
			method := structInfo.Method(i)
			log.Println(i, method.Name())
		}
		// struct GObject.ObjectClass panic
		// struct Gtk.WidgetClass not found
		methodInfo := structInfo.FindMethod("list_properties")
		if methodInfo != nil {
			log.Println("b2 found it")
		} else {
			log.Println("b2 not found")
		}
		structInfo.IterateAttributes(func(name, value string) {
			log.Printf("name: %v, value: %v\n", name, value)
		})

	}

	return
}
