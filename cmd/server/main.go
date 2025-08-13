package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	httpDelivery "github.com/soulstalker/task-api/internal/delivery/http"
	"github.com/soulstalker/task-api/internal/repo/memory"
	"github.com/soulstalker/task-api/internal/usecase"
)

func main() {
	// add context for os signal
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// init repo
	repo := memory.NewTaskRepoIM()
	// init usecase
	uc := usecase.NewTaskUC(repo)
	// init logger

	// init handler
	h := httpDelivery.NewHandler(uc)
	// init router
	r := httpDelivery.SetupRouter(h)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Канал для ошибок сервера
	serverErr := make(chan error, 1)

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Println("starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	log.Println("server started, waiting for requests...")

	// Ждем либо сигнала завершения, либо ошибки сервера
	select {
	case <-ctx.Done():
		log.Println("shutdown signal received")
	case err := <-serverErr:
		log.Printf("server error: %v", err)
		return
	}

	// Graceful shutdown
	log.Println("shutting down server...")
	shCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shCtx); err != nil {
		log.Printf("server shutdown error: %v", err)
	} else {
		log.Println("server shutdown complete")
	}
}
