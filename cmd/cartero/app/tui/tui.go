package tui

import (
	"fmt"
	"log"

	"github.com/awesome-gocui/gocui"
	"github.com/sebport0/cartero/internal/config"
)

type TUI struct {
	client *gocui.Gui
	cfg    *config.Config
}

func NewTUI() *TUI {
	a, err := gocui.NewGui(gocui.Output256, true)
	if err != nil {
		log.Panicln(err)
	}
	return &TUI{client: a}
}

func (t *TUI) Close() {
	t.client.Close()
}

func (t *TUI) Run() error {
	normal := &Normal{}
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

type Normal struct{}

func (n *Normal) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	// Collections view
	collectionsView, err := g.SetView("collections ", 0, 0, maxX/3, maxY, 0)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	collectionsView.Clear()
	fmt.Fprintf(collectionsView, "hello view\ncol1")

	// Request view
	// Response view
	// Documentation view
	return nil
}

func (t *TUI) quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}

func (t *TUI) enableMouse() {
	t.client.Cursor = true
	t.client.Mouse = true
}
