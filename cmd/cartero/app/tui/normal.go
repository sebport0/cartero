package tui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/sebport0/cartero/internal/config"
)

type Normal struct {
	cfg config.ModeCfg
}

func (n *Normal) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	// Collections view
	vCfg := n.cfg.Views[config.CollectionsView]
	collectionsView, err := g.SetView(
		vCfg.Name,
		vCfg.StartX,
		vCfg.StartY,
		maxX/vCfg.FinalX,
		maxY/vCfg.FinalY,
		vCfg.Overlaps)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	collectionsView.Editable = vCfg.Editable
	collectionsView.Editor = &CollectionsEditor{}
	_, err = g.SetCurrentView(vCfg.Name)
	if err != nil {
		return err
	}

	// Request view
	// Response view
	// Documentation view
	return nil
}

type CollectionsEditor struct{}

func (e *CollectionsEditor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyArrowUp || ch == 'k':
		v.MoveCursor(0, -1)
	case key == gocui.KeyArrowDown || ch == 'j':
		v.MoveCursor(0, 1)
	}
}
