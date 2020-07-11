package main

import (
	"log"
	"unsafe"

	"github.com/electricface/go-gir3/gi"
)

func main() {
	log.SetFlags(log.Lshortfile)
	main3()
}

func main1() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Gio", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		log.Fatal(err)
	}
	loadedNs := repo.LoadedNamespaces()
	log.Println(loadedNs)
	num := repo.NumInfo("Gio")
	log.Println(num)
	for i := 0; i < num; i++ {
		bi := repo.Info("Gio", i)
		name := bi.Name()
		log.Println(name)

		//if name == "File" {
		//	log.Println(bi.Type().String())
		//	ifcInfo := gi.ToInterfaceInfo(bi)
		//	numMethod := ifcInfo.NumMethod()
		//	for j := 0; j < numMethod; j++ {
		//		funcInfo := ifcInfo.Method(j)
		//		if funcInfo.Name() == "new_for_path" {
		//			// test it
		//			var ret gi.Argument
		//			path1 := "/home/tp1"
		//			_ = path1
		//			//err := funcInfo.Invoke([]gi.Argument{ gi.NewStringArgument(strPtr(&path1)) }, []gi.Argument{}, &ret)
		//			spew.Dump(ret)
		//			//if err != nil {
		//			//	log.Fatal(err)
		//			//}
		//			log.Println("invoke success")
		//		}
		//	}
		//	return
		//}
	}

	//repo.FindByName()
}

var invokerCacheGlib = gi.NewInvokerCache("GLib")
var invokerCacheGio = gi.NewInvokerCache("Gio")

func main3() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GLib", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		log.Fatal(err)
	}
	_, err = repo.Require("Gio", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		log.Fatal(err)
	}

	fp := FileNewForPath("/home/tp1/hello-world")
	log.Printf("fp: %p\n", fp)

	uri := FileGetUri(fp)
	log.Printf("uri: %q\n", uri)
	//
	//basename := FileGetBasename(fp)
	//log.Printf("basename: %q\n", basename)
	//
	//path1 := FileGetPath(fp)
	//log.Printf("path: %q\n",path1)
	//
	//dai := DesktopAppInfoNew("dde-control-center.desktop")
	//log.Printf("dai: %p\n", dai)

	acc := access("/tmp/", 1)
	log.Println("acc:", acc)

	buf := "------------------------------"
	result := ascii_formatd(buf, int32(len(buf)), "%f", 3.1415926)
	log.Println(" ascii_formatd result:", result)

	result0, endptr := ascii_strtod("6.232abc")
	log.Printf("ascii_strtod result: %v, endptr: %q\n", result0, endptr)

	strForB64 := "hello world"
	strLen := uint64(len(strForB64) + 1)
	resultForB64 := base64_decode_inplace(&strForB64, &strLen)
	log.Printf("strForB64: %s, strLen: %d, resultForB64: %s\n", strForB64, strLen, resultForB64)
}

const (
	File_new_for_path uint = iota
	File_get_uri
	File_get_path
	File_get_basename
	DesktopAppInfo_new
	F_access
	F_ascii_formatd
	F_ascii_strtod
	F_base64_decode_inplace
)

func DesktopAppInfoNew(desktopId string) (result unsafe.Pointer) {
	invoker, err := invokerCacheGio.Get(DesktopAppInfo_new, "DesktopAppInfo", "new")
	if err != nil {
		log.Fatal(err)
	}

	var ret gi.Argument
	pPath := gi.CString(desktopId)
	arg1 := gi.NewStringArgument(pPath)
	invoker.Call([]gi.Argument{arg1}, &ret, nil)
	gi.Free(pPath)

	result = ret.Pointer()

	return
}

func FileNewForPath(path1 string) (result unsafe.Pointer) {
	invoker, err := invokerCacheGio.Get(File_new_for_path, "File", "new_for_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}

	var ret gi.Argument
	pPath := gi.CString(path1)
	arg1 := gi.NewStringArgument(pPath)
	invoker.Call([]gi.Argument{arg1}, &ret, nil)
	gi.Free(pPath)
	result = ret.Pointer()
	return
}

func FileGetUri(p unsafe.Pointer) string {
	invoker, err := invokerCacheGio.Get(File_get_uri, "File", "get_uri")
	if err != nil {
		log.Fatal(err)
	}

	arg1 := gi.NewPointerArgument(p)
	var ret gi.Argument
	if err != nil {
		log.Fatal(err)
	}
	invoker.Call([]gi.Argument{arg1}, &ret, nil)
	return ret.String().Take()
}

func FileGetPath(p unsafe.Pointer) string {
	invoker, err := invokerCacheGio.Get(File_get_path, "File", "get_uri")
	if err != nil {
		log.Fatal(err)
	}

	arg1 := gi.NewPointerArgument(p)
	var ret gi.Argument
	invoker.Call([]gi.Argument{arg1}, &ret, nil)
	return ret.String().Take()
}

func FileGetBasename(p unsafe.Pointer) string {
	invoker, err := invokerCacheGio.Get(File_get_basename, "File", "get_uri")
	if err != nil {
		log.Fatal(err)
	}

	arg1 := gi.NewPointerArgument(p)
	var ret gi.Argument
	invoker.Call([]gi.Argument{arg1}, &ret, nil)
	return ret.String().Take()
}

// g_access
//func Access(filename string, mode int) int {
//	invoker, err := invokerCache.Get(F_Access, "access", "")
//	if err != nil {
//		log.Fatal(err)
//	}
//	pFilename := gi.CString(filename)
//	arg0 := gi.NewPointerArgument(pFilename)
//	arg1 := gi.NewIntArgument(mode)
//	args := []gi.Argument{arg0,arg1}
//	var ret gi.Argument
//	invoker.Call(args, &ret)
//
//	// after call
//
//	//invoker.Call(args, nil)
//}

func access(filename string, mode int32) int {
	invoker, err := invokerCacheGlib.Get(F_access, "access", "")
	if err != nil {
		log.Fatal(err)
	}
	p_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(p_filename)
	arg_mode := gi.NewInt32Argument(mode)
	args := []gi.Argument{arg_filename, arg_mode}
	var ret gi.Argument
	invoker.Call(args, &ret, nil)
	gi.Free(p_filename)
	return ret.Int()
}

func ascii_formatd(buffer string, buf_len int32, format string, d float64) (result string) {
	invoker, err := invokerCacheGlib.Get(F_ascii_formatd, "ascii_formatd", "")
	if err != nil {
		log.Fatal(err)
	}
	p_buffer := gi.CString(buffer)
	p_format := gi.CString(format)
	arg_buffer := gi.NewStringArgument(p_buffer)
	arg_buf_len := gi.NewInt32Argument(buf_len)
	arg_format := gi.NewStringArgument(p_format)
	arg_d := gi.NewDoubleArgument(d)
	args := []gi.Argument{arg_buffer, arg_buf_len, arg_format, arg_d}
	var ret gi.Argument
	invoker.Call(args, &ret, nil)
	//gi.Free(p_buffer)
	gi.Free(p_format)
	result = ret.String().Take()
	return
}

// g_ascii_strtod
func ascii_strtod(nptr string) (result float64, endptr string) {
	iv, err := invokerCacheGlib.Get(F_ascii_strtod, "ascii_strtod", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}

	var outArgs [1]gi.Argument
	c_nptr := gi.CString(nptr)
	arg_nptr := gi.NewStringArgument(c_nptr)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_nptr, arg_endptr}
	var ret gi.Argument

	iv.Call(args, &ret, &outArgs[0])

	result = ret.Double()
	val_endptr := outArgs[0].Pointer()
	endptr = gi.GoString(val_endptr)

	gi.Free(c_nptr)
	return
}

// g_base64_decode_inplace
/* guchar *
g_base64_decode_inplace (gchar *text,
                         gsize *out_len); */
func base64_decode_inplace(text *string, out_len *uint64) (result string) {
	iv, err := invokerCacheGlib.Get(F_base64_decode_inplace, "ascii_strtod", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_text := gi.CString(*text)
	//outArgs[0] = gi.NewStringArgument(c_text)
	outArgs[1] = gi.NewUint64Argument(*out_len)

	log.Println("begin debug")
	log.Println("outArgs[0]value", outArgs[0].String().Copy())
	log.Println("outArgs[1]size", outArgs[1].Size())
	log.Println("end debug")

	arg_text := gi.NewPointerArgument(c_text)
	arg_out_len := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_text, arg_out_len}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])

	log.Println("after call, begin debug")
	//log.Println("outArgs[0]value", outArgs[0].String().Copy())
	log.Println("outArgs[1]size", outArgs[1].Size())
	log.Println("end debug")

	log.Println("c_text:", gi.GoString(c_text))

	*text = outArgs[0].String().Copy()
	*out_len = outArgs[1].Size()
	log.Println("ret.Pointer:", ret.Pointer())
	result = ret.String().Copy()
	gi.Free(c_text)
	return
}
