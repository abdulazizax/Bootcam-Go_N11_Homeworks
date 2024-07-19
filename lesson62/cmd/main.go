package main

import (
	"csb/api"
	"csb/internal/config"
	"log"
)

func main() {
	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	api, err := api.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(api.Run().Run(configs.Server.Port))
}
