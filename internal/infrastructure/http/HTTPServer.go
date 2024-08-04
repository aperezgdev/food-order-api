package http_server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"

	"github.com/aperezgdev/food-order-api/env"
)

type Route interface {
	Pattern() string
	http.Handler
}

func NewHTTPServer(lc fx.Lifecycle, env env.EnvApp, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: ":" + env.PORT, Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func NewServerMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}

	return mux
}

