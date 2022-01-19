package main

import (
	"fmt"
	"log"

	"github.com/electricface/go-gir/gi"
	"github.com/electricface/go-gir/gtk-3.0"
)

const title = "Assistant"
const description = `
Demonstrates a sample multistep assistant. Assistants are used to divide
an operation into several simpler sequential steps, and to guide the user
through these steps
`

type AssistantApp struct {
	assistant gtk.Assistant
}

func newAssitantApp() *AssistantApp {
	aa := &AssistantApp{}
	aa.assistant = gtk.NewAssistant()
	aa.assistant.SetDefaultSize(-1, 300)
	aa.createPage1()
	aa.createPage2()
	aa.createPage3()

	aa.assistant.Connect(gtk.SigCancel, aa.onCloseCancel)
	aa.assistant.Connect(gtk.SigClose, aa.onCloseCancel)
	aa.assistant.Connect(gtk.SigApply, aa.onApply)
	aa.assistant.Connect(gtk.SigPrepare, aa.onPrepare)
	aa.assistant.Show()

	return aa
}

func (aa *AssistantApp) onCloseCancel() {
	log.Println("on close | cancel")
	aa.assistant.Destroy()
	gtk.MainQuit()
}

func (aa *AssistantApp) onApply() {
	log.Println("on apply")
}

func (aa *AssistantApp) onPrepare(args []interface{}) {
	var assistant gtk.Assistant
	var page gtk.Widget
	err := gi.Store(args, &assistant, &page)
	if err != nil {
		log.Fatal(err)
	}
	currentPage := assistant.GetCurrentPage()
	log.Println("prepare currentPage:", currentPage)
	nPages := assistant.GetNPages()
	title := fmt.Sprintf("Sample assistant (%v of %v)", currentPage+1, nPages)
	// NOTE: 这里设置的 title 没用了，只会把 page 的 title 作为窗口的 title。
	assistant.SetTitle(title)
}

func (aa *AssistantApp) onEntryChanged(args []interface{}) {
	var entry gtk.Entry
	err := gi.Store(args, &entry)
	if err != nil {
		log.Fatal(err)
	}
	pageNumber := aa.assistant.GetCurrentPage()
	currentPage := aa.assistant.GetNthPage(pageNumber)
	text := entry.GetText()

	if text != "" {
		aa.assistant.SetPageComplete(currentPage, true)
	} else {
		aa.assistant.SetPageComplete(currentPage, false)
	}
}

func (aa *AssistantApp) createPage1() {
	box := gtk.NewHBox(false, 12)
	box.SetBorderWidth(12)
	label := gtk.NewLabel("You must fill out this entry to continue:")
	box.PackStart(label, false, false, 0)

	entry := gtk.NewEntry()
	box.PackStart(entry, true, true, 0)
	entry.Connect(gtk.SigChanged, aa.onEntryChanged)

	box.ShowAll()
	aa.assistant.AppendPage(box)
	aa.assistant.SetPageTitle(box, "Page 1")
	aa.assistant.SetPageType(box, gtk.AssistantPageTypeIntro)

	pixBuf := aa.assistant.RenderIcon(gtk.STOCK_DIALOG_INFO, int32(gtk.IconSizeDialog), "")
	aa.assistant.SetPageHeaderImage(box, pixBuf)
}

func (aa *AssistantApp) createPage2() {
	box := gtk.NewHBox(false, 12)
	box.SetBorderWidth(12)
	checkButton := gtk.NewCheckButtonWithLabel("This is optional data, you may continue even if you do not check this")
	box.PackStart(checkButton, false, false, 0)

	box.ShowAll()
	aa.assistant.AppendPage(box)
	aa.assistant.SetPageComplete(box, true)
	aa.assistant.SetPageTitle(box, "Page 2")

	pixBuf := aa.assistant.RenderIcon(gtk.STOCK_DIALOG_INFO, int32(gtk.IconSizeDialog), "")
	aa.assistant.SetPageHeaderImage(box, pixBuf)
}

func (aa *AssistantApp) createPage3() {
	label := gtk.NewLabel("This is a confirmation page, press \"Apply\" to apply changes")
	label.Show()
	aa.assistant.AppendPage(label)
	aa.assistant.SetPageComplete(label, true)
	aa.assistant.SetPageTitle(label, "Confirmation")
	aa.assistant.SetPageType(label, gtk.AssistantPageTypeConfirm)

	pixBuf := aa.assistant.RenderIcon(gtk.STOCK_DIALOG_INFO, int32(gtk.IconSizeDialog), "")
	aa.assistant.SetPageHeaderImage(label, pixBuf)
}

func main() {
	gtk.Init(0, 0)
	newAssitantApp()
	gtk.Main()
}
