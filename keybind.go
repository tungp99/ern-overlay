package main

import (
	hook "github.com/robotn/gohook"
	"gopkg.in/ini.v1"
)

type Event int

const (
	RESET      Event = 0b0000
	RESUME     Event = 0b0001
	PAUSE      Event = 0b0010
	QUIT       Event = 0b0100
	HOT_RELOAD Event = 0b1000
)

var event = make(chan Event)

func hookKeybinds(cfg *ini.File) {
	reset_key_combo, toggle_key_combo, quit_key_combo, hotreload_key_combo := cfg.Section("keybind").Key("reset").Strings("+"),
		cfg.Section("keybind").Key("toggle").Strings("+"),
		cfg.Section("keybind").Key("quit").Strings("+"),
		cfg.Section("keybind").Key("hotreload").Strings("+")

	currentEv := PAUSE

	hook.Register(hook.KeyDown, reset_key_combo, func(e hook.Event) {
		currentEv = PAUSE
		event <- RESET
	})

	hook.Register(hook.KeyDown, toggle_key_combo, func(e hook.Event) {
		mask := PAUSE ^ RESUME
		currentEv ^= mask
		event <- currentEv
	})

	hook.Register(hook.KeyDown, quit_key_combo, func(e hook.Event) {
		event <- QUIT
		hook.End()
	})

	hook.Register(hook.KeyDown, hotreload_key_combo, func(e hook.Event) {
		currentEv = PAUSE
		event <- HOT_RELOAD
		cfg.Reload()
	})

	s := hook.Start()
	<-hook.Process(s)
}
