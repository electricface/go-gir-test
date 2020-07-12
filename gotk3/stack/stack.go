package main

import (
	"fmt"
	"github.com/electricface/go-gir3/gi"
	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
	"log"
)

const (
	winTitle = "stack example"
)

func main() {
	gtk.Init(0, 0)
	win := setupWindow(winTitle)
	box := newStackFull()
	win.Add(box)
	win.ShowAll()
	gtk.Main()
}

func setupWindow(title string) gtk.Window {
	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	win.SetTitle(title)
	win.Connect(gtk.SigDestroy, gtk.MainQuit)
	win.SetDefaultSize(800, 600)
	win.SetPosition(gtk.WindowPositionCenter)
	return win
}

func newStackFull() gtk.IWidget {
	stack := gtk.NewStack()
	sw := gtk.NewStackSwitcher()
	sw.SetStack(stack)

	boxText1 := newBoxText("Hello there!")
	boxRadio := newBoxRadio("choice 1", "choice 2", "choice 3", "choice 4")
	boxText2 := newBoxText("third page")

	stack.AddTitled(boxText1, "key1", "first page")
	stack.AddTitled(boxRadio, "key2", "second page")
	stack.AddTitled(boxText2, "key3", "third page")

	v, err := g.NewValueWith("list-add")
	if err != nil {
		log.Fatal(err)
	}
	stack.ChildSetProperty(boxRadio, "icon-name", v)
	v.Free()

	box := setupBox(gtk.OrientationVertical)
	box.PackStart(sw, false, false, 0)
	box.PackStart(stack, true, true, 0)
	return box
}

func newBoxText(content string) gtk.Widget {
	box := setupBox(gtk.OrientationVertical)
	tv := setupTView()
	setTextInTView(tv, content)
	box.PackStart(tv, true, true, 0)
	btn := setupBtn("Submit", func() {
		text := getTextFromTView(tv)
		fmt.Println(text)
	})
	box.Add(btn)
	return box.Widget
}

func setupBox(orient gtk.OrientationEnum) gtk.Box {
	box := gtk.NewBox(orient, 0)
	return box
}

func setupTView() gtk.TextView {
	tv := gtk.NewTextView()
	return tv
}

func setTextInTView(tv gtk.TextView, content string) {
	buffer := getBufferFromTView(tv)
	buffer.SetText(content, int32(len(content)))
}

func getBufferFromTView(tv gtk.TextView) gtk.TextBuffer {
	buffer := tv.GetBuffer()
	return buffer
}

func setupBtn(label string, onClick func()) gtk.Button {
	btn := gtk.NewButtonWithLabel(label)
	btn.Connect(gtk.SigClicked, onClick)
	return btn
}

func getTextFromTView(tv gtk.TextView) string {
	buffer := getBufferFromTView(tv)

	start := gtk.TextIter{P: gi.Malloc(gtk.SizeOfStructTextIter)}
	end := gtk.TextIter{P: gi.Malloc(gtk.SizeOfStructTextIter)}
	buffer.GetBounds(start, end)
	text := buffer.GetText(start, end, true)

	start.Free()
	end.Free()
	return text
}

func newBoxRadio(btns ...string) gtk.IWidget {
	var (
		prev gtk.RadioButton
		box  = setupBox(gtk.OrientationVertical)
	)

	for i, text := range btns {
		radio := gtk.NewRadioButtonWithLabelFromWidget(prev, text)
		box.PackStart(radio, false, false, 0)
		prev = radio
		i := i
		radio.Connect(gtk.SigToggled, func() {
			fmt.Println(i, radio.GetActive())
		})
	}
	return box
}
