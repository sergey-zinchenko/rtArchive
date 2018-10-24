package main

import (
	"log"
	"rtArchive/app"
)

func main() {
	a := app.NewApp()
	err := a.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	go a.ConnectGRPC()
}
