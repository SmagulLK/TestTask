package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (Serv *Server) Run(port string, handler http.Handler) error {
	Serv.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return Serv.httpServer.ListenAndServe()
}
