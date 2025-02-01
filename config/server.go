package config

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

type Server struct {
	host string
	port int
}

func (s Server) GetBindAddress() string {
	log.Error("[server] GetBindAddress")
	return fmt.Sprintf("%s:%d", s.host, s.port)
}
