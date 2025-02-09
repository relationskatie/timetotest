package config

import (
	"fmt"
)

type Server struct {
	host string
	port int
}

func (s Server) GetBindAddress() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}
