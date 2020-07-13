package main

import (
	"fmt"
	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gi"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
	"log"
	"net/url"
	"os/exec"
	"strconv"
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

	winAddItem := gtk.WrapWindow(builder.GetObject("winAddItem").P)
	buildUIWinAddItem(winAddItem, builder)
	winAddItem.ShowAll()
	return win
}

type subState struct {
	activeRb string
	url      string
	issue    string

	onActiveRbChanged func(activeRb string)
}

func (s *subState) handleActiveRbChanged() {
	log.Println(s.activeRb)
	s.onActiveRbChanged(s.activeRb)
}

func buildUIWinAddItem(win gtk.Window, builder gtk.Builder) {
	var state subState
	state.activeRb = "bug"

	rbBug := gtk.WrapRadioButton(builder.GetObject("rbBug").P)
	rbTask := gtk.WrapRadioButton(builder.GetObject("rbTask").P)
	rbIssue := gtk.WrapRadioButton(builder.GetObject("rbIssue").P)

	rbBug.Connect(gtk.SigToggled, func() {
		log.Println("rbBug toggled")
		if rbBug.GetActive() {
			state.activeRb = "bug"
			state.handleActiveRbChanged()
		}
	})
	rbTask.Connect(gtk.SigToggled, func() {
		log.Println("rbTask toggled")
		if rbTask.GetActive() {
			state.activeRb = "task"
			state.handleActiveRbChanged()
		}
	})
	rbIssue.Connect(gtk.SigToggled, func() {
		log.Println("rbIssue toggled")
		if rbIssue.GetActive() {
			state.activeRb = "issue"
			state.handleActiveRbChanged()
		}
	})

	entryNum := gtk.WrapEntry(builder.GetObject("entryNum").P)
	entryUrl := gtk.WrapEntry(builder.GetObject("entryUrl").P)

	entryUrl.Connect(gtk.SigChanged, func() {
		state.url = entryUrl.GetText()
	})

	updateUrlWithNum := func() {
		numStr := entryNum.GetText()
		num, err := strconv.Atoi(numStr)
		log.Println("num:", num)

		if state.activeRb == "bug" || state.activeRb == "task" {
			var urlTxt string
			if err == nil {
				urlTxt = fmt.Sprintf("https://pms.uniontech.com/zentao/%s-view-%d.html",
					state.activeRb, num)
			}
			entryUrl.SetText(urlTxt)
		}
	}

	entryNum.Connect(gtk.SigChanged, func() {
		updateUrlWithNum()
	})

	entryIssue := gtk.WrapEntry(builder.GetObject("entryIssue").P)
	entryIssue.Connect(gtk.SigChanged, func() {
		state.issue = entryIssue.GetText()
	})

	boxNum := gtk.WrapBox(builder.GetObject("boxNum").P)
	boxUrl := gtk.WrapBox(builder.GetObject("boxUrl").P)
	boxIssue := gtk.WrapBox(builder.GetObject("boxIssue").P)

	btnUrlTest := gtk.WrapButton(builder.GetObject("btnUrlTest").P)
	btnYes := gtk.WrapButton(builder.GetObject("btnYes").P)
	btnYes.Connect(gtk.SigClicked, func() {
		log.Println("btnYes clicked")

		var outStr string
		if state.activeRb == "bug" || state.activeRb == "task" {
			outStr = strings.Title(state.activeRb) + ": " + state.url
		} else {
			outStr = "Issue: " + state.issue
		}

		log.Println("outStr: ", outStr)
		win.Close()
	})
	btnNo := gtk.WrapButton(builder.GetObject("btnNo").P)
	btnNo.Connect(gtk.SigClicked, func() {
		log.Println("btnNo clicked")
		win.Close()
	})

	btnUrlTest.Connect(gtk.SigClicked, func() {
		urlTxt := entryUrl.GetText()
		if urlTxt == "" {
			log.Println("urlTxt is empty")
			return
		}
		_, err := url.Parse(urlTxt)
		if err != nil {
			log.Println("WARN:", err)
			return
		}
		go func() {
			err := exec.Command("xdg-open", urlTxt).Run()
			if err != nil {
				log.Println("WARN:", err)
			}
		}()
	})

	state.onActiveRbChanged = func(activeRb string) {
		switch activeRb {
		case "bug", "task":
			boxNum.SetSensitive(true)
			boxUrl.SetSensitive(true)
			boxIssue.SetSensitive(false)
			updateUrlWithNum()
		case "issue":
			boxNum.SetSensitive(false)
			boxUrl.SetSensitive(false)
			boxIssue.SetSensitive(true)
		}
	}
	state.onActiveRbChanged(state.activeRb)

}

func main() {
	gtk.Init(0, 0)
	win := buildUI()
	win.ShowAll()
	gtk.Main()
}
