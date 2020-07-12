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
	typ   string
	scope string
	title string
	desc  string
	log   string

	updateFn func(string)
}

func (s *State) handleAction(act interface{}) {
	switch a := act.(type) {
	case *actionUpdate:
		value := strings.TrimSpace(a.value)
		switch a.prop {
		case "type":
			s.typ = value
		case "scope":
			s.scope = value
		case "title":
			s.title = value
		case "desc":
			s.desc = value
		case "log":
			s.log = value
		}
		s.updatePreview()
	}
}

func (s *State) updatePreview() {
	var sb strings.Builder
	module := ""
	if s.scope != "" {
		module = "(" + s.scope + ")"
	}
	sb.WriteString(fmt.Sprintf("%s%s: %s\n", s.typ, module, s.title))
	sb.WriteString("\n")
	sb.WriteString(s.desc + "\n")
	sb.WriteString("\n")

	if s.log != "" {
		sb.WriteString("Log: " + s.log + "\n")
	}

	text := sb.String()
	s.updateFn(text)
}

type actionUpdate struct {
	prop, value string
}

func buildUI() gtk.Window {
	state := &State{}
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

	entryScope := gtk.WrapEntry(builder.GetObject("entryScope").P)
	entryScope.Connect(gtk.SigChanged, func() {
		log.Println("entryScope changed")
		state.handleAction(&actionUpdate{
			prop:  "scope",
			value: entryScope.GetText(),
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
	state.typ = "fix"
	state.updatePreview()

	win.Connect(gtk.SigDestroy, gtk.MainQuit)
	return win
}

func main() {
	gtk.Init(0, 0)
	win := buildUI()
	win.ShowAll()
	gtk.Main()
}
