//go:build darwin

package main

// #cgo CFLAGS: -x objective-c -fobjc-arc
// #cgo LDFLAGS: -framework Cocoa
// #include <Cocoa/Cocoa.h>
// void SetAlwaysOnTop(void* cocoaWindow) {
//     NSWindow *win = (__bridge NSWindow*)cocoaWindow;
//     [win makeFirstResponder:nil];
//     [win setLevel:NSScreenSaverWindowLevel];
//     [win setIgnoresMouseEvents:YES];
//     [win setAcceptsMouseMovedEvents:NO];
//     [win setCollectionBehavior:
//          NSWindowCollectionBehaviorCanJoinAllSpaces |
//          NSWindowCollectionBehaviorTransient |
//          NSWindowCollectionBehaviorFullScreenAuxiliary];
// }
import "C"

import "github.com/go-gl/glfw/v3.3/glfw"

func setTopMost(window *glfw.Window) {
	cocoaPtr := window.GetCocoaWindow()
	C.SetAlwaysOnTop(cocoaPtr)
}
