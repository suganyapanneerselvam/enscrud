package main

import (
	"fmt"
	"log"

	"ensweb_crud_demo/db"
	"ensweb_crud_demo/handler"
	"github.com/EnsurityTechnologies/config"
	"github.com/EnsurityTechnologies/ensweb"
	"github.com/EnsurityTechnologies/logger"
)

func main() {
	db.Connect()

	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	appLogger := logger.New(&logger.LoggerOptions{
		Name: "ensweb_crud_demo",
	})

	server, err := ensweb.NewServer(cfg, nil, appLogger)
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	server.AddRoute("/users", "GET", handler.GetUsers)
	server.AddRoute("/users/add", "POST", handler.AddUser)
	server.AddRoute("/users/update", "PUT", handler.UpdateUser)
	server.AddRoute("/users/delete", "DELETE", handler.DeleteUser)

	fmt.Println(" Server running on http://localhost:8080")
	server.Start()
}
