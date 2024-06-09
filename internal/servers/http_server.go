package servers

import (
	"context"
	"net/http"
	"time"
)

type HTTPServer struct {
	httpServer *http.Server
}

func (s *HTTPServer) Run(port string, handler http.Handler) error {
	if port == "" {
		port = "8080"
	}

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	return s.httpServer.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
