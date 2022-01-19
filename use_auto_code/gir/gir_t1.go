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

	"github.com/electricface/go-gir/girepository-2.0"
)

func main() {
	repo := girepository.RepositoryGetDefault1()
	if repo.P == nil {
		log.Fatal("repo is nil")
	}

	bi := repo.FindByName("GIRepository", "Repository")
	if bi.P == nil {
		log.Fatalln("bi is nil")
	}

	name := bi.GetName()
	log.Println("name:", name)

	typ := bi.GetType()
	log.Println("type:", typ)

	if typ == girepository.InfoTypeObject {
		log.Println("is object")
	}

	numMethods := girepository.ObjectInfoGetNMethods(bi)
	log.Println("num methods:", numMethods)

	for i := int32(0); i < numMethods; i++ {
		mi := girepository.ObjectInfoGetMethod(bi, i)
		symbol := girepository.FunctionInfoGetSymbol(mi)
		log.Println(i, symbol)
	}

	numConstants := girepository.ObjectInfoGetNConstants(bi)
	log.Println("num constants:", numConstants)
}
