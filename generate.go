//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func run(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Command failed: %s %v\n", name, args)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("🔧 Generating cross-platform build...")

	os.MkdirAll("dist", os.ModePerm)

	switch runtime.GOOS {
	case "windows":
		run("cmd", "/C", "copy", "/Y", "assets/config.ini", "dist/config.ini")
		run("cmd", "/C", "copy", "/Y", "assets/font.ttf", "dist/Montserrat-Medium.ttf")
	default:
		run("cp", "assets/config.ini", "dist/")
		run("cp", "assets/Montserrat-Medium.ttf", "dist/")
	}

	out := "ern-overlay"
	if runtime.GOOS == "windows" {
		out += ".exe"
	}
	run("go", "build", "-o", "dist/"+out, ".")

	fmt.Println("✅ Build complete")
}
