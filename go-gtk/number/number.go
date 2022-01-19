package main

import (
	"fmt"
	"log"
	"strconv"
	"unsafe"

	"github.com/electricface/go-gir/g-2.0"
	"github.com/electricface/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)

	dialog := gtk.NewDialog()
	dialog.SetTitle("number input")

	vbox := dialog.GetContentArea()

	label := gtk.NewLabel("Numnber:")
	vbox.Add(label)

	input := gtk.NewEntry()
	input.SetEditable(true)
	vbox.Add(input)

	input.Connect(gtk.SigInsertText, func(args []interface{}) {
		target := args[0].(g.Object)
		s := args[1].(string)
		length := args[2].(int32)
		log.Println(target, s, length)
		//arr := gi.Uint8Array{P: arrP, Len: length}
		position := (*int)(args[3].(unsafe.Pointer))
		log.Printf("positon: %p, val is %v\n", position, *position)

		if s == "." {
			if *position == 0 {
				log.Println("stop emission")
				g.SignalStopEmissionByName(target, "insert-text")
				//input.StopEmission("insert-text")
			}
		} else {
			_, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Println("stop emission")
				g.SignalStopEmissionByName(target, "insert-text")
			}
		}
	})

	button := gtk.NewButtonWithLabel("OK")
	button.Connect(gtk.SigClicked, func() {
		fmt.Println(input.GetText())
		gtk.MainQuit()
	})
	vbox.Add(button)

	dialog.ShowAll()
	gtk.Main()
}
