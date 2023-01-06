package ws

import (
	"fmt"
	"net/http"
	"time"

	"github.com/joaocarvalhodb1/arch-poc/shared/helpers"
)

const httpServerReadTimeout = 5 * time.Second
const httpServerWriteTimeout = 10 * time.Second

type HTTPServer struct {
	httpServer *http.Server
	port       string
	routes     http.Handler
	log        *helpers.Loggers
}

func NewHttpServer(routes http.Handler, port string, log *helpers.Loggers) *HTTPServer {
	httpServer := &http.Server{
		ReadTimeout:  httpServerReadTimeout,
		WriteTimeout: httpServerWriteTimeout,
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      routes,
	}
	server := &HTTPServer{
		httpServer: httpServer,
		port:       port,
		routes:     routes,
		log:        log,
	}
	return server
}

func (s *HTTPServer) Listen() (err error) {
	s.log.Debug("account app listening in port %s", s.httpServer.Addr)
	err = s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return err
}
