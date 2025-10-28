package main

import (
	"encoding/json"
	"ensweb_crud_demo/app"
	"os"
	"os/signal"
	"syscall"

	"github.com/EnsurityTechnologies/ensweb"
	"github.com/EnsurityTechnologies/logger"
)

func main() {

	log := logger.NewDefaultLog(nil, "enscrud", logger.Debug, "./logs/", 10)

	rd, err := os.ReadFile("config.json")
	if err != nil {
		log.Error("Failed to load config", "err", err)
		return
	}

	var cfg ensweb.Config

	err = json.Unmarshal(rd, &cfg)

	if err != nil {
		log.Error("Failed to unmarshal config", "err", err)
		return
	}

	a, err := app.NewApp(&cfg, log)
	if err != nil {
		log.Error("failed to create the app", "err", err)
		return
	}
	go a.Run()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)

	<-c
	a.Stop()
	log.Info("Shutting down...")
}
