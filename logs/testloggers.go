package logs

import (
	"github.com/dzwiedz90/helpdesk-service-users/config"
)

type TestLogger struct{}

func NewTestLoggers(a *config.Config) TestLogger {
	cfg = a
	return TestLogger{}
}

func (t *TestLogger) InfoLogger(message string) {}

func (t *TestLogger) ErrorLogger(message string) {}
