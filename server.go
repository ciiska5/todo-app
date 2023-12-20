package todoapp

import (
	"context"
	"net/http"
	"time"
)

// структура для запуска http-сервера
type Server struct {
	httpServer *http.Server
}

// запускает http-сервер
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          //1 MB - максимальный размер заголовка запроса, 1048576 байт
		ReadTimeout:    10 * time.Second, // 10 с
		WriteTimeout:   10 * time.Second, // 10 с
	}

	return s.httpServer.ListenAndServe()
}

// останавливает http-сервер
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
