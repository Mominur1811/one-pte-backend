package app

import (
	"one-pte-backend/config"
	"one-pte-backend/db"
	"one-pte-backend/web"
	"sync"
)

type Application struct {
	Wg *sync.WaitGroup
}

func NewApplication() *Application {
	return &Application{}
}

func (r *Application) Init() {
	config.LoadConfig()
	db.ConnectDB()
}

func (r *Application) Run() {
	web.StartRestServer(r.Wg)
}

func (r *Application) Wait() {
	r.Wg.Wait()
}

func (r *Application) Stop() {
	db.CloseDB()
}
