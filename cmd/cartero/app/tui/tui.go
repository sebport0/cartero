package tui

import (
	"github.com/rivo/tview"
	"github.com/sebport0/cartero/internal/config"
)

type TUI struct {
	app    *tview.Application
	config *config.Config
}

func NewTUI() *TUI {
	a := tview.NewApplication()
	return &TUI{app: a}
}
