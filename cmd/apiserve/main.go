package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/mioe/hello-golang/internal/app/apiserve"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserve.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserve.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserve.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
