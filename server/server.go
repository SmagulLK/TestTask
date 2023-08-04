package server

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (Serv *Server) Run(port string, handler http.Handler) error {
	log.Println("Server is running on port", port)
	Serv.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return Serv.httpServer.ListenAndServe()
}
