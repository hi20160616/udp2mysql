package server

import (
	"context"
	"log"
	"net/http"

	"github.com/hi20160616/udp2mysql/configs"
	myerr "github.com/hi20160616/udp2mysql/errors"
	"github.com/hi20160616/udp2mysql/web"
)

type webServer struct {
	s *web.WebServer
}

func NewWebServer() *webServer {
	return &webServer{s: web.NewWebServer(configs.V.Web.Addr)}
}

func (ws *webServer) Start(ctx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e)
			myerr.PanicLog(e)
		}
	}()

	if err := ws.s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return ctx.Err()
}

func (ws *webServer) Stop(ctx context.Context) error {
	if err := ws.s.Shutdown(context.Background()); err != nil {
		return err
	}
	return ctx.Err()
}
