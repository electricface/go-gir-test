package source

import (
	"log"
	"testing"
	"time"

	"github.com/linuxdeepin/go-gir/g-2.0"
)

func TestIdleAdd(t *testing.T) {
	_, err := g.IdleAdd(func() bool {
		log.Println("call idle func")
		time.Sleep(1 * time.Second)
		return true
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := g.MainContextDefault()
	mainLoop := g.NewMainLoop(ctx, false)
	go mainLoop.Run()
	time.Sleep(10 * time.Second)
	mainLoop.Quit()
}

func TestTimeoutAdd(t *testing.T) {
	_, err := g.TimeoutAdd(1*time.Second, func() bool {
		log.Println("call timeout func")
		return true
	})
	if err != nil {
		log.Fatal(err)
	}
	ctx := g.MainContextDefault()
	mainLoop := g.NewMainLoop(ctx, false)
	go mainLoop.Run()
	time.Sleep(10 * time.Second)
	mainLoop.Quit()
}
