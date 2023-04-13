package main

import (
	"cb/internal"
	"cb/internal/config"
	"log"
)

func main() {
	log.Println("config initializing")
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s\n", err.Error())
	}

	cfg := config.MustLoad()

	log.Println("Creating Application")
	app, err := internal.NewApp(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Running Application")
	app.Run()

}
