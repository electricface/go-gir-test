package main

import (
	"log"

	"github.com/electricface/go-gir/g-2.0"
	"github.com/electricface/go-gir/gi"
)

func main() {
	l := g.NewList()
	l.SetData(gi.Uint2Ptr(1))
	l.Append(gi.Uint2Ptr(2))
	l.Append(gi.Uint2Ptr(3))

	for e := l; e.P != nil; e = e.Next() {
		d := e.Data()
		log.Println(d)
	}
}
