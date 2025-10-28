package app

import (
	"ensweb_crud_demo/db"

	"github.com/EnsurityTechnologies/ensweb"
	"github.com/EnsurityTechnologies/logger"
	"gorm.io/gorm"
)

type App struct {
	ensweb.Server
	db  *gorm.DB
	log logger.Logger
}

func NewApp(cfg *ensweb.Config, log logger.Logger) (*App, error) {
	a := &App{
		log: log,
	}
	var err error

	a.Server, err = ensweb.NewServer(cfg, nil, log)
	if err != nil {
		log.Error("Failed to create server:", err)
		return nil, err
	}

	a.db, err = db.OpenDB(log)
	if err != nil {
		log.Error("Failed to open db", err)
		return nil, err
	}

	a.AddRoute("/users", "GET", a.GetUsers)
	a.AddRoute("/users/add", "POST", a.AddUser)
	a.AddRoute("/users/update", "PUT", a.UpdateUser)
	a.AddRoute("/users/delete", "DELETE", a.DeleteUser)

	return a, nil
}

func (a *App) Run() {
	a.Start()
}

func (a *App) Stop() {
	a.Shutdown()
}
