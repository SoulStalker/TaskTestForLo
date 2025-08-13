package http

import "github.com/gin-gonic/gin"

// SetupRouter запускает роутер
func SetupRouter(h *Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	g := r.Group("/tasks")
	{
		g.GET("", h.All)
		g.GET(":id", h.GetById)
		g.POST("", h.Create)
	}
	return r
}
