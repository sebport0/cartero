package config

var (
	NormalMode      = "normal"
	CollectionsView = "collections"
)

type ViewCfg struct {
	Name   string
	StartX int
	StartY int

	// FinalX should be an int that acts as a proportion of
	// the max X size. For example: FinalX = 3 translates to
	// a view's x1 = maxX/FinalX
	FinalX int

	// FinalY should be an int that acts as a proportion of
	// the max Y size. For example: FinalY = 3 translates to
	// a view's y1 = maxY/FinalY
	FinalY   int
	Overlaps byte

	// Editable enables the current view text to be edited
	Editable bool
}

type ModeCfg struct {
	Views map[string]ViewCfg
}

type Config struct {
	Modes      map[string]ModeCfg
	Fullscreen bool
}

func NewDefaultConfig() *Config {
	modes := map[string]ModeCfg{
		NormalMode: {
			Views: map[string]ViewCfg{
				CollectionsView: {
					Name:     CollectionsView,
					StartX:   0,
					StartY:   0,
					FinalX:   3,
					FinalY:   1,
					Overlaps: 0,
					Editable: true,
				},
			},
		},
	}
	return &Config{
		Modes:      modes,
		Fullscreen: true,
	}
}
