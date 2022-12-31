package reader

type Config struct {
	Address  string `toml:"address"`
	Cluster  string `toml:"cluster"`
	ClientID string `toml:"client_id"`
}
