package models

import (
	"fmt"
)

type USIOConfig struct {
	Host string
	Port string
}

func (config *USIOConfig) GetUSIOUri()  string {
	return fmt.Sprintf("http://%s:%s", config.Host,config.Port)
}