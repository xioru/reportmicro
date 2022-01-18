package main

import (
	"github.com/caarlos0/env/v6"
	"log"
	"reportmicro/internal/app"
	"reportmicro/internal/config"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln(err)
	}
	err := app.New(cfg).Run()
	if err != nil {
		log.Fatalln(err)
}