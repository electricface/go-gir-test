package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/electricface/go-gir/g-2.0"
	"github.com/electricface/go-gir/gdkpixbuf-2.0"
	gi "github.com/electricface/go-gir/gi"
	"github.com/electricface/go-gir/gtk-3.0"
)

func uniq(strings []string) (ret []string) {
	return
}

func authors() []string {
	if b, err := exec.Command("git", "log").Output(); err == nil {
		lines := strings.Split(string(b), "\n")

		var a []string
		r := regexp.MustCompile(`^Author:\s*([^ <]+).*$`)
		for _, e := range lines {
			ms := r.FindStringSubmatch(e)
			if ms == nil {
				continue
			}
			a = append(a, ms[1])
		}
		sort.Strings(a)
		var p string
		lines = []string{}
		for _, e := range a {
			if p == e {
				continue
			}
			lines = append(lines, e)
			p = e
		}
		return lines
	}
	return []string{"Yasuhiro Matsumoto <mattn.jp@gmail.com>"}
}

func main() {
	gtk.Init(0, 0)

	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetPosition(gtk.WindowPositionCenter)
	window.SetTitle("GTK - Go!")
	window.SetIconName("gtk-dialog-info")
	window.Connect(gtk.SigDestroy, func() {
		gtk.MainQuit()
	})

	//--------------------------------------------------------
	// GtkVBox
	//--------------------------------------------------------
	vbox := gtk.NewVBox(false, 1)

	//--------------------------------------------------------
	// GtkMenuBar
	//--------------------------------------------------------
	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	//--------------------------------------------------------
	// GtkVPaned
	//--------------------------------------------------------
	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	//--------------------------------------------------------
	// GtkFrame
	//--------------------------------------------------------
	frame1 := gtk.NewFrame("Demo")
	framebox1 := gtk.NewVBox(false, 1)
	frame1.Add(framebox1)

	frame2 := gtk.NewFrame("Demo")
	framebox2 := gtk.NewVBox(false, 1)
	frame2.Add(framebox2)

	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)

	//--------------------------------------------------------
	// GtkImage
	//--------------------------------------------------------
	//dir, _ := filepath.Split(os.Args[0])
	imagefile := "/home/tp1/program_lang/go/src/github.com/mattn/go-gtk/data/go-gtk-logo.png"

	label := gtk.NewLabel("Go Binding for GTK")
	//label.ModifyFontEasy("DejaVu Serif 15")
	framebox1.PackStart(label, false, true, 0)

	//--------------------------------------------------------
	// GtkEntry
	//--------------------------------------------------------
	entry := gtk.NewEntry()
	entry.SetText("Hello world")
	framebox1.Add(entry)

	image := gtk.NewImageFromFile(imagefile)
	framebox1.Add(image)

	//--------------------------------------------------------
	// GtkScale
	//--------------------------------------------------------
	scale := gtk.NewHScaleWithRange(0, 100, 1)
	scale.Connect(gtk.SigValueChanged, func() {
		fmt.Println("scale:", int(scale.GetValue()))
	})
	framebox2.Add(scale)

	//--------------------------------------------------------
	// GtkHBox
	//--------------------------------------------------------
	buttons := gtk.NewHBox(false, 1)

	//--------------------------------------------------------
	// GtkButton
	//--------------------------------------------------------
	button := gtk.NewButtonWithLabel("Button with label")
	button.Connect(gtk.SigClicked, func() {
		fmt.Println("button clicked:", button.GetLabel())
		//messagedialog := gtk.NewMessageDialog(
		//	button.GetTopLevelAsWindow(),
		//	gtk.DIALOG_MODAL,
		//	gtk.MESSAGE_INFO,
		//	gtk.BUTTONS_OK,
		//	entry.GetText())
		//messagedialog.Response(func() {
		//	fmt.Println("Dialog OK!")
		//
		//	//--------------------------------------------------------
		//	// GtkFileChooserDialog
		//	//--------------------------------------------------------
		//	filechooserdialog := gtk.NewFileChooserDialog(
		//		"Choose File...",
		//		button.GetTopLevelAsWindow(),
		//		gtk.FILE_CHOOSER_ACTION_OPEN,
		//		gtk.STOCK_OK,
		//		gtk.RESPONSE_ACCEPT)
		//	filter := gtk.NewFileFilter()
		//	filter.AddPattern("*.go")
		//	filechooserdialog.AddFilter(filter)
		//	filechooserdialog.Response(func() {
		//		fmt.Println(filechooserdialog.GetFilename())
		//		filechooserdialog.Destroy()
		//	})
		//	filechooserdialog.Run()
		//	messagedialog.Destroy()
		//})
		//messagedialog.Run()
	})
	buttons.Add(button)

	//--------------------------------------------------------
	// GtkFontButton
	//--------------------------------------------------------
	fontbutton := gtk.NewFontButton()
	fontbutton.Connect(gtk.SigFontSet, func() {
		fmt.Println("title:", fontbutton.GetTitle())
		fmt.Println("fontname:", fontbutton.GetFontName())
		fmt.Println("use_size:", fontbutton.GetUseSize())
		fmt.Println("show_size:", fontbutton.GetShowSize())
	})
	buttons.Add(fontbutton)
	framebox2.PackStart(buttons, false, false, 0)

	buttons = gtk.NewHBox(false, 1)

	//--------------------------------------------------------
	// GtkToggleButton
	//--------------------------------------------------------
	togglebutton := gtk.NewToggleButtonWithLabel("ToggleButton with label")
	togglebutton.Connect(gtk.SigToggled, func() {
		if togglebutton.GetActive() {
			togglebutton.SetLabel("ToggleButton ON!")
		} else {
			togglebutton.SetLabel("ToggleButton OFF!")
		}
	})
	buttons.Add(togglebutton)

	//--------------------------------------------------------
	// GtkCheckButton
	//--------------------------------------------------------
	checkbutton := gtk.NewCheckButtonWithLabel("CheckButton with label")
	checkbutton.Connect(gtk.SigToggled, func() {
		if checkbutton.GetActive() {
			checkbutton.SetLabel("CheckButton CHECKED!")
		} else {
			checkbutton.SetLabel("CheckButton UNCHECKED!")
		}
	})
	buttons.Add(checkbutton)

	//--------------------------------------------------------
	// GtkRadioButton
	//--------------------------------------------------------
	buttonbox := gtk.NewVBox(false, 1)
	radiofirst := gtk.NewRadioButtonWithLabel(g.SList{}, "Radio1")
	buttonbox.Add(radiofirst)
	buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio2"))
	buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio3"))
	buttons.Add(buttonbox)
	//radiobutton.SetMode(false);
	radiofirst.SetActive(true)

	framebox2.PackStart(buttons, false, false, 0)

	//--------------------------------------------------------
	// GtkVSeparator
	//--------------------------------------------------------
	vsep := gtk.NewVSeparator()
	framebox2.PackStart(vsep, false, false, 0)

	//--------------------------------------------------------
	// GtkComboBoxEntry
	//--------------------------------------------------------
	combos := gtk.NewHBox(false, 1)
	comboboxentry := gtk.NewComboBoxText()
	comboboxentry.AppendText("Monkey")
	comboboxentry.AppendText("Tiger")
	comboboxentry.AppendText("Elephant")
	comboboxentry.Connect(gtk.SigChanged, func() {
		fmt.Println("value:", comboboxentry.GetActiveText())
	})
	combos.Add(comboboxentry)

	//--------------------------------------------------------
	// GtkComboBox
	//--------------------------------------------------------
	combobox := gtk.NewComboBoxText()
	combobox.AppendText("Peach")
	combobox.AppendText("Banana")
	combobox.AppendText("Apple")
	combobox.SetActive(1)
	combobox.Connect(gtk.SigChanged, func() {
		fmt.Println("value:", combobox.GetActiveText())
	})
	combos.Add(combobox)

	framebox2.PackStart(combos, false, false, 0)

	//--------------------------------------------------------
	// GtkTextView
	//--------------------------------------------------------
	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.PolicyTypeAutomatic, gtk.PolicyTypeAutomatic)
	swin.SetShadowType(gtk.ShadowTypeIn)
	textview := gtk.NewTextView()
	buffer := textview.GetBuffer()
	start := gtk.TextIter{P: gi.Malloc(gtk.SizeOfStructTextIter)}
	end := gtk.TextIter{P: gi.Malloc(gtk.SizeOfStructTextIter)}
	buffer.GetStartIter(start)
	buffer.Insert(start, "Hello ", int32(len("Hello ")))
	buffer.GetEndIter(end)
	buffer.Insert(end, "World!", int32(len("World!")))

	tagTable := buffer.GetTagTable()
	tag := gtk.NewTextTag("bold")
	tagTable.Add(tag)

	val0, _ := g.NewValueWith("#FF0000")
	tag.SetProperty("background", val0)
	val0.Free()

	val1, _ := g.NewValueWith(700)
	tag.SetProperty("weight", val1)
	val1.Free()

	buffer.GetStartIter(start)
	buffer.GetEndIter(end)
	buffer.ApplyTag(tag, start, end)

	swin.Add(textview)
	framebox2.Add(swin)

	buffer.Connect(gtk.SigChanged, func() {
		fmt.Println("changed")
	})

	//--------------------------------------------------------
	// GtkMenuItem
	//--------------------------------------------------------
	cascademenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(cascademenu)
	submenu := gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	menuitem := gtk.NewMenuItemWithMnemonic("E_xit")
	menuitem.Connect(gtk.SigActivate, func() {
		gtk.MainQuit()
	})
	submenu.Append(menuitem)

	cascademenu = gtk.NewMenuItemWithMnemonic("_View")
	menubar.Append(cascademenu)
	submenu = gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	checkmenuitem := gtk.NewCheckMenuItemWithMnemonic("_Disable")
	checkmenuitem.Connect(gtk.SigActivate, func() {
		vpaned.SetSensitive(!checkmenuitem.GetActive())
	})
	submenu.Append(checkmenuitem)

	menuitem = gtk.NewMenuItemWithMnemonic("_Font")
	menuitem.Connect(gtk.SigActivate, func() {
		fsd := gtk.NewFontSelectionDialog("Font")
		fsd.SetFontName(fontbutton.GetFontName())
		fsd.Connect(gtk.SigResponse, func() {
			fmt.Println(fsd.GetFontName())
			fontbutton.SetFontName(fsd.GetFontName())
			fsd.Destroy()
		})
		fsd.SetTransientFor(window)
		fsd.Run()
	})
	submenu.Append(menuitem)

	cascademenu = gtk.NewMenuItemWithMnemonic("_Help")
	menubar.Append(cascademenu)
	submenu = gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	menuitem = gtk.NewMenuItemWithMnemonic("_About")
	menuitem.Connect(gtk.SigActivate, func() {
		dialog := gtk.NewAboutDialog()
		dialog.SetName("Go-Gtk Demo!")
		dialog.SetProgramName("demo")
		authorsArr := gi.NewCStrArrayZTWithStrings(authors()...)
		dialog.SetAuthors(authorsArr)
		authorsArr.FreeAll()

		//imagefile := "/home/tp1/program_lang/go/src/github.com/mattn/go-gtk/data/go-gtk-logo.png"
		imagefile := filepath.Join("/home/tp1/program_lang/go/src/github.com/mattn/go-gtk/data/mattn-logo.png")
		pixbuf, _ := gdkpixbuf.NewPixbufFromFile(imagefile)
		dialog.SetLogo(pixbuf)
		dialog.SetLicense("The library is available under the same terms and conditions as the Go, the BSD style license, and the LGPL (Lesser GNU Public License). The idea is that if you can use Go (and Gtk) in a project, you should also be able to use go-gtk.")
		dialog.SetWrapLicense(true)
		dialog.Run()
		dialog.Destroy()
	})
	submenu.Append(menuitem)

	//--------------------------------------------------------
	// GtkStatusbar
	//--------------------------------------------------------
	statusbar := gtk.NewStatusbar()
	context_id := statusbar.GetContextId("go-gtk")
	statusbar.Push(context_id, "GTK binding for Go!")

	framebox2.PackStart(statusbar, false, false, 0)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
