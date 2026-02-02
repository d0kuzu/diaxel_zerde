package config

import "time"

type Config struct {
	HTTPPort string

	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration

	AccessSecret  string
	RefreshSecret string
}

func MustLoad() *Config {
	return &Config{
		HTTPPort: "8080",

		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 30 * 24 * time.Hour,

		AccessSecret:  "ACCESS_SECRET",
		RefreshSecret: "REFRESH_SECRET",
	}
}
