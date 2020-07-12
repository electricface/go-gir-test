package main

import (
	"fmt"
	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gi"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
	"log"
	"strings"
)

type State struct {
	typ    string
	module string
	title  string
	desc   string
	log    string

	updateFn func(string)
}

func (s *State) handleAction(act interface{}) {
	switch a := act.(type) {
	case *actionUpdate:
		switch a.prop {
		case "type":
			s.typ = a.value
		case "module":
			s.module = a.value
		case "title":
			s.title = a.value
		case "desc":
			s.desc = a.value
		case "log":
			s.log = a.value
		}
		s.updatePreview()
	}
}

func (s *State) updatePreview() {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s(%s): %s\n", s.typ, s.module, s.title))
	sb.WriteString("\n")
	sb.WriteString(s.desc + "\n")
	sb.WriteString("\n")
	sb.WriteString("Log: " + s.log + "\n")

	text := sb.String()
	s.updateFn(text)
}

type actionUpdate struct {
	prop, value string
}

func buildUI() gtk.Window {
	state := &State{}
	//win := gtk.NewWindow(gtk.WindowTypeToplevel)
	builder := gtk.NewBuilder()
	_, err := builder.AddFromFile("./ui.glade")
	if err != nil {
		log.Fatal(err)
	}
	win := gtk.WrapWindow(builder.GetObject("window").P)

	comboType := gtk.WrapComboBox(builder.GetObject("comboType").P)
	comboType.Connect(gtk.SigChanged, func() {
		log.Println("comboType changed")
		model := comboType.GetModel()
		treeIter := gtk.TreeIter{P: gi.Malloc(gtk.SizeOfStructTreeIter)}
		comboType.GetActiveIter(treeIter)
		v := g.NewValue()
		model.GetValue(treeIter, 0, v)
		txt := v.GetString()
		state.handleAction(&actionUpdate{
			prop:  "type",
			value: txt,
		})
	})

	entryModule := gtk.WrapEntry(builder.GetObject("entryModule").P)
	entryModule.Connect(gtk.SigChanged, func() {
		log.Println("entryModule changed")
		state.handleAction(&actionUpdate{
			prop:  "module",
			value: entryModule.GetText(),
		})
	})

	entryTitle := gtk.WrapEntry(builder.GetObject("entryTitle").P)
	entryTitle.Connect(gtk.SigChanged, func() {
		log.Println("entryTitle changed")
		state.handleAction(&actionUpdate{
			prop:  "title",
			value: entryTitle.GetText(),
		})
	})

	tvDesc := gtk.WrapTextView(builder.GetObject("tvDesc").P)
	bufDesc := tvDesc.GetBuffer()
	bufDesc.Connect(gtk.SigChanged, func() {
		log.Println("bufDesc changed")
		start := gtk.TextIter{P: gi.Malloc(gtk.SizeOfStructTextIter)}
		end := gtk.TextIter{P: gi.Malloc(gtk.SizeOfStructTextIter)}
		bufDesc.GetBounds(start, end)
		txt := bufDesc.GetText(start, end, false)
		state.handleAction(&actionUpdate{
			prop:  "desc",
			value: txt,
		})
	})

	tvLog := gtk.WrapTextView(builder.GetObject("tvLog").P)
	bufLog := tvLog.GetBuffer()
	bufLog.Connect(gtk.SigChanged, func() {
		log.Println("bufLog changed")
		start := gtk.TextIter{P: gi.Malloc(gtk.SizeOfStructTextIter)}
		end := gtk.TextIter{P: gi.Malloc(gtk.SizeOfStructTextIter)}
		bufLog.GetBounds(start, end)
		txt := bufLog.GetText(start, end, false)
		state.handleAction(&actionUpdate{
			prop:  "log",
			value: txt,
		})
	})

	tvPreview := gtk.WrapTextView(builder.GetObject("tvPreview").P)
	bufPreview := tvPreview.GetBuffer()
	state.updateFn = func(s string) {
		bufPreview.SetText(s, int32(len(s)))
	}

	win.Connect(gtk.SigDestroy, gtk.MainQuit)
	return win
}

func main() {
	gtk.Init(0, 0)
	win := buildUI()
	win.ShowAll()
	gtk.Main()
}
