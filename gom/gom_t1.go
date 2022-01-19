package main

import (
	"log"

	"github.com/electricface/go-gir/gom-1.0"
)

func main() {
	adapter := gom.NewAdapter()
	log.Println(adapter)
}
