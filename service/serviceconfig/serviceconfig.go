package serviceconfig

import (
	"github.com/dzwiedz90/helpdesk-service-users/driver"
	"github.com/dzwiedz90/helpdesk-service-users/logs"
)

type ServerConfig struct {
	DB     *driver.DB
	Logger logs.LoggerInterface
}
