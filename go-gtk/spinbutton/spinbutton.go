package main

import (
	"fmt"
	"strconv"

	"github.com/electricface/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetPosition(gtk.WindowPositionCenter)
	window.SetTitle("GTK Go!")

	window.Connect(gtk.SigDestroy, func() {
		fmt.Println("got destroy!")
		gtk.MainQuit()
	})

	//--------------------------------------------------------
	// GtkHBox
	//--------------------------------------------------------
	fixed := gtk.NewFixed()

	//--------------------------------------------------------
	// GtkSpinButton
	//--------------------------------------------------------
	spinbutton1 := gtk.NewSpinButtonWithRange(1.0, 10.0, 1.0)
	spinbutton1.SetDigits(3)
	spinbutton1.Spin(gtk.SpinTypeStepForward, 7.0)
	fixed.Put(spinbutton1, 40, 50)

	spinbutton1.Connect(gtk.SigValueChanged, func() {
		val := int(spinbutton1.GetValueAsInt())
		fval := spinbutton1.GetValue()
		fmt.Println("SpinButton changed, new value: " + strconv.Itoa(val) + " | " + strconv.FormatFloat(fval, 'f', 2, 64))
		min, max := spinbutton1.GetRange()
		fmt.Println("Range: " + strconv.FormatFloat(min, 'f', 2, 64) + " " + strconv.FormatFloat(max, 'f', 2, 64))
		fmt.Println("Digits: " + strconv.Itoa(int(spinbutton1.GetDigits())))
	})

	adjustment := gtk.NewAdjustment(2.0, 1.0, 8.0, 2.0, 0.0, 0.0)
	spinbutton2 := gtk.NewSpinButton(adjustment, 1.0, 1)
	spinbutton2.SetRange(0.0, 20.0)
	spinbutton2.SetValue(18.0)
	spinbutton2.SetIncrements(2.0, 4.0)
	fixed.Put(spinbutton2, 200, 50)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(fixed)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
