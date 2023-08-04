package handler

import (
	"TestTask/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	video := router.Group("/videos")
	{
		video.GET("", h.GetVideos)
	}
	return router
}
