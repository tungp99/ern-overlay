//go:build windows

package main

/*
#cgo CFLAGS: -I./libs/include -DUNICODE
#cgo LDFLAGS: -L./libs -lglfw3dll -lgdi32 -lopengl32 -luser32 -lkernel32
#define GLFW_EXPOSE_NATIVE_WIN32
#include <GLFW/glfw3.h>
#include <GLFW/glfw3native.h>

GLFWwindow* global_window = NULL;

void set_glfw_window(GLFWwindow* win) {
    global_window = win;
}

void* get_win32_hwnd() {
    return glfwGetWin32Window(global_window);
}
*/
import "C"

import (
	"unsafe"

	"github.com/go-gl/glfw/v3.3/glfw"
	"golang.org/x/sys/windows"
)

var (
	user32           = windows.NewLazySystemDLL("user32.dll")
	procSetWindowPos = user32.NewProc("SetWindowPos")
)

const (
	HWND_TOPMOST   = ^uintptr(1 - 1)
	SWP_NOMOVE     = 0x0002
	SWP_NOSIZE     = 0x0001
	SWP_SHOWWINDOW = 0x0040
)

func setTopMost(window *glfw.Window) {
	C.set_glfw_window((*C.GLFWwindow)(unsafe.Pointer(window)))

	hwnd := uintptr(C.get_win32_hwnd())

	procSetWindowPos.Call(
		hwnd,
		HWND_TOPMOST,
		0, 0, 0, 0,
		SWP_NOMOVE|SWP_NOSIZE|SWP_SHOWWINDOW,
	)
}
