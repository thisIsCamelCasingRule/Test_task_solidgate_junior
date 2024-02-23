package config

type Config struct {
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`
}
