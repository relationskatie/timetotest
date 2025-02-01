package config

type Config struct {
	Server *Server
}

func New() (*Config, error) {
	return &Config{
		Server: &Server{
			host: "localhost",
			port: 8001,
		},
	}, nil
}
