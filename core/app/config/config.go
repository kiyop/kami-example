package config

type Config struct {
	General General `toml:"general"`
	Secret  Secret  `toml:"secret"`
	Session Session `toml:"session"`
}

type General struct{}

type Secret struct {
	SiteKey      string `toml:"site_key"`
	StretchCount int    `toml:"stretch_count"`
}

type Session struct {
	TTL int `toml:"ttl"`
}

func DefaultConfig() Config {
	return Config{
		Secret: Secret{
			StretchCount: 5000,
		},
		Session: Session{
			TTL: 14 * 24 * 60 * 60, // 2 週間
		},
	}
}
