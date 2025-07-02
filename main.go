package main

import (
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

//go:generate go run generate.go
func main() {
	cfg := load_config()
	go hookKeybinds(cfg)

	o := Overlay{}
	o.Config = cfg
	o.Initialize()
	defer o.Destroy()

	o.CreateWindow()

	countup := 0
	ticker := time.NewTicker(time.Second)
	ticker.Stop()

	for !o.Window.ShouldClose() {
		select {
		case state := <-event:
			switch state {
			case RESET:
				countup = 0
				o.DrawFrame(time.Duration(countup * int(time.Second)))
				ticker.Reset(time.Second)
				ticker.Stop()
			case RESUME:
				ticker.Reset(time.Second)
			case PAUSE:
				ticker.Stop()
			case QUIT:
				ticker.Stop()
				o.Window.SetShouldClose(true)
			case HOT_RELOAD:
				countup = 0
				o.DestroyWindow()
				o.CreateWindow()
				ticker.Reset(time.Second)
				ticker.Stop()
			}

		case <-ticker.C:
			countup++
			o.DrawFrame(time.Duration(countup * int(time.Second)))
		default:
			glfw.PollEvents()
		}
	}
}
