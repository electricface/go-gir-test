package main

import "C"
import (
	"log"
	"unsafe"

	"github.com/electricface/go-gir/gtk-3.0"
	"github.com/electricface/go-gir/gi"
)

func main() {
	win := gtk.NewWindow(gtk.WindowTypeToplevel)
	//g.TypeFundamental()
	//win.Object
	//win.SetBuildableProperty(gtk.Builder{}, "", g.Value{})
	//
	//win.P_Buildable()
	//var buildable gtk.IBuildable
	//buildable = win
	//_ = buildable

	size := unsafe.Sizeof(gi.GType(0))
	log.Println(size)

	var i gtk.IAdjustment
	i = nil
	val := i.P_Adjustment()
	log.Println(val)

	var tmp1 unsafe.Pointer
	if i != nil {
		tmp1 = i.P_Adjustment()
	}
	gi.NewPointerArgument(tmp1)
	//_, arr, err := glib.FileGetContents("/etc/fstab")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//stat, err := os.Stat("/etc/fstab")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//length := stat.Size()
	//
	//data := arr.AsSlice()
	//log.Printf("length: %v, len(data): %v\n", length, len(data))
	//fmt.Printf("%s\n", data)
	//arr.Free()

	//kf := glib.NewKeyFile()
	//_, err := kf.LoadFromFile("/home/tp1/go/src/github.com/electricface/go-gir3/test/use_auto_code/test_file", glib.KeyFileFlagsNone)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//arr, err := kf.GetDoubleList("Group1","ds")
	//if err != nil {
	//	log.Fatal(err)
	//}
	////arr := gi.DoubleArray{
	////	P:   unsafe.Pointer(uintptr(p)),
	////	Len: int(length),
	////}
	//vals := arr.Copy()
	//log.Println("vals:", vals)
	//
	//vals1 := arr.AsSlice()
	//log.Println("vals1:", vals1)
	//arr.Free()
	//
	//d1, err := kf.GetDouble("Group1", "d1")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("d1:", d1)
	//
	//arr1 := gi.NewDoubleArray(3.78,9.78,2.89)
	//kf.SetDoubleList("Group1", "outDs", arr1, uint64(arr1.Len))
	//arr1.Free()
	//
	//
	//arr2 := gi.NewBoolArray(true, false, true, false, true ,true, false, false)
	//kf.SetBooleanList("Group1", "outBs", arr2, uint64(arr2.Len))
	//arr2Slice := arr2.AsSlice()
	//log.Println("arr2Slice:", arr2Slice)
	//arr2.Free()
	//
	//arr3 := gi.NewInt32Array(1,2,3,4,5,6)
	//kf.SetIntegerList("Group1", "outIs", arr3, uint64(arr2.Len))
	//arr3.Free()
	//
	//_, err = kf.SaveToFile("/tmp/out")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//arr4, err := kf.GetBooleanList("Group1", "outBs")
	//if err == nil {
	//	log.Println("arr4:", arr4.AsSlice(), arr4.Copy())
	//	arr4.Free()
	//}  else {
	//	log.Println("WARN:", err)
	//}
	//
	//strArr, err := kf.GetStringList("Group1", "strv")
	//if err == nil {
	//	log.Println("strArr:", strArr.AsSlice())
	//	log.Println("strArr:", strArr.Copy())
	//	strArr.Free()
	//}
	//
	//si := unsafe.Sizeof(C.int(0))
	//log.Println("size of c.int:", si)
	//sb := unsafe.Sizeof(true)
	//log.Println("size of bool:", sb)

	//mode := int32(unix.W_OK)
	//result := glib.Access("/home/tp1/startdde-panic.txt", mode)
	//log.Println("result:", result)
	//
	//// 这一个会 panic
	////buf := "------------------------------"
	////result1 := glib.AsciiFormatd(buf, int32(len(buf)), "%f", 3.1415926)
	////log.Println(" ascii_formatd result:", result1)
	//ok := gio.ActionNameIsValid("changed*")
	//log.Println("ok:", ok)
	//
	//ok = gio.ActionNameIsValid1("changed*")
	//log.Println("ok:", ok)
	//
	//bind := gobject.WrapBinding(unsafe.Pointer(nil))
	//bind.BindProperty("", bind, "", 0)
}
