package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// add context for os signal
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// init repo

	// init usecase

	// init logger

	// init handler

	// init router

	srv := &http.Server{
		Addr: ":8080",
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
