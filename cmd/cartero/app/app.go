package app

import (
	"github.com/sebport0/cartero/cmd/cartero/app/backend"
	"github.com/sebport0/cartero/cmd/cartero/app/tui"
	"github.com/sebport0/cartero/internal/collection"
	"github.com/sebport0/cartero/internal/request"
)

type App struct {
	tui *tui.TUI
	be  *backend.Backend
}

func NewApp() *App {
	t := tui.NewTUI()
	be := backend.NewBackend()
	return &App{tui: t, be: be}
}

func (a *App) CreateEmptyCollection(name string) {
	c := collection.NewCollection(name, map[string]*request.Request{})
	a.be.AddCollection(c)
}

func (a *App) AddRequestToCollection(collection string, r *request.Request) {
	a.be.AddRequest(collection, r)
}

func (a *App) DeleteReqquest(collection, request string) {
	a.be.DeleteRequest(collection, request)
}
