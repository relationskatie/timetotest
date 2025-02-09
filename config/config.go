package config

type Config struct {
	Server   *Server
	Postgres *Postgres
}

func New() *Config {
	return &Config{
		Server: &Server{
			host: "localhost",
			port: 8001,
		},
		Postgres: &Postgres{
			host:     "localhost",
			port:     5432,
			user:     "relationskatie",
			password: "murder3472!",
			dbname:   "timetytest",
		},
	}
}
