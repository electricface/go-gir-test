package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/electricface/go-gir/g-2.0"
	"github.com/electricface/go-gir/gi"
)

func Test1(t *testing.T) {
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

	f4 := g.FileNewForCommandlineArg("./file_test.go")
	uri = f4.GetUri()
	log.Println("uri:", uri)
	uriScheme := f4.GetUriScheme()
	log.Println("uri scheme:", uriScheme)

	f5 := g.FileNewForCommandlineArgAndCwd("./file_test.go", "/tmp")
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
	byteArr := gi.NewUint8Array([]byte("hello world"))
	result, err := ops.Write(byteArr, uint64(byteArr.Len), nil)
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
	path0 = f8.GetPath()
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

func TestReadBytes(t *testing.T) {
	passwdF := g.FileNewForPath("/etc/passwd")
	iStream, err := passwdF.Read(nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		gb, err := iStream.ReadBytes(100, nil)
		if err != nil {
			log.Fatal(err)
		}
		size := gb.GetSize()
		if size == 0 {
			break
		}

		log.Println("size:", size)
		arr := gb.GetData()
		data := arr.AsSlice()
		fmt.Printf("data: %s\n", data)
	}
}

func TestReadAll(t *testing.T) {
	passwdF := g.FileNewForPath("/etc/passwd")
	iStream, err := passwdF.Read(nil)
	if err != nil {
		log.Fatal(err)
	}

	arr := gi.MakeUint8Array(100)
	defer arr.Free()

	for {
		result, bytesRead, err := iStream.ReadAll(arr, uint64(arr.Len), nil)
		if err != nil {
			log.Fatal(err)
		}
		// result 是否成功
		log.Printf("success: %v, bytes read: %v\n", result, bytesRead)
		if bytesRead == 0 {
			break
		}
		data := arr.AsSlice()
		fmt.Printf("data: %s\n", data[:bytesRead])
	}
}

func TestReadAllAsyncCh(t *testing.T) {
	passwdF := g.FileNewForPath("/etc/passwd")
	iStream, err := passwdF.Read(nil)
	if err != nil {
		log.Fatal(err)
	}

	mainCtx := g.MainContextDefault()
	mainLoop := g.NewMainLoop(mainCtx, false)
	go mainLoop.Run()

	arr := gi.MakeUint8Array(100)
	defer arr.Free()

	ch := make(chan bool)
	for {
		iStream.ReadAllAsync(arr, uint64(arr.Len), 0, nil, func(sourceObject g.Object, res g.AsyncResult) {
			result, bytesRead, err := iStream.ReadAllFinish(res)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(result, bytesRead)
			data := arr.AsSlice()[:bytesRead]
			log.Printf("data: %s\n", data)
			if bytesRead > 0 {
				ch <- false
			} else {
				ch <- true
			}
		})

		isEnd := <-ch
		if isEnd {
			break
		}
	}
	mainLoop.Quit()
}

func TestReadAllAsyncWg(t *testing.T) {
	passwdF := g.FileNewForPath("/etc/passwd")
	iStream, err := passwdF.Read(nil)
	if err != nil {
		log.Fatal(err)
	}

	mainCtx := g.MainContextDefault()
	mainLoop := g.NewMainLoop(mainCtx, false)
	go mainLoop.Run()

	arr := gi.MakeUint8Array(100)
	defer arr.Free()

	var wg sync.WaitGroup
	for {
		wg.Add(1)
		isEnd := false

		iStream.ReadAllAsync(arr, uint64(arr.Len), 0, nil, func(sourceObject g.Object, res g.AsyncResult) {
			result, bytesRead, err := iStream.ReadAllFinish(res)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(result, bytesRead)
			data := arr.AsSlice()[:bytesRead]
			log.Printf("data: %s\n", data)
			if bytesRead > 0 {
				isEnd = false
			} else {
				isEnd = true
			}
			wg.Done()
		})

		wg.Wait()
		if isEnd {
			break
		}
	}
	mainLoop.Quit()
}
