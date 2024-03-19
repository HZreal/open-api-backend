package config

type GinConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (g *GinConfig) GetAddr() string {
	return g.Host + ":" + g.Port
}
