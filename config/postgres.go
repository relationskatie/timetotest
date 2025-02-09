package config

import "fmt"

type Postgres struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func (p *Postgres) GetAddress() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", p.user, p.password, p.host, p.port, p.dbname)
}
