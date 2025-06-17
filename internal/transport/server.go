package transport

import (
	"context"
	"errors"
	"github.com/Killazius/workmate/internal/config"
	"github.com/Killazius/workmate/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"net"
	"net/http"
)

type Server struct {
	server *http.Server
	log    *zap.SugaredLogger
}

func NewServer(
	taskHandler *handlers.TaskHandler,
	log *zap.SugaredLogger,
	cfg config.HTTPConfig,
) *Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", taskHandler.CreateTask)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", taskHandler.GetTask)
			r.Delete("/", taskHandler.DeleteTask)
		})
	})

	return &Server{
		log: log,
		server: &http.Server{
			Addr:         net.JoinHostPort(cfg.Host, cfg.Port),
			ReadTimeout:  cfg.Timeout,
			WriteTimeout: cfg.Timeout,
			IdleTimeout:  cfg.IdleTimeout,
			Handler:      r,
		},
	}
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		s.log.Fatal("failed to run HTTP-server", zap.Error(err))
	}
}
func (s *Server) Run() error {
	err := s.server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}
func (s *Server) Stop(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		s.log.Error("failed to stop HTTP server", zap.Error(err))
	}
}
