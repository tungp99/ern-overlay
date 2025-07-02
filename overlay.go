package main

import (
	"log"
	"os"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gltext"
	"gopkg.in/ini.v1"
)

type Overlay struct {
	Config *ini.File
	Window *glfw.Window
	Font   *gltext.Font
}

func (o *Overlay) CreateWindow() {
	width, height, fontsz := o.Config.Section("visual").Key("width").MustInt(72),
		o.Config.Section("visual").Key("height").MustInt(32),
		o.Config.Section("visual").Key("fontsz").MustInt(48)

	pos_x, pos_y := o.Config.Section("visual").Key("position_x").MustInt(glfw.GetPrimaryMonitor().GetVideoMode().Width/2),
		o.Config.Section("visual").Key("position_y").MustInt(0)

	window, err := glfw.CreateWindow(width, height, "ERN Overlay", nil, nil)
	if err != nil {
		log.Panicln("failed to create window:", err)
	}

	window.SetPos(pos_x, pos_y)
	window.MakeContextCurrent()
	o.Window = window

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Panicln(err)
	}
	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.LIGHTING)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	f, err := os.Open(o.Config.Section("visual").Key("font").MustString("./Montserrat-Medium.ttf"))
	if err != nil {
		log.Panicln("cannot load font:", err)
	}
	defer f.Close()

	font, err := gltext.LoadTruetype(f, int32(fontsz), 32, 127, gltext.LeftToRight)
	if err != nil {
		log.Fatalln("cannot load font:", err)
	}
	o.Font = font
}

func (o *Overlay) DrawFrame(d time.Duration) {
	gl.ClearColor(0.0, 0.0, 0.0, 0.4)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()

	gl.Color3f(1, 1, 1)

	o.Font.Printf(0, 0, "%02d:%02d", int(d.Minutes()), int(d.Seconds()))

	o.Window.SwapBuffers()
}

func (o *Overlay) Initialize() {
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		log.Panicln("failed to init glfw:", err)
	}

	glfw.WindowHint(glfw.Focused, glfw.False)
	glfw.WindowHint(glfw.FocusOnShow, glfw.False)
	glfw.WindowHint(glfw.Floating, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
}

func (o *Overlay) Destroy() {
	o.DestroyWindow()
	glfw.Terminate()
}

func (o *Overlay) DestroyWindow() {
	if o.Font != nil {
		o.Font.Release()
	}
	if o.Window != nil {
		o.Window.Destroy()
	}
}
