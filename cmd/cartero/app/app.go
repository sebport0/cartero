package app

import (
	"github.com/sebport0/cartero/cmd/cartero/app/backend"
	"github.com/sebport0/cartero/cmd/cartero/app/tui"
	"github.com/sebport0/cartero/internal/collection"
	"github.com/sebport0/cartero/internal/config"
	"github.com/sebport0/cartero/internal/request"
)

type App struct {
	tui    *tui.TUI
	be     *backend.Backend
	config *config.Config
}

func NewApp() *App {
	c := config.NewDefaultConfig()
	t := tui.NewTUI(c)
	be := backend.NewBackend()
	return &App{tui: t, be: be, config: c}
}

func (a *App) Run() error {
	err := a.tui.Run()
	if err != nil {
		return err
	}
	defer a.tui.Close()
	return nil
}

func (a *App) CreateEmptyCollection(name string) {
	c := collection.NewCollection(name, map[string]*request.Request{})
	a.be.AddCollection(c)
}

func (a *App) AddRequestToCollection(collection string, r *request.Request) {
	a.be.AddRequest(collection, r)
}

func (a *App) DeleteRequest(collection, request string) {
	a.be.DeleteRequest(collection, request)
}
