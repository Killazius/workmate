package main

import (
	"context"
	"github.com/Killazius/workmate/internal/config"
	"github.com/Killazius/workmate/internal/logger"
	"github.com/Killazius/workmate/internal/repository"
	"github.com/Killazius/workmate/internal/service"
	"github.com/Killazius/workmate/internal/storage/taskstorage"
	"github.com/Killazius/workmate/internal/transport"
	"github.com/Killazius/workmate/internal/transport/handlers"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	log, err := logger.LoadFromConfig(cfg.LoggerPath)
	if err != nil {
		panic(err)
	}
	store := taskstorage.New()
	repo := repository.New(store)
	s := service.New(repo)
	handler := handlers.New(s, log)
	server := transport.NewServer(handler, log, cfg.HTTPServer)
	go server.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Info("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Stop(ctx)

}
