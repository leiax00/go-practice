package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"project-practice/api/weekly-task/service/v1"
	"project-practice/internal/weekly-task/conf"
	"project-practice/internal/weekly-task/service"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(conf *conf.Server, service *service.UserService) *http.Server {
	var opts []http.ServerOption
	if conf.Http.Network != "" {
		opts = append(opts, http.Network(conf.Http.Network))
	}
	if conf.Http.Addr != "" {
		opts = append(opts, http.Address(conf.Http.Addr))
	}
	if conf.Http.Timeout != nil {
		opts = append(opts, http.Timeout(conf.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	m := http.Middleware(
		middleware.Chain(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(log.DefaultLogger),
		),
	)
	srv.HandlePrefix("/", v1.NewUserHandler(service, m))
	return srv
}
