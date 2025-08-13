package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// add context for os signal

	// init repo

	// init usecase

	// init logger

	// init handler

	// init router

	r := gin.Default()
	r.Use(gin.Recovery())

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
