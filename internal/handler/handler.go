package handler

import (
	"githhub.com/VSBrilyakov/test-app/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		subscribe := api.Group("/subscribe")
		{
			subscribe.POST("/", h.createSubscribe)
			subscribe.GET("/:id", h.getSubscribe)
			subscribe.PUT("/:id", h.updateSubscribe)
			subscribe.DELETE("/:id", h.deleteSubscribe)
			subscribe.GET("/all", h.getAllSubscribes)
		}
	}

	return router
}
