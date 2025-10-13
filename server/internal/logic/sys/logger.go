package sys

import (
	"context"
	"server/internal/service"
)

type sSysLogger struct{}

func NewSysLogger() *sSysLogger {
	return &sSysLogger{}
}

func init() {
	service.RegisterSysLogger(NewSysLogger())
}

// PushAdminLoginLog 推送adminLog
func (s *sSysLogger) PushAdminLoginLog(ctx context.Context) {

	return
}

func (s *sSysLogger) PushApiLoginLog(ctx context.Context) {

	return
}
