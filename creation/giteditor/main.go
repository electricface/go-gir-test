package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unsafe"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gi"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

type line struct {
	type0   string
	content string
}

type State struct {
	typ   string
	scope string
	title string
	desc  string
	log   string

	lines []line
	//activeLine *line
	//activeRowIdx int
	updatePreviewFn func(string)
	addLineFn       func(l line)
}

func (s *State) addLine(l line) {
	s.lines = append(s.lines, l)
	s.addLineFn(l)
	s.updatePreview()
}

func (s *State) clearLines() {
	s.lines = nil
	s.updatePreview()
}

func (s *State) deleteLine(i int) {
	copy(s.lines[i:], s.lines[i+1:])
	s.lines[len(s.lines)-1] = line{}
	s.lines = s.lines[:len(s.lines)-1]
	s.updatePreview()
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

	if len(s.lines) > 0 {
		sb.WriteString("\n")
		for _, l := range s.lines {
			sb.WriteString(fmt.Sprintf("%s: %s\n", strings.Title(l.type0), l.content))
		}
	}

	text := sb.String()
	s.updatePreviewFn(text)
}

type actionUpdate struct {
	prop, value string
}

func buildUI(uiStr string) gtk.Window {
	state := &State{}
	builder := gtk.NewBuilder()
	_, err := builder.AddFromString(uiStr, uint64(len(uiStr)))
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
		start := newTextIter()
		end := newTextIter()
		bufDesc.GetBounds(start, end)
		txt := bufDesc.GetText(start, end, false)
		state.handleAction(&actionUpdate{
			prop:  "desc",
			value: txt,
		})
		gi.Free(start.P)
		gi.Free(end.P)
	})

	tvLog := gtk.WrapTextView(builder.GetObject("tvLog").P)
	bufLog := tvLog.GetBuffer()
	bufLog.Connect(gtk.SigChanged, func() {
		log.Println("bufLog changed")
		start := newTextIter()
		end := newTextIter()
		bufLog.GetBounds(start, end)
		txt := bufLog.GetText(start, end, false)
		state.handleAction(&actionUpdate{
			prop:  "log",
			value: txt,
		})
		gi.Free(start.P)
		gi.Free(end.P)
	})

	tvLines := gtk.WrapTreeView(builder.GetObject("tvLines").P)

	listStore := gtk.WrapListStore(builder.GetObject("liststore1").P)

	state.addLineFn = func(l line) {
		treeIter := newTreeIter()
		listStore.Append(treeIter)
		v1, _ := g.NewValueWith(l.type0)
		listStore.SetValue(treeIter, 0, v1)

		v2, _ := g.NewValueWith(l.content)
		listStore.SetValue(treeIter, 1, v2)

		v2.Free()
		gi.Free(treeIter.P)
	}

	tvPreview := gtk.WrapTextView(builder.GetObject("tvPreview").P)
	bufPreview := tvPreview.GetBuffer()
	state.updatePreviewFn = func(s string) {
		bufPreview.SetText(s, int32(len(s)))
	}
	state.typ = "fix"
	state.updatePreview()

	win.Connect(gtk.SigDestroy, func() {
		os.Exit(0)
	})

	winAddItem := gtk.WrapWindow(builder.GetObject("winAddItem").P)
	buildUIWinAddItem(winAddItem, builder, state)

	btnAddLine := gtk.WrapButton(builder.GetObject("btnAddLine").P)
	btnDeleteLine := gtk.WrapButton(builder.GetObject("btnDeleteLine").P)
	btnClearLines := gtk.WrapButton(builder.GetObject("btnClearLines").P)

	btnAddLine.Connect(gtk.SigClicked, func() {
		winAddItem.SetTransientFor(win)
		winAddItem.ShowAll()
	})
	btnDeleteLine.Connect(gtk.SigClicked, func() {
		log.Println("btnDeleteLine clicked")
		sel := tvLines.GetSelection()
		treePathList, model := sel.GetSelectedRows()

		if treePathList.Length() == 1 {
			treePath := gtk.TreePath{P: treePathList.NthData(0)}
			iter := newTreeIter()
			indices := treePath.GetIndices()
			id := int(indices.AsSlice()[0])
			log.Println("delete id", id)
			state.deleteLine(id)

			model.GetIter(iter, treePath)
			listStore.Remove(iter)
			g.Free(iter.P)
		}

		treePathList.FreeFull(func(item unsafe.Pointer) {
			gtk.TreePath{P: item}.Free()
		})

	})
	btnClearLines.Connect(gtk.SigClicked, func() {
		listStore.Clear()
		state.clearLines()
	})

	return win
}

func newTreeIter() gtk.TreeIter {
	return gtk.TreeIter{P: gi.Malloc(gtk.SizeOfStructTreeIter)}
}

func newTextIter() gtk.TextIter {
	return gtk.TextIter{P: gi.Malloc(gtk.SizeOfStructTextIter)}
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

func buildUIWinAddItem(win gtk.Window, builder gtk.Builder, parentState *State) {
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

		var content string
		if state.activeRb == "bug" || state.activeRb == "task" {
			content = state.url
		} else {
			content = state.issue
		}
		parentState.addLine(line{
			type0:   state.activeRb,
			content: content,
		})
		win.Hide()
	})
	btnNo := gtk.WrapButton(builder.GetObject("btnNo").P)
	btnNo.Connect(gtk.SigClicked, func() {
		log.Println("btnNo clicked")
		win.Hide()
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
	uiGladeData, err := Asset("ui.glade")
	if err != nil {
		log.Fatal(err)
	}
	uiStr := string(uiGladeData)

	app := gtk.NewApplication("com.deepin.giteditor", g.ApplicationFlagsFlagsNone)
	app.Connect(gtk.SigActivate, func(args []interface{}) {
		app := gtk.WrapApplication(args[0].(g.Object).P)
		win := buildUI(uiStr)
		app.AddWindow(win)
		win.ShowAll()
	})
	args := gi.NewCStrArrayWithStrings(os.Args...)
	defer args.FreeAll()
	status := app.Run(int32(len(os.Args)), args)
	app.Unref()
	os.Exit(int(status))
}
