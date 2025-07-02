package main

import (
	"log"

	"gopkg.in/ini.v1"
)

func load_config() *ini.File {
	var cfg, err = ini.Load("./assets/config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	return cfg
}
