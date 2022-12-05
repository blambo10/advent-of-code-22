package config

import (
	"github.com/gobuffalo/envy"
)

type Config struct {
	AOC AOC
}

type AOC struct {
	URL    string
	Cookie string
}

func GetConfig() *Config {
	return &Config{
		AOC: AOC{
			URL:    envy.Get("AOC_URL", ""),
			Cookie: envy.Get("AOC_SESSION_COOKIE", ""),
		},
	}
}

func GetAOCConfig() *AOC {
	return &GetConfig().AOC
}
