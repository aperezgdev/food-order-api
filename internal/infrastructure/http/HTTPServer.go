package http_server

import (
	"context"
	"log/slog"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/aperezgdev/food-order-api/env"
)

type Route interface {
	Pattern() string
	Method() string
	Handler(*gin.Context)
}

func NewHTTPGinServer(
	lc fx.Lifecycle,
	handler http.Handler,
	log *slog.Logger,
	env env.EnvApp,
) *http.Server {
	srv := &http.Server{Addr: ":" + env.PORT, Handler: handler}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				panic(err)
			}

			log.Info("Starting Gin Server")

			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Ending gin server")
			return nil
		},
	})
	return srv
}

func NewHTTPRouterGinGonic(routes []Route) http.Handler {
	r := gin.Default()

	for _, route := range routes {
		r.Handle(route.Method(), route.Pattern(), route.Handler)
	}

	return r
}
