package services

import (
	"api-repository/internal/config"
	"fmt"
	"time"
)

func GetServerStartedLogString(c *config.MainConfig, time time.Time, port int, name string) string {
	return fmt.Sprintf("SERVER %s started by PORT: %d at the TIME: %v \n"+
		"current CONF: %v", name, port, time, c)
}
