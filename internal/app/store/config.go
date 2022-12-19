package store

type Config struct {
	Port     string `toml:"port"`
	Host     string `toml:"host"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

func NewConfig() *Config {
	return &Config{}
}
