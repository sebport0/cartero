package tui

import (
	"log"

	"github.com/awesome-gocui/gocui"
	"github.com/sebport0/cartero/internal/config"
)

type TUI struct {
	client *gocui.Gui
	config *config.Config
}

func NewTUI(cfg *config.Config) *TUI {
	a, err := gocui.NewGui(gocui.Output256, true)
	if err != nil {
		log.Panicln(err)
	}
	return &TUI{client: a, config: cfg}
}

func (t *TUI) Close() {
	t.client.Close()
}

func (t *TUI) Run() error {
	normal := &Normal{t.config.Modes[config.NormalMode]}
	t.client.SetManager(normal)
	t.enableMouse()
	err := t.setKeyBindings()
	if err != nil {
		return err
	}
	if err := t.client.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func (t *TUI) setKeyBindings() error {
	err := t.client.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, t.quit)
	if err != nil {
		return err
	}
	return nil
}

func (t *TUI) quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}

func (t *TUI) enableMouse() {
	t.client.Cursor = true
	t.client.Mouse = true
}
