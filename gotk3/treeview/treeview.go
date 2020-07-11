/*
 * Copyright (C) 2019 ~ 2020 Uniontech Software Technology Co.,Ltd
 *
 * Author:
 *
 * Maintainer:
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"log"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gi"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

// IDs to access the tree view columns by
const (
	ColumnVersion = iota
	ColumnFeature
)

func createColumn(title string, id int) gtk.TreeViewColumn {
	cellRenderer := gtk.NewCellRendererText()
	col := gtk.NewTreeViewColumnWithAttribute(title, cellRenderer, "text", int32(id))
	return col
}

func main() {
	gtk.Init(0, 0)
	win := setupWindow("Go Feature Timeline")
	treeView, listStore := setupTreeView()

	// Add some rows to the list store
	addRow(listStore, "r57", "Gofix command added for rewriting code for new APIs")
	addRow(listStore, "r60", "URL parsing moved to new \"url\" package")
	addRow(listStore, "go1.0", "Rune type introduced as alias for int32")
	addRow(listStore, "go1.1", "Race detector added to tools")
	addRow(listStore, "go1.2", "Limit for number of threads added")
	addRow(listStore, "go1.3", "Support for various BSD's, Plan 9 and Solaris")

	win.Add(treeView)
	win.ShowAll()
	gtk.Main()
}

func addRow(listStore gtk.ListStore, version, feature string) {
	iter := gtk.TreeIter{P: gi.Malloc(gtk.SizeOfStructTreeIter)}
	listStore.Append(iter)
	defer gi.Free(iter.P)

	f, err := g.NewValueWith(feature)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Free()

	v, err := g.NewValueWith(version)
	if err != nil {
		log.Fatal(err)
	}
	defer v.Free()

	listStore.SetValue(iter, ColumnVersion, v)
	listStore.SetValue(iter, ColumnFeature, f)
}

func setupTreeView() (gtk.TreeView, gtk.ListStore) {
	treeView := gtk.NewTreeView()
	col1 := createColumn("version", ColumnVersion)
	col2 := createColumn("feature", ColumnFeature)
	treeView.AppendColumn(col1)
	treeView.AppendColumn(col2)
	arr := gi.NewGTypeArray(g.TYPE_STRING, g.TYPE_STRING)
	listStore := gtk.NewListStore(2, arr)
	treeView.SetModel(listStore)
	return treeView, listStore
}

func setupWindow(title string) gtk.Window {
	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	win.SetTitle(title)
	win.Connect(gtk.SigDestroy, func() {
		gtk.MainQuit()
	})
	win.SetPosition(gtk.WindowPositionCenter)
	win.SetDefaultSize(600, 300)
	return win
}
