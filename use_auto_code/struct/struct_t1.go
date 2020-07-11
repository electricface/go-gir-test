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

	"github.com/linuxdeepin/go-gir/gi"
	"github.com/linuxdeepin/go-gir/gtk-3.0"
)

func main() {
	gtk.Init(0, 0)
	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	objClass := win.GetClass()
	wrp := objClass.FindProperty("width-request")
	if wrp.P == nil {
		log.Fatal("not found property")
	}

	r := gi.DefaultRepository()
	_, err := r.Require("GObject", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		log.Fatalln(err)
	}
	psi := r.FindByName("GObject", "ParamSpec")
	if psi.P == nil {
		log.Fatalln("psi is nil")
	}
	psi0 := gi.WrapStructInfo(psi.P)
	nf := psi0.NumFields()
	log.Println("num fields:", nf)

	vtf := psi0.FindField("flags")
	if vtf.P == nil {
		log.Fatalln("vtf is nil")
	}

	vt, ok := vtf.GetField(win.P)
	if ok {
		log.Println("get field ok")
		log.Println(vt.Int64())
	} else {
		log.Println("get field not ok")
	}
}
