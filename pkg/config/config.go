package config

type Config struct {
}

func Get() *Config {
	cfg := Config{}
	return &cfg
}
