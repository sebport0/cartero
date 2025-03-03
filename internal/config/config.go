package config

var NormalMode string = ""

type Config struct {
	Mode       string
	Fullscreen bool
}

func NewConfig(mode string, fullscreen bool) *Config {
	if mode != NormalMode {
		panic("mode is not available")
	}
	return &Config{Mode: mode, Fullscreen: fullscreen}
}
