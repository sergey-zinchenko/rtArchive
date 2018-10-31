package main

import (
	log "github.com/sirupsen/logrus"
	"rtArchive/app"
	"rtArchive/config"
)

func init() {
	config.LoadEnvironment()
	if config.IsDevelopmentEnv() {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

func main() {
	a := app.NewApp()
	err := a.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	a.ConnectGRPC()
}
