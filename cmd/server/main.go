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
		Addr:    "localhost:8080",
		Handler: r,
	}

	<-ctx.Done() // Ждем сигнала на закрытие
	log.Println("shutdown signal received")

	shCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shCtx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}

	log.Println("shutdown complete")
}
