package tui

import (
	"log"

	"github.com/awesome-gocui/gocui"
	"github.com/sebport0/cartero/internal/config"
)

type TUI struct {
	app *gocui.Gui
	cfg *config.Config
}

func NewTUI() *TUI {
	a, err := gocui.NewGui(gocui.Output256, true)
	if err != nil {
		log.Panicln(err)
	}
	return &TUI{app: a}
}

func (t *TUI) Close() {
	t.app.Close()
}

func (t *TUI) Run() error {
	err := t.setKeyBindings()
	if err != nil {
		return err
	}
	t.enableMouse()
	// t.app.SetManagerFunc(t.normalLayout)
	if err := t.app.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func (t *TUI) setKeyBindings() error {
	err := t.app.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, t.quit)
	if err != nil {
		return err
	}
	return nil
}

func (t *TUI) normalLayout(*gocui.Gui) error {
	return nil
}

func (t *TUI) quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}

func (t *TUI) enableMouse() {
	t.app.Cursor = true
	t.app.Mouse = true
}
