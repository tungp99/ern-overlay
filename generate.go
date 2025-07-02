//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	log.Println("🔧 Generating cross-platform build...")

	build_dir := filepath.Join("dist", runtime.GOOS, runtime.GOARCH)
	err := os.MkdirAll(filepath.Join(build_dir, "assets"), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	entries, err := os.ReadDir("assets")
	if err != nil {
		log.Fatalln(err)
	}

	for _, e := range entries {
		switch runtime.GOOS {
		case "windows":
			run("cmd", "/C", "copy", "/Y", filepath.Join("assets", e.Name()), filepath.Join(build_dir, "assets", e.Name()))
		default:
			run("cp", filepath.Join("assets", e.Name()), filepath.Join(build_dir, "assets", e.Name()))
		}
	}

	out := "ern-overlay"
	if runtime.GOOS == "windows" {
		out += ".exe"
	}
	run("go", "build", "-o", filepath.Join(build_dir, out), ".")

	fmt.Println("✅ Build complete")
}
