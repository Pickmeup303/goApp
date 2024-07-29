package config

type config struct {
	ServerPort string
	ServerHost string
}

func DefaultConfig() *config {
	return &config{
		ServerPort: "8080",
		ServerHost: "localhost",
	}
}
