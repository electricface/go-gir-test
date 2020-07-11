/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"container/list"
	"fmt"
	"log"

	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

var labelList = list.New()

func main() {
	gtk.Init(0, 0)

	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	win.SetTitle("Add/Remove Widgets Example")
	win.Connect(gtk.SigDestroy, func() {
		gtk.MainQuit()
	})

	win.Add(windowWidget())
	win.ShowAll()

	gtk.Main()
}

func windowWidget() gtk.Widget {
	grid := gtk.NewGrid()
	grid.SetOrientation(gtk.OrientationVertical)

	sw := gtk.NewScrolledWindow(nil, nil)

	grid.Attach(sw, 0, 0, 2, 1)
	sw.SetHexpand(true)
	sw.SetVexpand(true)

	labelsGrid := gtk.NewGrid()
	labelsGrid.SetOrientation(gtk.OrientationVertical)

	sw.Add(labelsGrid)
	labelsGrid.SetHexpand(true)

	insertBtn := gtk.NewButtonWithLabel("Add a label")
	removeBtn := gtk.NewButtonWithLabel("Remove a label")

	nLabels := 1
	insertBtn.Connect(gtk.SigClicked, func() {
		var s string
		if nLabels == 1 {
			s = fmt.Sprintf("Inserted %d label.", nLabels)
		} else {
			s = fmt.Sprintf("Inserted %d labels.", nLabels)
		}
		label := gtk.NewLabel(s)

		labelList.PushBack(label)
		labelsGrid.Add(label)
		label.SetHexpand(true)
		labelsGrid.ShowAll()

		nLabels++
	})

	removeBtn.Connect(gtk.SigClicked, func() {
		e := labelList.Front()
		if e == nil {
			log.Print("Nothing to remove")
			return
		}
		lab, ok := labelList.Remove(e).(gtk.Label)
		if !ok {
			log.Print("Element to remove is not a *gtk.Label")
			return
		}
		lab.Destroy()
	})

	grid.Attach(insertBtn, 0, 1, 1, 1)
	grid.Attach(removeBtn, 1, 1, 1, 1)

	return grid.Widget
}
