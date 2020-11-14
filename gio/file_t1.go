package main

import (
	"log"
	"os"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/linuxdeepin/go-gir/gi"
)

func main() {
	f1 := g.FileNewForPath("/tmp")
	uri := f1.GetUri()
	log.Println("uri:", uri)

	f2 := g.FileNewForPath1("/tmp")
	uri = f2.GetUri()
	log.Println("uri:", uri)

	f3 := g.FileNewForUri("file:///tmp")
	uri = f3.GetUri()
	log.Println(uri)
	log.Println(f3.GetPath())

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pwd)

	f4 := g.FileNewForCommandlineArg("./file_t1.go")
	uri = f4.GetUri()
	log.Println("uri:", uri)
	uriScheme := f4.GetUriScheme()
	log.Println("uri scheme:", uriScheme)

	f5 := g.FileNewForCommandlineArgAndCwd("./file_t1.go", "/tmp")
	path0 := f5.GetPath()
	log.Println("path0:", path0)

	// 写临时文件
	f6, out, err := g.FileNewTmp(gi.NilStr)
	if err != nil {
		log.Fatal(err)
	}
	path0 = f6.GetPath()
	log.Println("path0:", path0)
	ops := out.GetOutputStream()
	arr := []byte("hello world")
	byteArr := gi.NewUint8Array(arr...)
	result, err := ops.Write(byteArr, uint64(len(arr)), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("result:", result)

	out.Unref()

	f7 := g.FileParseName("file:///home/del1/abc")
	path0 = f7.GetPath()
	log.Println("path0:", path0)
	parseName := f7.GetParseName()
	log.Println("parse name:", parseName)

	f8 := f7.Dup()
	path0  = f8.GetPath()
	log.Println(path0)

	hash := f8.Hash()
	log.Println("hash:", hash)

	log.Println("f7 = f8 ? ", f7.Equal(f8))
	log.Println("f7 = f6 ? ", f7.Equal(f6))

	basename := f8.GetBasename()
	log.Println("basename:", basename)

	path0 = f8.PeekPath()
	log.Println("path:", path0)

	uri = f8.GetUri()
	log.Println("uri:", uri)

	parent := f8.GetParent()
	path0 = parent.GetPath()
	log.Println("path:", path0)

	f9 := g.FileNewForPath("/")
	parent = f9.GetParent()
	path0 = parent.GetPath()
	log.Println("path:", path0)
}