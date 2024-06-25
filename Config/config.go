package my_go_module

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Name string `json:"name"`
}

func NewConfig(host, port string) *Config {
	return &Config{
		Host: host,
		Port: port,
	}
}
