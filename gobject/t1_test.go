package gobject

import (
	"log"
	"testing"

	"github.com/linuxdeepin/go-gir/g-2.0"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	busConn, err := g.BusGetSync(g.BusTypeSession, nil)
	if err != nil {
		log.Fatal(err)
	}

	vGuid, err := busConn.Get("guid")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("str:", vGuid.GetString())
	vGuid.Free()

	vStream, err := busConn.Get("stream")
	if err != nil {
		log.Fatal(err)
	}
	defer vStream.Free()
	ioStream := g.WrapIOStream(vStream.GetObject().P)
	isClose := ioStream.IsClosed()
	log.Println("is closed:", isClose)

	var ioStream1 g.IOStream
	err = vStream.Store(&ioStream1)
	isClose = ioStream1.IsClosed()
	log.Println("2 is closed:", isClose)

	var p g.ParamSpec
	err = vStream.Store(&p)
	assert.Error(t, err)

	var myGuid string
	var myStream g.IOStream
	err = busConn.GetProperties([]string{"guid", "stream", "guid"}, &myGuid, &myStream, &myGuid)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(myGuid)
	log.Println(myStream.IsClosed())
}
