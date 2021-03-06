package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AuthenticationService) AuthenticationService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthenticationService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthenticationService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthenticationService) AuthenticationService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetOpenid(ctx context.Context, code string) (status bool, errinfo string, data *UserRes) {
	defer func() {
		l.logger.Log("method", "GetOpenid", "code", code, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetOpenid(ctx, code)
}
func (l loggingMiddleware) AdminLogin(ctx context.Context, name string, password string) (status bool, errinfo string, data string) {
	defer func() {
		l.logger.Log("method", "AdminLogin", "name", name, "password", password, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.AdminLogin(ctx, name, password)
}
func (l loggingMiddleware) EnterpriseLogin(ctx context.Context, name string, password string) (status bool, errinfo string, data string) {
	defer func() {
		l.logger.Log("method", "EnterpriseLogin", "name", name, "password", password, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.EnterpriseLogin(ctx, name, password)
}
func (l loggingMiddleware) Logout(ctx context.Context) (status bool, errinfo string, data string) {
	defer func() {
		l.logger.Log("method", "Logout", "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.Logout(ctx)
}
