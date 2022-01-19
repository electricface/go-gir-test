package main

import (
	"log"

	"github.com/electricface/go-gir/g-2.0"
	"github.com/electricface/go-gir/gi"
)

func main() {
	val, err := g.NewValue()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("val isValid:", val.IsValid())

	val, err = g.NewValueWith(123)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("val isValid:", val.IsValid())

	v := val.GetInt()
	log.Println("get:", v)
	val.Free()

	val, err = g.NewValueWith(uint(123))
	if err != nil {
		log.Fatalln(err)
	}
	vUint := val.GetUint()
	log.Println("get uint:", vUint)
	val.Free()

	val, err = g.NewValueWith(gi.Long(123))
	if err != nil {
		log.Fatalln(err)
	}
	vLong := val.GetLong()
	log.Println("get long:", vLong)
	val.Free()

	val, err = g.NewValueWith(gi.Ulong(123))
	if err != nil {
		log.Fatalln(err)
	}
	vUlong := val.GetUlong()
	log.Println("get ulong:", vUlong)
	val.Free()

	val, err = g.NewValueWith("hell world")
	if err != nil {
		log.Fatalln(err)
	}
	vStr := val.GetString()
	log.Println("get str:", vStr)
	val.Free()
}

func main1() {
	log.Println("new value")
	val, err := g.NewValue()
	if err != nil {
		log.Fatal(err)
	}

	val.Init(g.TYPE_INT)
	val.SetInt(110)

	valInt := val.GetInt()
	log.Println("int:", valInt)

	valUint := val.GetUint()
	valUint = val.GetUint()
	log.Println("uint:", valUint)

	val.Unset()
	val.Init(g.TYPE_STRING)
	val.SetString("hello world")
	//val.SetStaticString("hello world") // 不适合
	//val.SetStringTakeOwnership("hello world") // 不适合
	valStr := val.GetString()
	log.Println("valStr:", valStr)

	val.Unset()
	val.Init(g.TYPE_OBJECT)
	kf := g.NewKeyFile()

	log.Println("set object kf")
	val.SetObject(g.WrapObject(kf.P))

	gs := g.NewSettings("com.deepin.dde.dock")
	log.Println("set object gs")
	val.SetObject(gs)

	val.Unset()
	val.Init(g.SettingsGetType())
	log.Println("set instance gs.P")
	val.SetInstance(gs.P)

	//otherVal := g.NewValue()
	//otherVal.Init(g.ValueGetType())
	//log.Println("set instance otherVal")
	//val.SetInstance(gs.P)

	log.Println("free 1")
	val.Free()
	log.Println("free 2")
	val.Free()
	log.Println("end")

}
