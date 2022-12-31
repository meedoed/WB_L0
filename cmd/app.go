package main

import (
	apiserver2 "WB_L0/internal/apiserver"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver2.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}
	s := apiserver2.New(config)

	if err = s.Start(); err != nil {
		log.Fatal(err)
	}
}
